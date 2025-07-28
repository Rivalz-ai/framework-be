package service

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	marketplaceDto "github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	httpClient "github.com/Rivalz-ai/framework-be/modules/marketplace/service/http"
	"github.com/Rivalz-ai/framework-be/modules/node/dto"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	//"fmt"
	"bytes"
	"io"
	"math/big"
	"net/http"
)

/*
	type Node struct {
		NodeUrls string
		NodeSecrets string
		Client *ethclient.Client
		server *server.Server
	}
*/
type Node struct {
	URL     string
	Secret  string
	Weight  int
	Client  *ethclient.Client
	Key     string
	ChainID int
}

func (n *Node) GetChainID() (int, error) {
	if n.ChainID != 0 {
		return n.ChainID, nil
	}

	chainID, err := n.Client.NetworkID(context.Background())
	if err != nil {
		return 0, err
	}
	return int(chainID.Int64()), nil
}

// nodeSV.CallMethod(ctx,"tokenOfOwnerByIndex",wallet_address,sv.server.ExtendConfig.Pag.ContractAddress,pagABI.PAG_ABI,index)
func (n *Node) CallMethod(ctx context.Context, method, walletAddress, token_address, ABI string, args ...interface{}) (interface{}, error) {
	//temp test
	contractAddr := token_address
	//Chuyển đổi địa chỉ ví và địa chỉ token thành kiểu common.Address
	wallet_address := common.HexToAddress(walletAddress)
	contract_address := common.HexToAddress(contractAddr)
	//fmt.Println("contract_address: ",contract_address)
	//
	// Load the contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		log.Error("Failed to load the contract ABI: "+err.Error(), "CallMethod", method+":"+walletAddress)
		return 0, err
	}
	index := big.NewInt(0) // Thay 0 bằng index bạn muốn lấy
	// Prepare data for the contract call
	data, err := parsedABI.Pack(method, wallet_address, index)
	if err != nil {
		log.Error("Failed to pack the data: "+err.Error(), "CallMethod", method+":"+walletAddress)
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
		log.Error("Failed to call the contract: "+err.Error(), "CallMethod", method+":"+walletAddress)
		return 0, err
	}
	if len(result) == 0 {
		log.Error("No result returned from the contract", "CallMethod", method+":"+walletAddress)
		return 0, nil
	}
	// Parse the result
	var data_resp interface{}
	err = parsedABI.UnpackIntoInterface(&data_resp, method, result)
	if err != nil {
		log.Error("Failed to unpack the result: "+err.Error(), "CallMethod", method+":"+walletAddress)
		return 0, err
	}
	//decimals := new(big.Int).SetBytes(result).Uint64()
	//fmt.Println("decimals",decimals)
	//fmt.Println("balance origin: ",balance)
	return data_resp, nil
}

// GetRewardTokenInfo lấy thông tin token (Decimals + Name)
func (n *Node) GetRewardTokenInfo(ctx context.Context, tokenAddress, ABI string) (*dto.TokenInfo, error) {
	contractAddress := common.HexToAddress(tokenAddress)
	var logData = utils.Dictionary()
	logData["token_address"] = tokenAddress
	// Load ABI
	parsedABI, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		log.Error("Failed to load the contract ABI: "+err.Error(), "GetRewardTokenInfo", logData)
		return nil, err
	}

	// Call decimals()
	data, err := parsedABI.Pack("decimals")
	if err != nil {
		log.Error("Failed to pack decimals(): "+err.Error(), "GetRewardTokenInfo", logData)
		return nil, fmt.Errorf("failed to pack decimals(): %v", err)
	}
	result, err := n.Client.CallContract(ctx, ethereum.CallMsg{To: &contractAddress, Data: data}, nil)
	if err != nil {
		log.Error("Failed to call decimals(): "+err.Error(), "GetRewardTokenInfo", logData)
		return nil, fmt.Errorf("failed to call decimals(): %v", err)
	}

	var decimals uint8
	err = parsedABI.UnpackIntoInterface(&decimals, "decimals", result)
	if err != nil {
		log.Error("Failed to unpack decimals: "+err.Error(), "GetRewardTokenInfo", logData)
		return nil, fmt.Errorf("failed to unpack decimals: %v", err)
	}

	// Call name()
	data, err = parsedABI.Pack("name")
	if err != nil {
		return nil, fmt.Errorf("failed to pack name(): %v", err)
	}
	result, err = n.Client.CallContract(ctx, ethereum.CallMsg{To: &contractAddress, Data: data}, nil)
	if err != nil {
		log.Error("Failed to call name(): "+err.Error(), "GetRewardTokenInfo", logData)
		return nil, fmt.Errorf("failed to call name(): %v", err)
	}

	var name string
	err = parsedABI.UnpackIntoInterface(&name, "name", result)
	if err != nil {
		log.Error("Failed to unpack name: "+err.Error(), "GetRewardTokenInfo", logData)
		return nil, fmt.Errorf("failed to unpack name: %v", err)
	}

	return &dto.TokenInfo{
		RewardAddress: tokenAddress,
		Decimals:      decimals,
		Name:          name,
	}, nil
}

