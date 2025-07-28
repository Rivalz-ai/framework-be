package service

import (
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	//"github.com/Rivalz-ai/framework-be/modules/node/dto"
	"context"

	abistake "github.com/Rivalz-ai/framework-be/define/abi"
	"github.com/ethereum/go-ethereum"

	//"errors"
	//"fmt"
	"math/big"
)

/*
	async claimedResourceAmount(resourceClaimerAddress: string, walletAddress: string) {
	    const resourceClaimerContract = this.createContract(
	      resourceClaimerAddress,
	      RESOURCE_CLAIMER_ABI
	    )
	    const claimedAmount = await resourceClaimerContract.getClaimedAmount(walletAddress)
	    return Number(claimedAmount)
	  }
*/
func (n *Node) ClaimedResourceAmount(ctx context.Context, resource_claimer_address, wallet_address string) (int64, error) {
	logData := utils.Dictionary()
	logData["resource_claimer_address"] = resource_claimer_address
	logData["wallet_address"] = wallet_address

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(abistake.RESOURCE_CLAIM_ABI))
	if err != nil {
		log.Error("Failed to parse ABI:"+err.Error(), "ClaimedResourceAmount", logData)
		return 0, err
	}

	// Encode function call
	walletAddress := common.HexToAddress(wallet_address)
	data, err := parsedABI.Pack("getClaimedAmount", walletAddress)
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "ClaimedResourceAmount", logData)
		return 0, err
	}

	// Call contract
	contractAddress := common.HexToAddress(resource_claimer_address)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "ClaimedResourceAmount", logData)
		return 0, err
	}
	if len(result) == 0 {
		log.Warn("Empty result", "ClaimedResourceAmount", logData)
		return 0, nil
	}
	// Decode result
	var claimedAmount *big.Int
	err = parsedABI.UnpackIntoInterface(&claimedAmount, "getClaimedAmount", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "ClaimedResourceAmount", logData)
		return 0, err
	}

	return claimedAmount.Int64(), nil
}
