package service

import (
	"context"
	"strings"

	abistake "github.com/Rivalz-ai/framework-be/define/abi"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/Rivalz-ai/framework-be/modules/node/dto"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	//"errors"
	//"fmt"
	"math/big"
)

// ExtendConfig.StakingMulticallContract
func (n *Node) GetResourceStakeBalanceBatch(ctx context.Context, wallet_address string, staking_contract_address []string, multicall_contract_address string) (int64, error) {
	logData := utils.Dictionary()
	logData["wallet_address"] = wallet_address
	logData["staking_contract_address"] = staking_contract_address
	//
	walletAddress := common.HexToAddress(wallet_address)
	multiCallContractAddress := common.HexToAddress(multicall_contract_address)
	//
	var stakingContracts = []common.Address{}
	for _, address := range staking_contract_address {
		stakingContracts = append(stakingContracts, common.HexToAddress(address))
	}
	//
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(abistake.STAKING_MULTICALL_ABI))
	if err != nil {
		log.Error("Failed to parse ABI:"+err.Error(), "GetResourceStakeBalanceBatch", logData)
		return 0, err
	}
	// Encode input data
	data, err := parsedABI.Pack("stakedOfUsers", walletAddress, stakingContracts)
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "GetResourceStakeBalanceBatch", logData)
		return 0, err
	}
	// Call contract
	callMsg := ethereum.CallMsg{
		To:   &multiCallContractAddress,
		Data: data,
	}

	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "GetResourceStakeBalanceBatch", logData)
		return 0, err
	}
	//
	var stakedAmounts []*big.Int
	err = parsedABI.UnpackIntoInterface(&stakedAmounts, "stakedOfUsers", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetResourceStakeBalanceBatch", logData)
		return 0, err
	}
	total_amount := int64(0)
	for _, stake := range stakedAmounts {
		total_amount += stake.Int64()
	}
	return total_amount, nil
}
func (n *Node) GetResourceStakeBalance(ctx context.Context, wallet_address, staking_contract_address string) ([]dto.StakeInfo, error) {
	logData := utils.Dictionary()
	logData["wallet_address"] = wallet_address
	logData["staking_contract_address"] = staking_contract_address
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(abistake.RESOURCE_STAKING_ABI))
	if err != nil {
		log.Error("Failed to parse ABI:"+err.Error(), "GetResourceStakeBalance", logData)
		return nil, err
	}

	// Encode function call data
	userAddress := common.HexToAddress(wallet_address) // Thay bằng địa chỉ user cần query
	contractAddress := common.HexToAddress(staking_contract_address)
	data, err := parsedABI.Pack("getStakeInfo", userAddress)
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "GetResourceStakeBalance", logData)
		return nil, err
	}

	// Call contract
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "GetResourceStakeBalance", logData)
		return nil, err
	}

	// Decode result
	var stakeInfo []dto.StakeInfo
	//
	err = parsedABI.UnpackIntoInterface(&stakeInfo, "getStakeInfo", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetResourceStakeBalance", logData)
		return nil, err
	}

	return stakeInfo, nil
}

func (n *Node) GetTokenRewards(ctx context.Context, staking_contract_address string) ([]string, error) {
	logData := utils.Dictionary()
	logData["staking_contract_address"] = staking_contract_address
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(abistake.RESOURCE_STAKING_ABI))
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetTokenReward", logData)
		return nil, err
	}

	// Encode function call
	data, err := parsedABI.Pack("getRewardTokens")
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "GetTokenReward", logData)
		return nil, err
	}

	// Call contract
	contractAddress := common.HexToAddress(staking_contract_address)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "GetTokenReward", logData)
		return nil, err
	}

	// Decode result
	var rewardTokens []common.Address
	err = parsedABI.UnpackIntoInterface(&rewardTokens, "getRewardTokens", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetTokenReward", logData)
		return nil, err
	}

	var tokenAddresses []string
	for _, token := range rewardTokens {
		if token != (common.Address{}) { //0x00000000000
			tokenAddresses = append(tokenAddresses, token.Hex())
		}
	}
	return tokenAddresses, nil
}

func (n *Node) GetTotalStakedTokens(ctx context.Context, staking_contract_address string) (int64, error) {
	logData := utils.Dictionary()
	logData["staking_contract_address"] = staking_contract_address
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(abistake.RESOURCE_STAKING_ABI))
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetTotalStakedTokens", logData)
		return 0, err
	}

	// Encode function call
	data, err := parsedABI.Pack("totalStakedTokens")
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "GetTotalStakedTokens", logData)
		return 0, err
	}

	// Call contract
	contractAddress := common.HexToAddress(staking_contract_address)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "GetTotalStakedTokens", logData)
		return 0, err
	}

	// Decode result
	// Decode result
	var totalStakedTokens *big.Int
	err = parsedABI.UnpackIntoInterface(&totalStakedTokens, "totalStakedTokens", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetTotalStakedTokens", logData)
		return 0, err
	}
	return totalStakedTokens.Int64(), nil
}