type Block struct {
	Number          string        `json:"number"`
	Hash            string        `json:"hash"`
	ParentHash      string        `json:"parentHash"`
	Nonce           string        `json:"nonce"`
	Timestamp       string        `json:"timestamp"`
	Difficulty      string        `json:"difficulty"`
	GasLimit        string        `json:"gasLimit"`
	GasUsed         string        `json:"gasUsed"`
	Miner           string        `json:"miner"`
	Transactions    []Transaction `json:"transactions"`
	BaseFeePerGas   string        `json:"baseFeePerGas"`
	Size            string        `json:"size"`
	TotalDifficulty string        `json:"totalDifficulty"`
}

// Transaction represents an Ethereum transaction
type Transaction struct {
	Hash                 string `json:"hash"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	Value                string `json:"value"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	Nonce                string `json:"nonce"`
	BlockHash            string `json:"blockHash"`
	BlockNumber          string `json:"blockNumber"`
	TransactionIndex     string `json:"transactionIndex"`
	Input                string `json:"input"`
	Type                 string `json:"type"`
	MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
}

type RPCRequest struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// RPCResponse represents a JSON-RPC response
type RPCResponse struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (n *Node) makeRPCCall(method string, params []interface{}) (*RPCResponse, error) {
	request := RPCRequest{
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", n.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		httpReq.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rpcResp RPCResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return nil, err
	}

	if rpcResp.Error != nil {
		return nil, fmt.Errorf("RPC error: %s (code: %d)", rpcResp.Error.Message, rpcResp.Error.Code)
	}

	return &rpcResp, nil
}

