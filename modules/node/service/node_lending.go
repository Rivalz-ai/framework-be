package service

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"sync"
	"time"

	abiLending "github.com/Rivalz-ai/framework-be/define/abi"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func (n *Node) LendingGetBorrowFeePerMonth(ctx context.Context, lendingAddress string) (*big.Int, error) {
	parsedABI, err := abi.JSON(strings.NewReader(abiLending.LENDING_ABI))
	if err != nil {
		return nil, err
	}

	data, err := parsedABI.Pack("calculateBorrowFee", big.NewInt(30))
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(lendingAddress)

	// Create a call message to the contract
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	// Call the contract
	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call the contract: "+err.Error(), "LendingGetBorrowFeePerMonth", lendingAddress)
		return nil, err
	}
	if len(result) == 0 {
		log.Error("No result returned from the contract", "LendingGetBorrowFeePerMonth", lendingAddress)
		return nil, nil
	}
	// Parse the result
	var data_resp *big.Int
	err = parsedABI.UnpackIntoInterface(&data_resp, "calculateBorrowFee", result)
	if err != nil {
		log.Error("Failed to unpack the result: "+err.Error(), "LendingGetBorrowFeePerMonth", lendingAddress)
		return nil, err
	}

	return data_resp, nil
}

func (n *Node) LendingTotalLendingAmount(ctx context.Context, lendingAddress string) (*big.Int, error) {
	parsedABI, err := abi.JSON(strings.NewReader(abiLending.LENDING_ABI))
	if err != nil {
		return nil, err
	}

	data, err := parsedABI.Pack("sTotalLentTokens")
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(lendingAddress)

	// Create a call message to the contract
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	// Call the contract
	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call the contract: "+err.Error(), "LendingTotalLendingAmount", lendingAddress)
		return nil, err
	}
	if len(result) == 0 {
		log.Error("No result returned from the contract", "LendingTotalLendingAmount", lendingAddress)
		return nil, nil
	}

	// Parse the result
	var data_resp *big.Int
	err = parsedABI.UnpackIntoInterface(&data_resp, "sTotalLentTokens", result)
	if err != nil {
		log.Error("Failed to unpack the result: "+err.Error(), "LendingTotalLendingAmount", lendingAddress)
		return nil, err
	}

	return data_resp, nil
}

func (n *Node) LendingGetTotalBorrowAmount(ctx context.Context, lendingAddress string) (*big.Int, error) {
	parsedABI, err := abi.JSON(strings.NewReader(abiLending.LENDING_ABI))
	if err != nil {
		return nil, err
	}

	data, err := parsedABI.Pack("sTotalBorrowedTokens")
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(lendingAddress)

	// Create a call message to the contract
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	// Call the contract
	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call the contract: "+err.Error(), "LendingGetTotalBorrowAmount", lendingAddress)
		return nil, err
	}
	if len(result) == 0 {
		log.Error("No result returned from the contract", "LendingGetTotalBorrowAmount", lendingAddress)
		return nil, nil
	}

	// Parse the result
	var data_resp *big.Int
	err = parsedABI.UnpackIntoInterface(&data_resp, "sTotalBorrowedTokens", result)
	if err != nil {
		log.Error("Failed to unpack the result: "+err.Error(), "LendingGetTotalBorrowAmount", lendingAddress)
		return nil, err
	}

	return data_resp, nil
}

func (n *Node) LendingGetBorrowInfoOfBorrower(ctx context.Context, lendingAddress string, borrower string) ([]map[string]interface{}, error) {
	contractAddress := common.HexToAddress(lendingAddress)
	lendingContract, err := abiLending.NewLending(contractAddress, n.Client)
	if err != nil {
		log.Error("Failed to create the contract: "+err.Error(), "LendingGetBorrowInfoOfBorrower", lendingAddress)
		return nil, err
	}

	tokenIds, err := lendingContract.BorrowedTokensOfOwner(nil, common.HexToAddress(borrower))
	if err != nil {
		return nil, err
	}

	blockPerDay, err := lendingContract.SBlocksPerDay(nil)
	if err != nil {
		return nil, err
	}

	secondsPerBlock := new(big.Int).Div(new(big.Int).SetInt64(24*60*60), blockPerDay)

	borrowInfo := make([]map[string]interface{}, len(tokenIds))
	var wg sync.WaitGroup
	var errCapture error

	for i, tokenId := range tokenIds {
		wg.Add(1)
		go func(i int, tokenId *big.Int) {
			defer wg.Done()
			info, err := lendingContract.GetBorrowInfo(nil, tokenId)
			if err != nil {
				log.Error("Failed to get borrow info: "+err.Error(), "LendingGetBorrowInfoOfBorrower", lendingAddress)
				errCapture = err
				return
			}

			block, err := n.GetBlockByNumber(ctx, info.StartBlock)
			if err != nil {
				log.Error("Failed to get block: "+err.Error(), "LendingGetBorrowInfoOfBorrower", lendingAddress)
				errCapture = err
				return
			}

			startTimestamp := block.Timestamp
			blockDifference := new(big.Int).Sub(info.EndBlock, info.StartBlock)
			durationInSeconds := new(big.Int).Mul(blockDifference, secondsPerBlock)
			startTimestampInt, ok := new(big.Int).SetString(startTimestamp[2:], 16)
			if !ok {
				log.Error("Failed to convert start timestamp to big.Int", "LendingGetBorrowInfoOfBorrower", lendingAddress)
				errCapture = errors.New("failed to convert start timestamp to big.Int")
				return
			}
			endTime := new(big.Int).Add(startTimestampInt, durationInSeconds)
			endTimeDate := time.Unix(int64(endTime.Uint64()), 0)

			borrowInfo[i] = map[string]interface{}{
				"borrower":    info.Borrower.Hex(),
				"startBlock":  info.StartBlock.Int64(),
				"endBlock":    info.EndBlock.Int64(),
				"feePerBlock": info.FeePerBlock.Int64(),
				"mutual":      &endTimeDate,
				"tokenId":     tokenId.Int64(),
			}
		}(i, tokenId)
	}

	wg.Wait()

	if errCapture != nil {
		return nil, errCapture
	}

	return borrowInfo, nil
}

func (n *Node) LendingPendingReward(ctx context.Context, lendingAddress string, lender string) (map[string]interface{}, error) {
	lendingContract, err := abiLending.NewLending(common.HexToAddress(lendingAddress), n.Client)
	if err != nil {
		return nil, err
	}

	rewardToken, err := lendingContract.SRewardToken(nil)
	if err != nil {
		return nil, err
	}

	pendingReward, err := lendingContract.GetPendingRewards(nil, common.HexToAddress(lender))
	if err != nil {
		return nil, err
	}

	lendingAmount, err := lendingContract.LentTokenCount(nil, common.HexToAddress(lender))
	if err != nil {
		return nil, err
	}

	totalLendingAmount, err := lendingContract.STotalLentTokens(nil)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"rewardToken":        rewardToken.Hex(),
		"pendingRewards":     pendingReward.Int64(),
		"myLendingAmount":    lendingAmount.Int64(),
		"totalLendingAmount": totalLendingAmount.Int64(),
		"logo":               "https://api.rivalz.ai/fragmentz/rome/rewards/reward_default.png",
	}, nil
}
