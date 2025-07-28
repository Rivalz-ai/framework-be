package service

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/framework/log"
	marketplaceDto "github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	httpClient "github.com/Rivalz-ai/framework-be/modules/marketplace/service/http"
	protjectDto "github.com/Rivalz-ai/framework-be/modules/project/dto"
	httpClientCGK "github.com/Rivalz-ai/framework-be/modules/project/service/httpclient"
	"github.com/Rivalz-ai/framework-be/util"
	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MarketplaceService) helperGetMkpSetting(ctx context.Context) (*marketplaceDto.MkpSetting, error) {
	// get from cache
	var mkpSetting *marketplaceDto.MkpSetting
	rs := s.server.Redis.Client.Get(ctx, "mkp_setting")
	if rs.Err() == nil {
		err := json.Unmarshal([]byte(rs.Val()), &mkpSetting)
		if err != nil {
			return nil, err
		}
		return mkpSetting, nil
	}

	// get from db
	err := s.colMkpSetting.FindOne(ctx, bson.M{}).Decode(&mkpSetting)
	if err != nil {
		return nil, err
	}

	// set to cache
	json, err := json.Marshal(mkpSetting)
	if err != nil {
		return nil, err
	}
	s.server.Redis.Client.Set(ctx, "mkp_setting", string(json), time.Minute*5)

	return mkpSetting, nil
}

