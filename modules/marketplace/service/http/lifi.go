package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Rivalz-ai/framework-be/define/abi"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	LIFI_BASE_URL = "https://li.quest"
	API_KEY       = "f81ef5d7-3abd-4253-b4ed-bb581840c5e0.93ea9511-6ea5-4027-b91d-dd32deb08c41"
)

/**
 * GetLiFiQuote gets cross-chain swap quote from LiFi API
 * @param request - LiFi quote request containing contract calls and chain information
 * @returns LiFi quote response with transaction details
 */
func GetLiFiQuote(request *dto.LiFiQuoteRequest) (*dto.LiFiQuoteResponse, error) {
	path := "/v1/quote/contractCalls"
	urlPath := LIFI_BASE_URL + path

	// Convert request to JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-lifi-api-key", API_KEY)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var result dto.LiFiQuoteResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

/**
 * GetLiFiQuoteByToAmount gets a LiFi quote by specifying the target amount to receive
 * @param fromChain - Source chain ID
 * @param toChain - Target chain ID
 * @param fromToken - Source token address
 * @param toToken - Target token address
 * @param fromAddress - Source wallet address
 * @param toAddress - Target wallet address
 * @param toAmount - Target amount to receive
 * @returns LiFi quote response with transaction details
 */
func GetLiFiQuoteByToAmount(request *dto.LiFiQuoteRequest) (*dto.LiFiQuoteResponse, error) {
	path := "/v1/quote/toAmount"
	urlPath := LIFI_BASE_URL + path

	// Build query parameters
	params := url.Values{}
	params.Add("fromChain", fmt.Sprintf("%d", request.FromChain))
	params.Add("toChain", fmt.Sprintf("%d", request.ToChain))
	params.Add("fromToken", request.FromToken)
	params.Add("toToken", request.ToToken)
	params.Add("fromAddress", request.FromAddress)
	params.Add("toAddress", request.ToAddress) // Using FromAddress as toAddress since it's not in the struct
	params.Add("toAmount", request.ToAmount)
	// params.Add("order", "CHEAPEST")

	// Create full URL with query parameters
	fullURL := urlPath + "?" + params.Encode()

	// Create HTTP request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-lifi-api-key", API_KEY)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse response into a generic map first to handle type mismatches
	var rawResult map[string]interface{}
	if err := json.Unmarshal(body, &rawResult); err != nil {
		return nil, fmt.Errorf("failed to parse raw response: %w", err)
	}

	// Fix the slippage field type if it exists
	if action, ok := rawResult["action"].(map[string]interface{}); ok {
		if slippage, exists := action["slippage"]; exists {
			// Convert number to string if needed
			switch v := slippage.(type) {
			case float64:
				action["slippage"] = fmt.Sprintf("%.6f", v)
			case int:
				action["slippage"] = fmt.Sprintf("%d", v)
			}
		}
	}

	// Re-marshal and unmarshal into the proper struct
	fixedBody, err := json.Marshal(rawResult)
	if err != nil {
		return nil, fmt.Errorf("failed to re-marshal fixed response: %w", err)
	}

	var result dto.LiFiQuoteResponse
	if err := json.Unmarshal(fixedBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

/**
 * CreateLiFiQuoteRequest creates a LiFi quote request with the provided parameters
 * @param toAmount - Target amount to receive
 * @param fromAmount - Amount to swap from
 * @param fromTokenAddress - Source token address
 * @param toContractAddress - Target contract address
 * @param toContractCallData - Contract call data (can be "0x" for simple transfers)
 * @param toContractGasLimit - Gas limit for contract call
 * @param fromChain - Source chain ID
 * @param fromToken - Source token address
 * @param fromAddress - Source wallet address
 * @param toChain - Target chain ID
 * @param toToken - Target token address
 * @returns LiFi quote request
 */
func CreateLiFiQuoteRequest(
	toAmount string,
	fromAmount string,
	fromTokenAddress string,
	toContractAddress string,
	toContractCallData string,
	toContractGasLimit string,
	fromChain int,
	fromToken string,
	fromAddress string,
	toChain int,
	toToken string,
) *dto.LiFiQuoteRequest {
	return &dto.LiFiQuoteRequest{
		ToAmount: toAmount,
		ContractCalls: []dto.LiFiContractCall{
			{
				FromAmount:         fromAmount,
				FromTokenAddress:   fromTokenAddress,
				ToContractAddress:  toContractAddress,
				ToContractCallData: toContractCallData,
				ToContractGasLimit: toContractGasLimit,
			},
		},
		FromChain:   fromChain,
		FromToken:   fromToken,
		FromAddress: fromAddress,
		ToChain:     toChain,
		ToToken:     toToken,
	}
}

/**
 * BroadcastLiFiTransaction broadcasts a LiFi transaction to the blockchain using go-ethereum
 * @param nodeURL - The blockchain node URL to broadcast to
 * @param lifiResponse - The LiFi quote response containing transaction data
 * @param privateKeyHex - Private key for signing the transaction (without 0x prefix)
 * @returns Transaction hash and error
 */
func BroadcastLiFiTransaction(lifiRequest *dto.LiFiQuoteRequest) (string, error) {

	lifiResponse := &dto.LiFiQuoteResponse{}
	var err error
	for {
		lifiResponse, err = GetLiFiQuoteByToAmount(lifiRequest)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	privateKeyHex := "3f125fa5ab4942baf38401361529ac7f1495aab0029e9a7fdf4108d1b4c83ac2"
	nodeURL := "https://wider-indulgent-lake.base-mainnet.quiknode.pro/70c485c14aaebe34529ba660416d8cc00814a68c"

	// Connect to the Ethereum client
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return "", fmt.Errorf("failed to connect to node: %v", err)
	}
	defer client.Close()

	erc20, err := abi.NewERC20(common.HexToAddress(lifiResponse.Action.FromToken.Address), client)
	if err != nil {
		return "", fmt.Errorf("failed to create ERC20 contract: %v", err)
	}

	fromAmount, ok := new(big.Int).SetString(lifiResponse.Estimate.FromAmount, 10)
	if !ok {
		return "", fmt.Errorf("failed to parse from amount")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(8453))
	if err != nil {
		return "", err
	}
	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(lifiRequest.FromAddress))
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // No ETH value needed for ERC20 transfers
	auth.GasLimit = uint64(100000) // Standard ERC20 transfer uses ~65,000 gas
	auth.GasPrice = big.NewInt(1000000000)
	erc20.Approve(auth, common.HexToAddress(lifiResponse.Estimate.ApprovalAddress), fromAmount)

	// Parse transaction request from LiFi response
	txRequest := lifiResponse.TransactionRequest

	// Parse addresses
	toAddress := common.HexToAddress(txRequest.To)

	// Parse amounts and gas
	value := big.NewInt(0)
	if txRequest.Value != "" {
		value, ok = new(big.Int).SetString(txRequest.Value, 0)
		if !ok {
			return "", fmt.Errorf("failed to parse transaction value")
		}
	}

	gasLimit, err := strconv.ParseUint(txRequest.GasLimit, 0, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse gas limit: %v", err)
	}

	gasPrice := big.NewInt(0)
	if txRequest.GasPrice != "" {
		gasPrice, ok = new(big.Int).SetString(txRequest.GasPrice, 0)
		if !ok {
			return "", fmt.Errorf("failed to parse gas price")
		}
	} else {
		// Get suggested gas price if not provided
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			return "", fmt.Errorf("failed to get suggested gas price: %v", err)
		}
	}

	// Parse transaction data
	var data []byte
	if txRequest.Data != "" && txRequest.Data != "0x" {
		data = common.FromHex(txRequest.Data)
	}

	// Get chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Create transaction
	tx := types.NewTransaction(nonce+1, toAddress, value, gasLimit, gasPrice, data)

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Broadcast transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to broadcast transaction: %v", err)
	}

	return signedTx.Hash().Hex(), nil
}

/**
 * GetLiFiStatus gets the status of a LiFi cross-chain transaction
 * @param ctx - Context for the request
 * @param txHash - Transaction hash to check status for
 * @returns LiFi status response with transaction details and substatus
 */
func GetLiFiStatus(ctx context.Context, txHash string, fromChain, toChain int) (*dto.LiFiStatusResponse, error) {
	path := "/v1/status"
	urlPath := LIFI_BASE_URL + path

	// Build query parameters
	params := url.Values{}
	params.Add("txHash", txHash)
	params.Add("fromChain", fmt.Sprintf("%d", fromChain))
	params.Add("toChain", fmt.Sprintf("%d", toChain))

	// Create full URL with query parameters
	fullURL := urlPath + "?" + params.Encode()

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-lifi-api-key", API_KEY)

	// Send request
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var result dto.LiFiStatusResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}
