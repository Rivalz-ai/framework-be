package service

import (
	"context"
	"math/big"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func (n *Node) GetTotalSupply(ctx context.Context, tokenAddress string, ABI string) (string, error) {
	// get from node

	logData := utils.Dictionary()
	logData["token_address"] = tokenAddress

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		log.Error("Failed to parse ABI:"+err.Error(), "GetTotalSupply", logData)
		return "", err
	}

	// Encode function call
	data, err := parsedABI.Pack("totalSupply")
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "GetTotalSupply", logData)
		return "", err
	}

	// Call contract
	contractAddress := common.HexToAddress(tokenAddress)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "GetTotalSupply", logData)
		return "", err
	}
	if len(result) == 0 {
		log.Warn("Empty result", "GetTotalSupply", logData)
		return "", nil
	}

	var totalSupply *big.Int
	err = parsedABI.UnpackIntoInterface(&totalSupply, "totalSupply", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetTotalSupply", logData)
		return "", err
	}

	return totalSupply.String(), nil
}

func (n *Node) GetDecimals(ctx context.Context, tokenAddress string, ABI string) (int, error) {
	// get from node

	logData := utils.Dictionary()
	logData["token_address"] = tokenAddress

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		log.Error("Failed to parse ABI:"+err.Error(), "GetDecimals", logData)
		return 0, err
	}

	// Encode function call
	data, err := parsedABI.Pack("decimals")
	if err != nil {
		log.Error("Failed to pack data:"+err.Error(), "GetDecimals", logData)
		return 0, err
	}

	// Call contract
	contractAddress := common.HexToAddress(tokenAddress)
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call contract:"+err.Error(), "GetDecimals", logData)
		return 0, err
	}
	if len(result) == 0 {
		log.Warn("Empty result", "GetDecimals", logData)
		return 0, nil
	}

	var decimals *big.Int
	err = parsedABI.UnpackIntoInterface(&decimals, "decimals", result)
	if err != nil {
		log.Error("Failed to unpack result:"+err.Error(), "GetDecimals", logData)
		return 0, err
	}

	return int(decimals.Int64()), nil
}
