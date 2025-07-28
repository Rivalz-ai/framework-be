package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	redis "github.com/redis/go-redis/v9"
)

const (
	EtherscanAPIKey  = "66E7UH166981Y8G21SBTT4E1MMED6VCH1U"
	EtherscanBaseURL = "https://api.etherscan.io/v2/api"
)

type EtherscanClient struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	cache      *redis.Client
}

type GasOracleResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  struct {
		LastBlock       string `json:"LastBlock"`
		SafeGasPrice    string `json:"SafeGasPrice"`
		ProposeGasPrice string `json:"ProposeGasPrice"`
		FastGasPrice    string `json:"FastGasPrice"`
		SuggestBaseFee  string `json:"suggestBaseFee"`
		GasUsedRatio    string `json:"gasUsedRatio"`
	} `json:"result"`
}

type GasPrices struct {
	LastBlock       string `json:"last_block"`
	SafeGasPrice    string `json:"safe_gas_price"`
	ProposeGasPrice string `json:"propose_gas_price"`
	FastGasPrice    string `json:"fast_gas_price"`
	SuggestBaseFee  string `json:"suggest_base_fee"`
	GasUsedRatio    string `json:"gas_used_ratio"`
}

var etherscanClient *EtherscanClient

// NewEtherscanClient creates a new Etherscan client instance
func NewEtherscanClient(cacheClient *redis.Client) *EtherscanClient {
	if etherscanClient != nil {
		return etherscanClient
	}

	etherscanClient = &EtherscanClient{
		apiKey:  EtherscanAPIKey,
		baseURL: EtherscanBaseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		cache: cacheClient,
	}
	return etherscanClient
}

// GetGasPrice fetches current gas prices from Etherscan API
func (c *EtherscanClient) GetGasPrice(ctx context.Context, chainID int) (*GasPrices, error) {

	if c.cache == nil {
		return nil, fmt.Errorf("cache client is nil")
	}

	keyCache := "mkp_gas_price_" + strconv.Itoa(chainID)

	// get from cache
	rs := c.cache.Get(ctx, keyCache)
	if rs.Err() == nil {
		var gasPrices *GasPrices
		err := json.Unmarshal([]byte(rs.Val()), &gasPrices)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		return gasPrices, nil
	}

	// get from etherscan
	url := fmt.Sprintf("%s?chainid=%d&module=gastracker&action=gasoracle&apikey=%s",
		c.baseURL, chainID, c.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var gasResponse GasOracleResponse
	if err := json.Unmarshal(body, &gasResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if gasResponse.Status != "1" {
		return nil, fmt.Errorf("API returned error: %s", gasResponse.Message)
	}

	gasPrices := &GasPrices{
		LastBlock:       gasResponse.Result.LastBlock,
		SafeGasPrice:    gasResponse.Result.SafeGasPrice,
		ProposeGasPrice: gasResponse.Result.ProposeGasPrice,
		FastGasPrice:    gasResponse.Result.FastGasPrice,
		SuggestBaseFee:  gasResponse.Result.SuggestBaseFee,
		GasUsedRatio:    gasResponse.Result.GasUsedRatio,
	}

	json, err := json.Marshal(gasPrices)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	c.cache.Set(ctx, keyCache, string(json), time.Second*3)

	return gasPrices, nil
}
