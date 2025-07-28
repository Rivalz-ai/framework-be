package service

import (
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	//"fmt"
	//"github.com/Rivalz-ai/framework-be/modules/node/dto"
	"context"
	//"fmt"
	"math/big"

	abipag "github.com/Rivalz-ai/framework-be/define/abi"
)

// nodeSV.CallMethod(ctx,"tokenOfOwnerByIndex",wallet_address,sv.server.ExtendConfig.Pag.ContractAddress,pagABI.PAG_ABI,index)
func (n *Node) TokenOfOwnerByIndex(ctx context.Context, walletAddress, token_address string, i int) (int64, error) {

	//Chuyển đổi địa chỉ ví và địa chỉ token thành kiểu common.Address
	wallet_address := common.HexToAddress(walletAddress)
	contract_address := common.HexToAddress(token_address)
	//fmt.Println("contract_address: ",contract_address)
	//
	// Load the contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(abipag.PAG_ABI))
	if err != nil {
		log.Error("Failed to load the contract ABI: "+err.Error(), "TokenOfOwnerByIndex", walletAddress)
		return 0, err
	}
	index := big.NewInt(int64(i)) // Thay 0 bằng index bạn muốn lấy
	// Prepare data for the contract call
	data, err := parsedABI.Pack("tokenOfOwnerByIndex", wallet_address, index)
	if err != nil {
		log.Error("Failed to pack the data: "+err.Error(), "TokenOfOwnerByIndex", walletAddress)
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
		log.Error("Failed to call the contract: "+err.Error(), "TokenOfOwnerByIndex", walletAddress)
		return 0, err
	}
	if len(result) == 0 {
		log.Error("No result returned from the contract", "TokenOfOwnerByIndex", walletAddress)
		return 0, nil
	}
	// Parse the result
	var data_resp *big.Int
	err = parsedABI.UnpackIntoInterface(&data_resp, "tokenOfOwnerByIndex", result)
	if err != nil {
		log.Error("Failed to unpack the result: "+err.Error(), "TokenOfOwnerByIndex", walletAddress)
		return 0, err
	}
	//decimals := new(big.Int).SetBytes(result).Uint64()
	//fmt.Println("decimals",decimals)
	//fmt.Println("balance origin: ",balance)
	return data_resp.Int64(), nil
}