func (n *Node) GetPendingRewards(ctx context.Context, wallet_address, staking_contract_address string) ([]dto.PendingReward, error) {
	logData := utils.Dictionary()
	logData["staking_contract_address"] = staking_contract_address
	logData["wallet_address"] = wallet_address
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(abistake.RESOURCE_STAKING_ABI))
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetPendingRewards", logData)
		return nil, err
	}
	walletAddress := common.HexToAddress(wallet_address)
	// Encode function call
	data, err := parsedABI.Pack("getPendingRewards", walletAddress)
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "GetPendingRewards", logData)
		return nil, err
	}

	// Call contract
	contractAddress := common.HexToAddress(staking_contract_address)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "GetPendingRewards", logData)
		return nil, err
	}

	// Decode result
	// Decode result
	type PendingRewards struct {
		Tokens  []common.Address
		Amounts []*big.Int
	}
	var rewards PendingRewards
	err = parsedABI.UnpackIntoInterface(&rewards, "getPendingRewards", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetPendingRewards", logData)
		return nil, err
	}
	if len(rewards.Tokens) != len(rewards.Amounts) {
		logData["rewards"] = rewards
		log.Error("Invalid result from node", "GetPendingRewards", logData)
		return nil, err
	}
	//convert result to slice
	var pendingRewards []dto.PendingReward
	for i, token := range rewards.Tokens {
		if token == (common.Address{}) {
			continue
		}
		pendingRewards = append(pendingRewards, dto.PendingReward{
			TokenAddress: token.Hex(),
			Amount:       rewards.Amounts[i].Int64(),
		})
		//fmt.Println("Token:", token.Hex(), "Amount:", rewards.Amounts[i].String())
	}
	return pendingRewards, nil
}

func (n *Node) GetRewardPerDay(ctx context.Context, staking_contract_address, reward_token_address string) (int64, error) {
	logData := utils.Dictionary()
	logData["staking_contract_address"] = staking_contract_address
	logData["reward_token_address"] = reward_token_address

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(abistake.RESOURCE_STAKING_ABI))
	if err != nil {
		log.Error("Failed to parse ABI:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}

	// Encode function call for getCurrentRewardRate
	rewardTokenAddress := common.HexToAddress(reward_token_address)
	data, err := parsedABI.Pack("getCurrentRewardRate", rewardTokenAddress)
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}

	// Call contract for getCurrentRewardRate
	contractAddress := common.HexToAddress(staking_contract_address)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}

	// Decode result for getCurrentRewardRate
	var rewardPerBlock *big.Int
	err = parsedABI.UnpackIntoInterface(&rewardPerBlock, "getCurrentRewardRate", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}

	// Encode function call for sBlocksPerDay
	data, err = parsedABI.Pack("sBlocksPerDay")
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}

	// Call contract for sBlocksPerDay
	callMsg.Data = data
	result, err = n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}

	// Decode result for sBlocksPerDay
	var blocksPerDay *big.Int
	err = parsedABI.UnpackIntoInterface(&blocksPerDay, "sBlocksPerDay", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}

	// Calculate reward per month
	rewardPerDay := new(big.Int).Mul(rewardPerBlock, blocksPerDay)

	return rewardPerDay.Int64(), nil
}

func (n *Node) GetRewardPerMonth(ctx context.Context, staking_contract_address, reward_token_address string) (int64, error) {
	rewardPerDay, err := n.GetRewardPerDay(ctx, staking_contract_address, reward_token_address)
	if err != nil {
		return 0, err
	}
	rewardPerMonth := new(big.Int).Mul(big.NewInt(rewardPerDay), big.NewInt(30))
	return rewardPerMonth.Int64(), nil
}

func (n *Node) GetRewardPerWeek(ctx context.Context, staking_contract_address, reward_token_address string) (int64, error) {
	rewardPerDay, err := n.GetRewardPerDay(ctx, staking_contract_address, reward_token_address)
	if err != nil {
		return 0, err
	}
	rewardPerMonth := new(big.Int).Mul(big.NewInt(rewardPerDay), big.NewInt(7))
	return rewardPerMonth.Int64(), nil
}

func (n *Node) SBlocksPerDay(ctx context.Context, staking_contract_address string) (int64, error) {
	logData := utils.Dictionary()
	logData["staking_contract_address"] = staking_contract_address

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(abistake.RESOURCE_STAKING_ABI))
	if err != nil {
		log.Error("Failed to parse ABI:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}

	// Call contract for getCurrentRewardRate
	contractAddress := common.HexToAddress(staking_contract_address)
	// Encode function call for sBlocksPerDay
	data, err := parsedABI.Pack("sBlocksPerDay")
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	// Call contract for sBlocksPerDay
	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}

	// Decode result for sBlocksPerDay
	var blocksPerDay *big.Int
	err = parsedABI.UnpackIntoInterface(&blocksPerDay, "sBlocksPerDay", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetRewardPerDay", logData)
		return 0, err
	}

	return blocksPerDay.Int64(), nil
}
