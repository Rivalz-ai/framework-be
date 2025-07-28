package http

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
)

var (
	APIKEY     = "c469f3b6-a856-4887-bd24-7ef7fa67e791"
	PASSPHRASE = "132acbDFE!@#"
	SECRET     = "3727EDFE83EA685CD320B684EC905ECA"
	BASE_URL   = "https://www.okx.com/"
	CHAIN_ID   = "8453"
)

type OkxConfig struct {
	ApiKey     string
	Passphrase string
	Secret     string
}

var okxConfigs = []OkxConfig{
	{
		ApiKey:     "c469f3b6-a856-4887-bd24-7ef7fa67e791",
		Passphrase: "132acbDFE!@#",
		Secret:     "3727EDFE83EA685CD320B684EC905ECA",
	},
	{
		ApiKey:     "b7231427-2a7b-419e-a8b2-86265ffeb858",
		Passphrase: "Z@MB5eUdXw2xWpk",
		Secret:     "836CA615133A27A6D8D96F0B3F55448D",
	},
	{
		ApiKey:     "e8d7bdec-aaf5-4668-a07b-fbcd7cd5d936",
		Passphrase: "e2!U8VtCzaYBAe7",
		Secret:     "E0AD09113ACD74440B43A9D74183DBAD",
	},
	{
		ApiKey:     "5d1bf78e-c6ed-45fc-b24e-aeadd4e6db11",
		Passphrase: "db!LC-tvKqdY4K-",
		Secret:     "07EF3E60F935E1C1D28CC86FE083761B",
	},
}

var UsdcMapping map[int]string = map[int]string{
	int(define.BASE):     "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
	int(define.ETHEREUM): "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
	int(define.BSC):      "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d",
	int(define.SOLANA):   "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
}

func GetOkxTokenPrice(tokenAddress string, decimals int, chainId int) (float64, float64, error) {
	// Define the base URL for OKX API
	urlPath := BASE_URL + "/api/v5/dex/aggregator/quote"

	// Calculate the amount based on decimals
	// amountInt := pow(10, decimals)
	// amount := strconv.Itoa(int(amountInt))

	decimalMultiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	amountInt := new(big.Int).Mul(decimalMultiplier, big.NewInt(1))
	amount := amountInt.String()

	// Set up query parameters
	values := url.Values{}
	// values.Add("chainIndex", strconv.Itoa(chainId))
	values.Add("fromTokenAddress", tokenAddress)
	values.Add("toTokenAddress", UsdcMapping[chainId]) // USDC on Base
	values.Add("amount", amount)

	if chainId == int(define.SOLANA) {
		chainId = 501
		values.Add("chainIndex", strconv.Itoa(chainId))
	} else {
		values.Add("chainIndex", strconv.Itoa(chainId))
	}

	// values.Add("slippage", "0.5")
	queryString := "?" + values.Encode()

	// Prepare authentication
	// timestamp := time.Now().UTC().Format(time.RFC3339)
	// requestPath := "/api/v5/dex/aggregator/quote"
	// headers := getHeaders(timestamp, "GET", requestPath, queryString)
	var errCapture error
	for _, okxConfig := range okxConfigs {
		headers := GenHeader("GET", okxConfig.ApiKey, okxConfig.Secret, "/api/v5/dex/aggregator/quote"+queryString, okxConfig.Passphrase)

		// Create request
		req, err := http.NewRequest("GET", urlPath+queryString, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			errCapture = err
			continue
		}

		// Add headers
		for k, v := range headers {
			req.Header.Add(k, v)
		}

		// Send request
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			errCapture = err
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			errCapture = err
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Request failed with status code: %d, body: %s\n", resp.StatusCode, string(body))
			errCapture = fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
			continue
		}

		// Parse the response
		var response map[string]interface{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error unmarshalling JSON	:", err)
			errCapture = err
			continue
		}

		// Extract the price information from the response
		data, ok := response["data"].([]interface{})
		if !ok || len(data) == 0 {
			errCapture = fmt.Errorf("invalid response format")
			continue
		}

		quoteData, ok := data[0].(map[string]interface{})
		if !ok {
			errCapture = fmt.Errorf("invalid quote data format")
			continue
		}

		toTokenAmount, ok := quoteData["toTokenAmount"].(string)
		if !ok {
			errCapture = fmt.Errorf("invalid toTokenAmount format")
			continue
		}

		// Convert the amount to a decimal
		toAmount, err := strconv.ParseFloat(toTokenAmount, 64)
		if err != nil {
			errCapture = fmt.Errorf("error parsing toTokenAmount: %w", err)
			continue
		}

		// Calculate the price (1 token = X USDC)
		decimalDiv := 1e6
		if chainId == 56 {
			decimalDiv = 1e18
		}
		price := toAmount / decimalDiv // add 1% to make sure the price is enough

		// get price impact
		priceImpact, ok := quoteData["priceImpactPercentage"].(string)
		if !ok {
			errCapture = fmt.Errorf("invalid priceImpactPercentage format")
			continue
		}

		priceImpactFloat, err := strconv.ParseFloat(priceImpact, 64)
		if err != nil {
			errCapture = fmt.Errorf("error parsing priceImpact: %w", err)
			continue
		}
		priceImpactFloat = math.Round(priceImpactFloat*1000) / 1000 // Round to 3 decimal places
		priceImpactFloat = math.Abs(priceImpactFloat)

		return price, priceImpactFloat, nil
	}
	if errCapture != nil {
		return 0, 0, errCapture
	}
	return 0, 0, fmt.Errorf("no price found")
}