func (s *MarketplaceService) helperGetRAgentEstValue(ctx context.Context, ragentId string) (float64, float64, float64, float64, error) {

	keyCache := "mkp_ragent_est_price_" + ragentId

	// get from cache
	rs := s.server.Redis.Client.Get(ctx, keyCache)
	if rs.Err() == nil {
		var mkpRagent *marketplaceDto.MkpRagentEstPriceCache
		err := json.Unmarshal([]byte(rs.Val()), &mkpRagent)
		if err != nil {
			return 0, 0, 0, 0, err
		}
		return mkpRagent.RagentEstPrice, mkpRagent.TokenProjectPrice, mkpRagent.RizTokenPrice, mkpRagent.PriceImpact, nil
	}

	agentId, err := primitive.ObjectIDFromHex(ragentId)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	filter := bson.M{"_id": agentId}
	var mkpRagent *marketplaceDto.MkpRagent
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// for testing
	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "local" {
		mkpRagent.TokenAddress = "0x67543cf0304c19ca62ac95ba82fd4f4b40788dc1"
		s.server.ExtendConfig.RIZAddress = "0x67543cf0304c19ca62ac95ba82fd4f4b40788dc1"
	}

	var price float64
	var priceImpact float64

	// get price of token
	// get from cache first
	price, priceImpact, _ = httpClient.GetOkxTokenPrice(mkpRagent.TokenAddress, mkpRagent.TokenDecimals, mkpRagent.ChainId)

	if mkpRagent.ChainId != int(define.BASE) && price == 0 {
		decimalMultiplier := math.Pow(10, float64(mkpRagent.TokenDecimals))
		toAmount := big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(int64(decimalMultiplier)))
		lifiResponse, err := httpClient.GetLiFiQuoteByToAmount(&marketplaceDto.LiFiQuoteRequest{
			ToAmount:  toAmount.String(),
			FromToken: s.server.ExtendConfig.USDCAddress,
			ToToken:   mkpRagent.TokenAddress,
			FromChain: int(define.BASE),
			ToChain:   mkpRagent.ChainId,
			ContractCalls: []marketplaceDto.LiFiContractCall{
				{
					FromTokenAddress:   s.server.ExtendConfig.USDCAddress,
					ToContractAddress:  mkpRagent.TokenAddress,
					ToContractCallData: "0x",
					ToContractGasLimit: "200000",
				},
			},
			FromAddress: "0xc5815b405F6af3B80e2E2c2A2766E41fC584459b",
		})
		if err == nil {
			feeCosts := 0.0
			for _, feeCost := range lifiResponse.Estimate.FeeCosts {
				feeCostAmount, _ := new(big.Float).SetString(feeCost.AmountUSD)
				feeCostAmountFloat, _ := feeCostAmount.Float64()
				feeCosts += feeCostAmountFloat
			}
			toAmount, _ := new(big.Float).SetString(lifiResponse.Estimate.ToAmountMin)
			priceToToken, _ := toAmount.Quo(toAmount, big.NewFloat(decimalMultiplier)).Float64()
			fromAmount, _ := new(big.Float).SetString(lifiResponse.Estimate.FromAmount)
			fromAmountFloat, _ := fromAmount.Quo(fromAmount, big.NewFloat(1e6)).Float64()
			price = (fromAmountFloat + feeCosts) / priceToToken
			priceImpact, _ = strconv.ParseFloat(lifiResponse.Action.Slippage, 64)
			priceImpact = priceImpact * 100
		} else {
			log.Error("Error get lifi quote: "+err.Error(), "GET_TOKEN_PRICE", map[string]interface{}{
				"ragentId":         ragentId,
				"tokenAddress":     mkpRagent.TokenAddress,
				"chainId":          mkpRagent.ChainId,
				"fromTokenAddress": s.server.ExtendConfig.USDCAddress,
				"toTokenAddress":   mkpRagent.TokenAddress,
				"fromChain":        int(define.BASE),
				"toChain":          mkpRagent.ChainId,
			})
		}
	}

	// Reduce sleep time to improve performance
	time.Sleep(time.Millisecond * 50)

	// get price of riz token
	// priceRizToken, _, _ := httpClient.GetOkxTokenPrice(s.server.ExtendConfig.RIZAddress, 8, int(define.BASE))

	var priceRizToken float64

	// get price of riz token from cache
	cacheKey := "token_price_" + s.server.ExtendConfig.RIZAddress + "_" + strconv.Itoa(int(define.BASE))
	rs = s.server.Redis.Client.Get(ctx, cacheKey)
	if rs.Err() == nil {
		var priceRizTokenTmp float64
		err = json.Unmarshal([]byte(rs.Val()), &priceRizTokenTmp)
		if err == nil {
			priceRizToken = priceRizTokenTmp
		}
	} else {
		priceRizToken, _, _ = httpClient.GetOkxTokenPrice(s.server.ExtendConfig.RIZAddress, 8, int(define.BASE))
		json, err := json.Marshal(priceRizToken)
		if err != nil {
			return 0, 0, 0, 0, err
		}
		s.server.Redis.Client.Set(ctx, cacheKey, string(json), time.Second*10)
	}

	// make sure get highest price to make sure the price is enough to swap
	tokenList := []string{mkpRagent.TokenAddress, s.server.ExtendConfig.RIZAddress}
	if len(tokenList) > 0 {
		tokenInfo, err := httpClientCGK.GetCoinGecKoPrice(s.server.ExtendConfig.CoinGeckoURL, util.GetChainName(mkpRagent.ChainId), s.server.ExtendConfig.CoinGeckoToken, tokenList, map[string]interface{}{})
		if err == nil {
			pPrice := tokenInfo[mkpRagent.TokenAddress]["usd"]
			if pPrice > price {
				price = pPrice
			}
			pPrice = tokenInfo[s.server.ExtendConfig.RIZAddress]["usd"]
			if pPrice > priceRizToken {
				priceRizToken = pPrice
			}
		} else {
			log.Error("Error get coin gecko price: "+err.Error(), "GET_TOKEN_PRICE", map[string]interface{}{
				"ragentId":         ragentId,
				"tokenAddress":     mkpRagent.TokenAddress,
				"chainId":          mkpRagent.ChainId,
				"fromTokenAddress": s.server.ExtendConfig.USDCAddress,
				"toTokenAddress":   mkpRagent.TokenAddress,
				"fromChain":        int(define.BASE),
				"toChain":          mkpRagent.ChainId,
			})
		}
	}

	// get swap slippage
	walletAddress, err := s.helperGetAddressFromPrivateKey(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// amount need = project token price * number of project token * 1e6
	amountUsdcProjectToken := price * mkpRagent.ProjectToken
	amountInFloat := new(big.Float).Mul(big.NewFloat(amountUsdcProjectToken), big.NewFloat(1e6))
	amountDstToken, _ := amountInFloat.Int(nil)
	swapSlippage, err := httpClient.GetOkxSwapWithAutoSlippage(httpClient.UsdcMapping[mkpRagent.ChainId], mkpRagent.TokenAddress, amountDstToken.String(), walletAddress)
	if err == nil {
		slippageUsed, _ := strconv.ParseFloat(swapSlippage.Tx.Slippage, 64)
		priceImpact = slippageUsed * 100
	}

	// if can not get price impact, use slippage rate on setting
	if priceImpact == 0 || priceImpact > 5 {
		mkpSetting, err := s.helperGetMkpSetting(ctx)
		if err != nil {
			return 0, 0, 0, 0, err
		}
		priceImpact = mkpSetting.SlippageRate * 100
		if mkpRagent.ChainId != int(define.BASE) {
			priceImpact = 1
		}
	}

	// set to cache
	rAgentEstPrice := float64(mkpRagent.ProjectToken)*price + float64(mkpRagent.RizToken)*priceRizToken
	mkpRagentEstPriceCache := &marketplaceDto.MkpRagentEstPriceCache{
		TokenProjectPrice: price,
		RizTokenPrice:     priceRizToken,
		RagentEstPrice:    rAgentEstPrice,
		PriceImpact:       priceImpact,
	}
	json, err := json.Marshal(mkpRagentEstPriceCache)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	s.server.Redis.Client.Set(ctx, keyCache, string(json), time.Second*5)

	return rAgentEstPrice * 1.005, price * 1.005, priceRizToken * 1.005, priceImpact, nil
}