func (n *Node) GetBlockByNumber(ctx context.Context, blockNumber *big.Int) (*Block, error) {
	var blockParam interface{}
	if blockNumber == nil {
		blockParam = "latest"
	} else {
		blockParam = fmt.Sprintf("0x%x", blockNumber)
	}

	response, err := n.makeRPCCall("eth_getBlockByNumber", []interface{}{blockParam, true})
	if err != nil {
		return nil, err
	}

	var block Block
	if err := json.Unmarshal(response.Result, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

func (n *Node) BroadcastTransaction(ctx context.Context, fromWallet string, swapResponse *marketplaceDto.OkxSwapResponse, privateKeyHex string) (string, error) {
	// --- Step 2: Connect to the blockchain node ---
	ethClient := n.Client

	// client, err := ethclient.Dial("https://rpc.ankr.com/base/dc3359a3d6c4f6866d0e59e41b886d8806cba7197232edf7412c79644595b948")
	// if err != nil {
	// 	panic(err)
	// }
	// ethClient := client
	start := time.Now()

	// --- Step 3: Get the current nonce for your account ---
	nonce, err := ethClient.PendingNonceAt(ctx, common.HexToAddress(fromWallet))
	if err != nil {
		fmt.Println("Failed to get nonce:", err)
		return "", err
	}

	chainID, err := ethClient.NetworkID(ctx)
	if err != nil {
		fmt.Println("Failed to get chain ID:", err)
		return "", err
	}

	// --- Step 4: Parse the transaction details from the 1inch response ---
	gasPrice := big.NewInt(0)
	if swapResponse.Tx.GasPrice != "" {
		gasPrice, _ = new(big.Int).SetString(swapResponse.Tx.GasPrice, 10)
	}

	gasPriceTmp, err := n.Client.SuggestGasPrice(ctx)
	if err == nil {
		gasPriceTmp = gasPriceTmp.Mul(gasPriceTmp, big.NewInt(115)) // 10% more to speed up
		gasPriceTmp = gasPriceTmp.Div(gasPriceTmp, big.NewInt(100))
		if gasPriceTmp.Cmp(gasPrice) > 0 {
			gasPrice = gasPriceTmp
		}
	}

	// get from etherscan
	etherClient := httpClient.NewEtherscanClient(nil)
	gasPriceEther, err := etherClient.GetGasPrice(ctx, int(chainID.Int64()))
	if err == nil {
		gasPriceFl, ok := new(big.Float).SetString(gasPriceEther.FastGasPrice)
		if ok {
			gasPriceFl.Mul(gasPriceFl, big.NewFloat(1e9))
			gasPriceTmp, _ := gasPriceFl.Int(nil)
			if gasPriceTmp.Cmp(gasPrice) > 0 {
				gasPrice = gasPriceTmp
			}
		}
	}

	gasLimit := uint64(1000000)
	if swapResponse.Tx.Gas != "" {
		gasLimit, _ = strconv.ParseUint(swapResponse.Tx.Gas, 10, 64)
		gasLimit = gasLimit * 105 / 100
	}
	toAddress := common.HexToAddress(swapResponse.Tx.To)
	value := new(big.Int) // Often zero for ERC-20 swaps
	if swapResponse.Tx.Value != "" {
		value, _ = new(big.Int).SetString(swapResponse.Tx.Value, 10)
	}
	data, err := hex.DecodeString(swapResponse.Tx.Data[2:])
	if err != nil {
		fmt.Println("Failed to decode transaction data:", err)
		return "", err
	}

	gasTipCap, err := ethClient.SuggestGasTipCap(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get gas tip cap: %v", err)
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		To:        &toAddress,
		Value:     value,
		Gas:       gasLimit,
		GasFeeCap: gasPrice,
		GasTipCap: gasTipCap,
		Data:      data,
	})

	// --- Step 5: Create the transaction object ---
	// tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// --- Step 6: Sign the transaction with your private key ---
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		fmt.Println("Failed to decode private key:", err)
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		fmt.Println("Failed to sign transaction:", err)
		return "", err
	}

	// --- Step 7: Broadcast the signed transaction to the network ---
	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("Failed to send transaction:", err)
		return "", err
	}

	// bind receipt
	receipt, err := bind.WaitMined(ctx, ethClient, signedTx)
	if err != nil {
		return "", err
	}

	log.Info("BROADCAST_TRANSACTION", "TIME_PROCESS", map[string]interface{}{
		"time_process": time.Since(start).Seconds(),
		"tx_hash":      signedTx.Hash().Hex(),
		"node":         n.URL,
	})

	if receipt.Status == 1 {
	} else {
		return "", errors.New("transaction failed")
	}

	return signedTx.Hash().Hex(), nil
}

func (n *Node) BroadcastSignedTransaction(ctx context.Context, signedTransaction string) (string, error) {
	ethClient := n.Client

	signedTxBytes, err := hex.DecodeString(strings.TrimPrefix(signedTransaction, "0x"))
	if err != nil {
		return "", fmt.Errorf("failed to decode signed transaction hex: %v", err)
	}

	var tx types.Transaction
	if err := tx.UnmarshalBinary(signedTxBytes); err != nil {
		return "", fmt.Errorf("failed to unmarshal signed transaction: %v", err)
	}

	err = ethClient.SendTransaction(ctx, &tx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	receipt, err := bind.WaitMined(ctx, ethClient, &tx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}
	if receipt.Status == 1 {
	} else {
		return "", errors.New("transaction failed")
	}

	return tx.Hash().Hex(), nil
}
