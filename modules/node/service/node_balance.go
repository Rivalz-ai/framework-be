package service

import (
	"context"
	"math/big"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	//"fmt"
)

func (n *Node) GetBalance(ctx context.Context, walletAddress, token_address, ABI string, decimals int) (int64, error) {
	//temp test
	//Chuyển đổi địa chỉ ví và địa chỉ token thành kiểu common.Address
	wallet_address := common.HexToAddress(walletAddress)
	contract_address := common.HexToAddress(token_address)
	//fmt.Println("contract_address: ",contract_address)
	//
	// Load the contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		log.Error("Failed to load the contract ABI: "+err.Error(), "GetBalance", walletAddress)
		return 0, err
	}

	// Prepare data for the contract call
	data, err := parsedABI.Pack("balanceOf", wallet_address)
	if err != nil {
		log.Error("Failed to pack the data: "+err.Error(), "GetBalance", walletAddress)
		return 0, err
	}

	// Create a call message to the contract
	callMsg := ethereum.CallMsg{
		To:   &contract_address,
		Data: data,
	}

	// Call the contract
	result, err := n.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		log.Error("Failed to call the contract: "+err.Error(), "GetBalance", walletAddress)
		return 0, err
	}
	if len(result) == 0 {
		log.Error("No result returned from the contract", "GetBalance", walletAddress)
		return 0, nil
	}
	// Parse the result
	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		log.Error("Failed to unpack the result: "+err.Error(), "GetBalance", walletAddress)
		return 0, err
	}
	//decimals := new(big.Int).SetBytes(result).Uint64()
	//fmt.Println("decimals",decimals)
	if decimals == 0 {
		return balance.Int64(), nil
	}
	balance_dec := balance.Int64() / int64(decimals)
	return balance_dec, nil
}