/**
 * GetSwapQuote gets swap quote from DEX API
 * @param fromTokenAddress - Source token address
 * @param toTokenAddress - Destination token address
 * @param amount - Amount to swap
 * @param slippage - Maximum slippage (e.g., "0.5" for 0.5%)
 * @returns Swap quote
 */
func GetOkxSwap(fromTokenAddress, toTokenAddress, amount string, slippage string, userWallet string) (*dto.OkxSwapResponse, error) {
	slippageValue := "0.05"
	if slippage != "" {
		slippageValue = slippage
	}

	path := "/api/v5/dex/aggregator/swap"
	urlPath := BASE_URL + path

	// Create params map
	params := map[string]string{
		"chainIndex":        CHAIN_ID,
		"fromTokenAddress":  fromTokenAddress,
		"toTokenAddress":    toTokenAddress,
		"amount":            amount,
		"slippage":          slippageValue,
		"userWalletAddress": userWallet,
		// "priceImpactProtectionPercentage": slippageValue,
	}

	// Convert params to query string
	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}
	queryString := "?" + values.Encode()

	// Prepare authentication
	// timestamp := time.Now().UTC().Format(time.RFC3339)
	// requestPath := "/api/v5/" + path
	// headers := getHeaders(timestamp, "GET", requestPath, queryString)
	var errCapture error
	for _, okxConfig := range okxConfigs {
		headers := GenHeader("GET", okxConfig.ApiKey, okxConfig.Secret, "/api/v5/dex/aggregator/swap"+queryString, okxConfig.Passphrase)

		// Create request
		req, err := http.NewRequest("GET", urlPath+queryString, nil)
		if err != nil {
			errCapture = fmt.Errorf("failed to create request: %w", err)
			continue
		}

		// Add headers
		for k, v := range headers {
			req.Header.Add(k, v)
		}

		// Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errCapture = fmt.Errorf("failed to send request: %w", err)
			continue
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errCapture = fmt.Errorf("failed to read response: %w", err)
			continue
		}

		// Parse response
		var result struct {
			Code string                `json:"code"`
			Msg  string                `json:"msg"`
			Data []dto.OkxSwapResponse `json:"data"`
		}
		if err := json.Unmarshal(body, &result); err != nil {
			errCapture = fmt.Errorf("failed to parse response: %w", err)
			continue
		}

		// Check response
		if result.Code != "0" {
			errCapture = fmt.Errorf("API Error: %s", result.Msg)
			continue
		}

		if len(result.Data) == 0 {
			errCapture = fmt.Errorf("no data returned")
			continue
		}

		return &result.Data[0], nil
	}
	if errCapture != nil {
		return nil, errCapture
	}
	return nil, fmt.Errorf("no swap data")
}

func GetOkxSwapWithAutoSlippage(fromTokenAddress, toTokenAddress, amount string, userWallet string) (*dto.OkxSwapResponse, error) {
	path := "/api/v5/dex/aggregator/swap"
	urlPath := BASE_URL + path

	// Create params map
	params := map[string]string{
		"chainIndex":        CHAIN_ID,
		"fromTokenAddress":  fromTokenAddress,
		"toTokenAddress":    toTokenAddress,
		"amount":            amount,
		"slippage":          "0.05",
		"userWalletAddress": userWallet,
		"autoSlippage":      "true",
		"maxAutoSlippage":   "0.05",
		// "priceImpactProtectionPercentage": slippageValue,
	}

	// Convert params to query string
	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}
	queryString := "?" + values.Encode()

	// Prepare authentication
	// timestamp := time.Now().UTC().Format(time.RFC3339)
	// requestPath := "/api/v5/" + path
	// headers := getHeaders(timestamp, "GET", requestPath, queryString)
	var errCapture error
	for _, okxConfig := range okxConfigs {
		headers := GenHeader("GET", okxConfig.ApiKey, okxConfig.Secret, "/api/v5/dex/aggregator/swap"+queryString, okxConfig.Passphrase)

		// Create request
		req, err := http.NewRequest("GET", urlPath+queryString, nil)
		if err != nil {
			errCapture = fmt.Errorf("failed to create request: %w", err)
			continue
		}

		// Add headers
		for k, v := range headers {
			req.Header.Add(k, v)
		}

		// Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errCapture = fmt.Errorf("failed to send request: %w", err)
			continue
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errCapture = fmt.Errorf("failed to read response: %w", err)
			continue
		}

		// Parse response
		var result struct {
			Code string                `json:"code"`
			Msg  string                `json:"msg"`
			Data []dto.OkxSwapResponse `json:"data"`
		}
		if err := json.Unmarshal(body, &result); err != nil {
			errCapture = fmt.Errorf("failed to parse response: %w", err)
			continue
		}

		// Check response
		if result.Code != "0" {
			errCapture = fmt.Errorf("API Error: %s", result.Msg)
			continue
		}

		if len(result.Data) == 0 {
			errCapture = fmt.Errorf("no data returned")
			continue
		}

		return &result.Data[0], nil
	}
	if errCapture != nil {
		return nil, errCapture
	}
	return nil, fmt.Errorf("no swap data")
}