func (s *MarketplaceService) helperTokenPrice(ctx context.Context, tokenAddress string, chainId int) (map[string]float64, error) {
	cacheKey := "market_cap_" + tokenAddress
	rs := s.server.Redis.Client.Get(ctx, cacheKey)
	if rs.Err() == nil {
		tokenInfo, err := rs.Result()
		if err != nil {
			return nil, err
		}
		var tokenInfoRs map[string]float64
		err = json.Unmarshal([]byte(tokenInfo), &tokenInfoRs)
		if err != nil {
			return nil, err
		}
		return tokenInfoRs, nil
	}

	tokenInfo, err := httpClientCGK.GetCoinGecKoPrice(s.server.ExtendConfig.CoinGeckoURL, util.GetChainName(chainId), s.server.ExtendConfig.CoinGeckoToken, []string{tokenAddress}, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	// Store the complete map for this token in Redis
	jsonData, err := json.Marshal(tokenInfo[tokenAddress])
	if err != nil {
		return nil, err
	}
	s.server.Redis.Client.Set(ctx, cacheKey, string(jsonData), time.Minute*5)

	return tokenInfo[tokenAddress], nil
}

func (s *MarketplaceService) helperGetAddressFromPrivateKey(ctx context.Context, secretKey string) (string, error) {
	// prepare for sign transaction
	privateKey, err := crypto.HexToECDSA(secretKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return strings.ToLower(fromAddress.Hex()), nil
}

func (s *MarketplaceService) helperGetTokenInfo(ctx context.Context, tokenAddress string) (*protjectDto.TokenInfo, error) {
	cacheKey := "token_info_" + tokenAddress
	rs := s.server.Redis.Client.Get(ctx, cacheKey)
	if rs.Err() == nil {
		var tokenInfo *protjectDto.TokenInfo
		err := json.Unmarshal([]byte(rs.Val()), &tokenInfo)
		if err != nil {
			return nil, err
		}
		return tokenInfo, nil
	}

	tokenInfo, err := httpClientCGK.GetCoinGecKoTokenInfo(s.server.Redis.Client, s.server.ExtendConfig.CoinGeckoURL, s.server.ExtendConfig.CoinGeckoToken, tokenAddress, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	// set to cache
	json, err := json.Marshal(tokenInfo)
	if err != nil {
		return nil, err
	}
	s.server.Redis.Client.Set(ctx, cacheKey, string(json), time.Hour*24)

	return tokenInfo, nil
}

func (s *MarketplaceService) helperGetTokenPriceOkx(ctx context.Context, tokenAddress string, decimals int, chainId int) (float64, float64, error) {
	tokenAddress = strings.ToLower(tokenAddress)
	price, priceImpact, err := httpClient.GetOkxTokenPrice(tokenAddress, decimals, chainId)
	if err != nil || price == 0 {
		priceTmp, err := s.helperTokenPrice(ctx, tokenAddress, chainId)
		if err != nil {
			return 0, 0.05, err
		}
		price = priceTmp["usd"]
	}

	slippageNumber := math.Abs(priceImpact * 1.25)
	if slippageNumber < 1 {
		slippageNumber = 1
	}
	slippageNumber = slippageNumber / 100

	return price, slippageNumber, nil
}
