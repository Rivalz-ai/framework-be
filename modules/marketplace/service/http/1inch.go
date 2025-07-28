package http

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
)

// QuoteResponse represents the structure of the JSON response from the 1inch quote endpoint.
type QuoteResponse struct {
	DstAmount string `json:"dstAmount"`
}

func Get1inchTokenPrice(tokenAddress string, decimals int) (float64, error) {
	chainId := 8453
	if os.Getenv("ENV") == "dev" {
		chainId = 84532
	}

	// Replace with the actual contract addresses for BNB and USDC on the correct network.
	// For Binance Smart Chain (BSC):
	bnbAddress := tokenAddress
	usdcAddress := "0x8ac76a512cc950d982f0168e5175163476bbaa5c"

	amountInt := pow(10, decimals)
	amount := strconv.Itoa(int(amountInt))

	url := fmt.Sprintf("https://api.1inch.dev/swap/v6.0/%d/quote?src=%s&dst=%s&amount=%s",
		chainId, bnbAddress, usdcAddress, amount)

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0, err
	}

	// You might need to include an API key in the header if required by the 1inch API in the future.
	// req.Header.Add("Authorization", "Bearer YOUR_API_KEY")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d, body: %s\n", resp.StatusCode, string(body))
		return 0, err
	}

	var quoteResponse QuoteResponse
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return 0, err
	}

	// Calculate the price of 1 BNB in USDC
	bnbAmount := float64(1)
	usdcAmount, err := weiToDecimal(quoteResponse.DstAmount, 6) // USDC has 6 decimals
	if err != nil {
		fmt.Println("Error converting USDC amount:", err)
		return 0, err
	}
	price := usdcAmount / bnbAmount

	fmt.Printf("Approximate price of 1 BNB in USDC: %f\n", price)
	return price, nil
}

// weiToDecimal converts a value in wei (smallest unit) to a decimal representation.
func weiToDecimal(wei string, decimals int) (float64, error) {
	intValue, success := new(big.Int).SetString(wei, 10)
	if !success {
		return 0, fmt.Errorf("failed to parse wei value: %s", wei)
	}

	decimalValue := new(big.Float).SetInt(intValue)
	divisor := new(big.Float).SetFloat64(float64(pow(10, decimals)))
	decimalValue.Quo(decimalValue, divisor)

	floatValue, _ := decimalValue.Float64()
	return floatValue, nil
}

// Helper function for power calculation
func pow(base, exponent int) int64 {
	result := int64(1)
	for i := 0; i < exponent; i++ {
		result *= int64(base)
	}
	return result
}

func GenerateCallableData(srcToken, dstToken, fromWallet, originWallet string, amount string) (*dto.SwapResponse, error) {
	chainId := 8453
	if os.Getenv("ENV") == "dev" {
		chainId = 84532
	}
	method := "GET"
	srcToken = "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913" // usdc
	if os.Getenv("ENV") == "dev" {
		srcToken = "0x081827b8C3Aa05287b5aA2bC3051fbE638F33152" // usdc
	}
	apiUrl := fmt.Sprintf("https://api.1inch.dev/swap/v6.0/%d/swap?src=%s&dst=%s&amount=%s&from=%s&origin=%s&slippage=0", chainId, srcToken, dstToken, amount, fromWallet, originWallet)

	req, err := http.NewRequest(method, apiUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer undefined")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var swapResponse dto.SwapResponse
	err = json.Unmarshal(body, &swapResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return &swapResponse, nil
}