func OkxApproveTransaction(toTokenAddress, amount string) (*dto.OkxApproveResponse, error) {
	path := "/api/v5/dex/aggregator/approve-transaction"
	urlPath := BASE_URL + path

	// Create params map
	params := map[string]string{
		"chainIndex":           CHAIN_ID,
		"tokenContractAddress": toTokenAddress,
		"approveAmount":        amount,
	}

	// Convert params to query string
	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}
	queryString := "?" + values.Encode()

	// Prepare authentication
	// timestamp := time.Now().UTC().Format(time.RFC3339)
	// requestPath := "/api/v5/" + path
	// headers := getHeaders(timestamp, "GET", requestPath, queryString)
	var errCapture error
	for _, okxConfig := range okxConfigs {
		headers := GenHeader("GET", okxConfig.ApiKey, okxConfig.Secret, "/api/v5/dex/aggregator/approve-transaction"+queryString, okxConfig.Passphrase)

		// Create request
		req, err := http.NewRequest("GET", urlPath+queryString, nil)
		if err != nil {
			errCapture = fmt.Errorf("failed to create request: %w", err)
			continue
		}

		// Add headers
		for k, v := range headers {
			req.Header.Add(k, v)
		}

		// Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errCapture = fmt.Errorf("failed to send request: %w", err)
			continue
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errCapture = fmt.Errorf("failed to read response: %w", err)
			continue
		}

		// Parse response
		var result struct {
			Code string                   `json:"code"`
			Msg  string                   `json:"msg"`
			Data []dto.OkxApproveResponse `json:"data"`
		}
		if err := json.Unmarshal(body, &result); err != nil {
			errCapture = fmt.Errorf("failed to parse response: %w", err)
			continue
		}

		// Check response
		if result.Code != "0" {
			errCapture = fmt.Errorf("API Error: %s", result.Msg)
			continue
		}

		if len(result.Data) == 0 {
			errCapture = fmt.Errorf("no data returned")
			continue
		}

		return &result.Data[0], nil
	}
	if errCapture != nil {
		return nil, errCapture
	}
	return nil, fmt.Errorf("no approve data")
}

func getHeaders(timestamp, method, requestPath, queryString string) map[string]string {
	// Create the prehash string by concatenating required parts
	preHash := timestamp + method + requestPath + queryString

	// Create a new HMAC-SHA256 hash
	h := hmac.New(sha256.New, []byte(APIKEY))
	h.Write([]byte(preHash))

	// Get the hash and encode it to base64
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// Return headers
	return map[string]string{
		"OK-ACCESS-PROJECT":    "d1b814619e792f428d238fafd8e9435d",
		"OK-ACCESS-KEY":        APIKEY,
		"OK-ACCESS-SIGN":       signature,
		"OK-ACCESS-TIMESTAMP":  timestamp,
		"OK-ACCESS-PASSPHRASE": "Hackintosh123@", // Add your passphrase if required
		"Content-Type":         "application/json",
	}
}

func GenHeader(method, apiKey, secret, path, passPhrase string) map[string]string {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	message := timestamp + method + path
	signature := genSignature(message, secret)
	return map[string]string{
		"OK-ACCESS-PROJECT":    "d1b814619e792f428d238fafd8e9435d",
		"Content-Type":         "application/json",
		"OK-ACCESS-KEY":        apiKey,
		"OK-ACCESS-SIGN":       signature,
		"OK-ACCESS-TIMESTAMP":  timestamp,
		"OK-ACCESS-PASSPHRASE": passPhrase,
	}
}

func genSignature(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func BuildParams(requestPath string, params map[string]string) string {
	urlParams := url.Values{}
	for k := range params {
		urlParams.Add(k, params[k])
	}
	return requestPath + "?" + urlParams.Encode()
}
