package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	marketplaceDto "github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	nodeService "github.com/Rivalz-ai/framework-be/modules/node/service"
	"github.com/Rivalz-ai/framework-be/modules/project/dto"
	"github.com/Rivalz-ai/framework-be/server"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
)

type MarketplaceService struct {
	server                    *server.Server
	colMarketplace            *mongo.Collection
	colTransaction            *mongo.Collection
	colMkpMyAgent             *mongo.Collection
	colMkpRAgent              *mongo.Collection
	colProject                *mongo.Collection
	colMkpSetting             *mongo.Collection
	colMkpUser                *mongo.Collection
	colMkpBid                 *mongo.Collection
	colMkpUserBid             *mongo.Collection
	colMkpSwapWallet          *mongo.Collection
	colSession                *mongo.Collection
	colUserReward             *mongo.Collection
	colUser                   *mongo.Collection
	colUserHoldAgent          *mongo.Collection
	colSessionLevel           *mongo.Collection
	colSessionUserPointReward *mongo.Collection
}

func NewMarketplaceService(sv *server.Server) (*MarketplaceService, error) {
	db, err := sv.Mgo.GetDB("rome")
	if err != nil {
		return nil, err
	}
	return &MarketplaceService{
		server:                    sv,
		colMarketplace:            db.Collection("marketplace"),
		colTransaction:            db.Collection("transactions"),
		colMkpMyAgent:             db.Collection("mkp_ragent_user_buy"),
		colMkpRAgent:              db.Collection("mkp_ragent"),
		colProject:                db.Collection("projects"),
		colMkpSetting:             db.Collection("mkp_setting"),
		colMkpUser:                db.Collection("mkp_user"),
		colMkpBid:                 db.Collection("mkp_bid"),
		colMkpUserBid:             db.Collection("mkp_user_bid"),
		colMkpSwapWallet:          db.Collection("mkp_swap_wallet"),
		colSession:                db.Collection("season_info"),
		colUserReward:             db.Collection("userrewards"),
		colUser:                   db.Collection("users"),
		colUserHoldAgent:          db.Collection("mkp_user_hold_snapshot"),
		colSessionLevel:           db.Collection("session_level"),
		colSessionUserPointReward: db.Collection("session_user_point_reward"),
	}, nil
}

func (s *MarketplaceService) GetTokenPrice(ctx context.Context, tokenAddress string, decimals int, chainId int) (float64, error) {
	if chainId == 0 {
		chainId = int(define.BASE)
	}
	price, _, err := s.helperGetTokenPriceOkx(ctx, tokenAddress, decimals, chainId)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func (s *MarketplaceService) GetTotalValueByWallet(ctx context.Context, walletAddress string) (float64, error) {
	walletAddress = strings.ToLower(walletAddress)
	filter := bson.M{"wallet": walletAddress}
	var mkpUsers []*marketplaceDto.MkpUser
	cursor, err := s.colMkpUser.Find(ctx, filter)
	if err != nil {
		return 0, err
	}
	err = cursor.All(ctx, &mkpUsers)
	if err != nil {
		return 0, err
	}

	if len(mkpUsers) == 0 {
		return 0, nil
	}

	// Use a mutex for safely updating the total value
	var mu sync.Mutex
	totalValue := 0.0

	// Use errgroup to process in parallel
	errGroup, ctx := errgroup.WithContext(ctx)

	for _, mkpUser := range mkpUsers {
		mkpUser := mkpUser // Capture the loop variable

		errGroup.Go(func() error {
			rAgentEstValue, _, _, _, err := s.helperGetRAgentEstValue(ctx, mkpUser.RagentId)
			if err != nil {
				return err
			}

			// Safely update the total value
			mu.Lock()
			totalValue += rAgentEstValue * float64(mkpUser.Quantity)
			mu.Unlock()

			return nil
		})
	}

	// Wait for all goroutines and handle any errors
	if err := errGroup.Wait(); err != nil {
		return 0, err
	}

	return totalValue, nil
}

func (s *MarketplaceService) GetRagentDetail(ctx context.Context, ragent_id, wallet_address string) (*marketplaceDto.RagentDetailResponse, error) {
	wallet_address = strings.ToLower(wallet_address)
	// get mkp_user
	filter := bson.M{"wallet": wallet_address, "ragent_id": ragent_id}
	var mkpUser *marketplaceDto.MkpUser
	err := s.colMkpUser.FindOne(ctx, filter).Decode(&mkpUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user does not have any token")
		}
		return nil, err
	}

	quantity := mkpUser.Quantity

	// get ragent from ragent_id
	ragentObjectId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return nil, err
	}
	filter = bson.M{"_id": ragentObjectId}
	var ragent *marketplaceDto.MkpRagent
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&ragent)
	if err != nil {
		return nil, err
	}

	amountToken := math.Floor(ragent.ProjectToken*100) / 100
	amountRiz := math.Floor(ragent.RizToken*100) / 100

	return &marketplaceDto.RagentDetailResponse{
		Balance:     quantity,
		CompanyName: ragent.Company,
		CompanyLogo: ragent.Logo,
		TokenLogo:   ragent.TokenLogo,
		TokenSymbol: ragent.TokenSymbol,
		AmountToken: amountToken,
		AmountRiz:   amountRiz,
	}, nil
}

func (s *MarketplaceService) GetRagentSellUnwrapInfo(ctx context.Context, ragentId, walletAddress string) (*marketplaceDto.RagentSellUnwrapInfoResponse, error) {
	walletKeeper, err := s.helperGetAddressFromPrivateKey(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey)
	if err != nil {
		return nil, err
	}

	// get ragent from ragent_id
	ragentObjectId, err := primitive.ObjectIDFromHex(ragentId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": ragentObjectId}
	var ragent *marketplaceDto.MkpRagent
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&ragent)
	if err != nil {
		return nil, err
	}

	// get mkp_user from wallet_address and ragent_id
	filter = bson.M{"wallet": walletAddress, "ragent_id": ragentId}
	var mkpUser *marketplaceDto.MkpUser
	err = s.colMkpUser.FindOne(ctx, filter).Decode(&mkpUser)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		mkpUser = nil
	}

	tokenIds := []int64{}
	if mkpUser != nil {
		for _, tokenId := range mkpUser.AccountList {
			tokenIdInt, err := strconv.ParseInt(tokenId, 10, 64)
			if err != nil {
				return nil, err
			}
			tokenIds = append(tokenIds, tokenIdInt)
		}
	}

	return &marketplaceDto.RagentSellUnwrapInfoResponse{
		NFTContractAddress: ragent.RagentTokenAddress,
		TokenIds:           tokenIds,
		WalletKeeper:       walletKeeper,
	}, nil
}

func (s *MarketplaceService) GetStatus(ctx context.Context, tracking_id string) (string, error) {
	myAgentId, err := primitive.ObjectIDFromHex(tracking_id)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": myAgentId}
	var mkpRagentUserBuy *marketplaceDto.MkpRagentUserBuy
	err = s.colMkpMyAgent.FindOne(ctx, filter).Decode(&mkpRagentUserBuy)
	if err != nil {
		return "", err
	}
	return mkpRagentUserBuy.Status, nil
}

func (s *MarketplaceService) GetMkpRAgent(ctx context.Context, walletAddress string, page, limit int) ([]*marketplaceDto.MkpRagent, error) {
	walletAddress = strings.ToLower(walletAddress)

	// get from cache
	cacheKey := fmt.Sprintf("mkp_ragent_%s_%d_%d", walletAddress, page, limit)
	ragentsCache, err := s.server.Redis.Client.Get(ctx, cacheKey).Result()
	if err == nil {
		var ragents []*marketplaceDto.MkpRagent
		err = json.Unmarshal([]byte(ragentsCache), &ragents)
		if err == nil {
			return ragents, nil
		}
	}

	filter := bson.M{"status": "active"}
	var ragents []*marketplaceDto.MkpRagent
	skip := int64((page - 1) * limit)
	limit64 := int64(limit)
	cursor, err := s.colMkpRAgent.Find(ctx, filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit64,
		Sort:  bson.M{"_id": 1},
	})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &ragents)
	if err != nil {
		return nil, err
	}

	// Get mkp setting early to avoid repeated lookups
	mkpSetting, err := s.helperGetMkpSetting(ctx)
	if err != nil {
		return nil, err
	}

	// Fetch agent IDs owned by wallet in one database query
	var mapAgentId = make(map[string]bool)
	if walletAddress != "0x" {
		rs, err := s.colMkpUser.Distinct(ctx, "ragent_id", bson.M{"wallet": walletAddress, "quantity": bson.M{"$gt": 0}})
		if err != nil {
			return nil, err
		}

		for _, agentId := range rs {
			mapAgentId[agentId.(string)] = true
		}
	}

	// Use errgroup to handle errors from goroutines
	errGroup, ctx := errgroup.WithContext(ctx)

	// Create mutex to protect concurrent map access
	var mu sync.Mutex

	// Process each ragent concurrently
	for i := range ragents {
		i := i // Capture loop variable
		ragent := ragents[i]

		errGroup.Go(func() error {
			ragentPrice, _, _, priceImpact, err := s.helperGetRAgentEstValue(ctx, ragent.ID.Hex())
			if err != nil {
				return err
			}

			slippage := math.Abs(priceImpact)
			if slippage < 1 {
				slippage = 1
			}
			slippage = slippage / 100

			// Get token price
			tokenPrice, err := s.helperTokenPrice(ctx, ragent.TokenAddress, ragent.ChainId)
			if err != nil {
				tokenPrice = make(map[string]float64)
			}

			// get token riz price
			rizTokenPrice, err := s.helperTokenPrice(ctx, s.server.ExtendConfig.RIZAddress, int(define.BASE))
			if err != nil {
				return err
			}

			// count number of bids
			count, err := s.colMkpBid.CountDocuments(ctx, bson.M{"ragent_id": ragent.ID.Hex(), "status": "active"})
			if err != nil {
				return err
			}

			// Update ragent with calculated values under mutex lock
			mu.Lock()
			defer mu.Unlock()

			// get riz token price change in 24h
			rizTokenPriceChange := 0.0
			if rizTokenPrice["usd_24h_change"] != 0 {
				rizTokenPriceChange = math.Abs(rizTokenPrice["usd_24h_change"])
			}

			ragent.MarketCap = tokenPrice["usd_market_cap"]
			if ragent.MarketCap <= 0 {
				ragent.AgentType = "Pre-TG"
			} else {
				ragent.AgentType = "Infra"
			}
			ragent.APY = 5
			ragent.BuyPrice = ragentPrice * (1 + mkpSetting.BuyFee + slippage)
			if tokenPrice["usd_24h_change"] != 0 {
				ragent.Buy24hChange = math.Abs((tokenPrice["usd_24h_change"]*ragent.ProjectToken + rizTokenPriceChange*ragent.RizToken) / (ragent.ProjectToken + ragent.RizToken))
			}

			ragent.SellPrice = ragentPrice * (1 - slippage - mkpSetting.SellFee)
			if tokenPrice["usd_24h_change"] != 0 {
				ragent.Sell24hChange = math.Abs((tokenPrice["usd_24h_change"]*ragent.ProjectToken + rizTokenPriceChange*ragent.RizToken) / (ragent.ProjectToken + ragent.RizToken))
			}

			if _, ok := mapAgentId[ragent.ID.Hex()]; ok {
				ragent.IsMyAgentOnly = true
			}

			// ragent.Bids = bids
			// if len(bids) == 0 {
			// 	ragent.Bids = []*marketplaceDto.BidResponse{}
			// }
			ragent.ActiveBid = int(count)

			return nil
		})
	}

	// Wait for all goroutines to complete and check for errors
	if err := errGroup.Wait(); err != nil {
		return nil, err
	}

	// save to cache
	ragentsJson, err := json.Marshal(ragents)
	if err != nil {
		return nil, err
	}

	s.server.Redis.Client.Set(ctx, cacheKey, string(ragentsJson), time.Second*60)

	return ragents, nil
}

func (s *MarketplaceService) GetMkpMyAgent(ctx context.Context, wallet_address string) ([]*marketplaceDto.MkpRagentUserBuyResponse, float64, error) {
	wallet_address = strings.ToLower(wallet_address)
	filter := bson.M{"wallet": wallet_address, "quantity": bson.M{"$gt": 0}}
	var mkpUsers []*marketplaceDto.MkpUser
	cursor, err := s.colMkpUser.Find(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	err = cursor.All(ctx, &mkpUsers)
	if err != nil {
		return nil, 0, err
	}

	// Get mkp setting early
	mkpSetting, err := s.helperGetMkpSetting(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Prepare results array with capacity
	mkpMyAgentsResponse := make([]*marketplaceDto.MkpRagentUserBuyResponse, len(mkpUsers))

	// Use atomic for thread-safe total value calculation
	var totalValueAtomic atomic.Value
	totalValueAtomic.Store(0.0)
	//get current level of user for calculate apy
	userLevel, err := s.GetCurrentLevelUser(ctx, wallet_address)
	if err != nil {
		return nil, 0, err
	}
	apy := getAPYBootRateInt(int(utils.StringToInt64(userLevel.SessionLevel.Name)))
	// Create a mutex for protecting the total value calculation
	var mu sync.Mutex

	// Use errgroup to handle errors from goroutines
	errGroup, ctx := errgroup.WithContext(ctx)

	// Process each user concurrently
	for i, mkpUser := range mkpUsers {
		i, mkpUser := i, mkpUser // Capture loop variables

		errGroup.Go(func() error {
			// get ragent from ragent_id
			ragentId, err := primitive.ObjectIDFromHex(mkpUser.RagentId)
			if err != nil {
				return err
			}
			filter := bson.M{"_id": ragentId}
			var ragent *marketplaceDto.MkpRagent
			err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&ragent)
			if err != nil {
				return err
			}

			// get market cap of token in ragent
			tokenPrice, err := s.helperTokenPrice(ctx, ragent.TokenAddress, ragent.ChainId)
			if err == nil {
				ragent.MarketCap = tokenPrice["usd_market_cap"]

			}
			if ragent.MarketCap <= 0 {
				ragent.AgentType = "Pre-TG"
			} else {
				ragent.AgentType = "Infra"
			}
			rAgentEstValue, _, _, priceImpact, err := s.helperGetRAgentEstValue(ctx, ragent.ID.Hex())
			if err != nil {
				return err
			}

			slippage := math.Abs(priceImpact)
			if slippage < 1 {
				slippage = 1
			}
			slippage = slippage / 100

			// count number of bids
			count, err := s.colMkpBid.CountDocuments(ctx, bson.M{"ragent_id": ragent.ID.Hex(), "status": "active"})
			if err != nil {
				return err
			}

			// Create response with gathered data
			myrAgent := &marketplaceDto.MkpRagentUserBuyResponse{
				WalletAddress: mkpUser.Wallet,
				MkpRagentId:   mkpUser.RagentId,
				MarketCap:     ragent.MarketCap,
				SellPrice:     rAgentEstValue * (1 - slippage) * (1 - mkpSetting.SellFee),
				X:             ragent.X,
				Coingecko:     ragent.Coingecko,
				CookieDAO:     ragent.CookieDAO,
				DEXScreener:   ragent.DEXScreener,
				USDValue:      rAgentEstValue * float64(mkpUser.Quantity),
				RagentName:    ragent.Name,
				Company:       ragent.Company,
				CompanyLogo:   ragent.Logo,
				TokenLogo:     ragent.TokenLogo,
				Quantity:      mkpUser.Quantity,
				ActiveBid:     count,
				ChainId:       ragent.ChainId,
				APY:           apy,
			}
			if myrAgent.MarketCap <= 0 {
				myrAgent.AgentType = "Pre-TG"
			} else {
				myrAgent.AgentType = "Infra"
			}
			// Update the response array and add to total under mutex lock
			mu.Lock()
			defer mu.Unlock()

			mkpMyAgentsResponse[i] = myrAgent
			currentTotal := totalValueAtomic.Load().(float64)
			totalValueAtomic.Store(currentTotal + myrAgent.USDValue)

			return nil
		})
	}

	// Wait for all goroutines to complete and check for errors
	if err := errGroup.Wait(); err != nil {
		return nil, 0, err
	}

	return mkpMyAgentsResponse, totalValueAtomic.Load().(float64), nil
}

func (s *MarketplaceService) SwarmOwner(ctx context.Context, wallet_address string) (bool, error) {
	wallet_address = strings.ToLower(wallet_address)
	// check if wallet is an owner of any project
	filter := bson.M{"creator": wallet_address}
	var projects []*dto.Project
	cursor, err := s.colProject.Find(ctx, filter)
	if err != nil {
		return false, err
	}
	err = cursor.All(ctx, &projects)
	if err != nil {
		return false, err
	}
	if len(projects) > 0 {
		return true, nil
	}
	return false, nil
}

func (s *MarketplaceService) GetQuantityToken(ctx context.Context, ragent_id string, quantity int64, side string, wallet_address string) (*marketplaceDto.GetQuantityTokenResponse, map[string]marketplaceDto.BalanceTokenRagent, error) {

	// get mkp setting
	mkpSetting, err := s.helperGetMkpSetting(ctx)
	if err != nil {
		return nil, nil, err
	}

	// get ragent
	agentId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return nil, nil, err
	}
	filter := bson.M{"_id": agentId}
	var mkpRagent *marketplaceDto.MkpRagent
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return nil, nil, err
	}

	// get ragent est value and price from 1inch
	rAgentEstValue, priceProject, priceRiz, priceImpact, err := s.helperGetRAgentEstValue(ctx, ragent_id)
	if err != nil {
		return nil, nil, err
	}

	if side == "buy" && (priceProject == 0 || priceRiz == 0) {
		return nil, nil, errors.New("internal error: can not get price")
	}

	slippage := math.Abs(priceImpact)
	if slippage < 1 {
		slippage = 1
	}
	slippage = slippage / 100

	fee := (1 + mkpSetting.BuyFee + mkpSetting.SwapFee + slippage)
	if side == "sell" {
		fee = (1 - mkpSetting.SellFee - mkpSetting.SwapFee - slippage)
	} else if side == "unwrap" {
		fee = (1 - mkpSetting.UnwrapFee)
	}

	totalPrice := rAgentEstValue * float64(quantity) * fee

	// if sell -> we need to calculate the amount of token need to sell of each account
	var mapAccountBalance = make(map[string]marketplaceDto.BalanceTokenRagent)

	var tokenIds = []int64{}
	if side == "sell" {
		isNeedReCalculate := false

		// call erc721 to get list of token id
		// nodeClient, err := nodeService.NewNodeService(s.server)
		// if err != nil {
		// 	return nil, nil, err
		// }
		// tokenIds, err := nodeClient.GetTokenIds(ctx, mkpRagent.RagentTokenAddress, wallet_address, quantity)
		// if err != nil {
		// 	return nil, nil, err
		// }

		// get mkp_user by wallet address + ragent id
		filter = bson.M{"wallet": wallet_address, "ragent_id": ragent_id}
		var mkpUser *marketplaceDto.MkpUser
		err = s.colMkpUser.FindOne(ctx, filter).Decode(&mkpUser)
		if err != nil {
			return nil, nil, err
		}

		if mkpUser.AccountList == nil || len(mkpUser.AccountList) < int(quantity) {
			return nil, nil, errors.New("account not enough")
		}

		// calculate config amount of token in each account
		projectTokenBigFloat := new(big.Float).SetFloat64(mkpRagent.ProjectToken)
		decimalMultiplier := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(mkpRagent.TokenDecimals)), nil))
		result := new(big.Float).Mul(projectTokenBigFloat, decimalMultiplier)
		amountDstToken, _ := result.Int(nil)
		amountRizToken := big.NewInt(int64(mkpRagent.RizToken * 1e8))

		nodeClient, err := nodeService.NewNodeService(s.server)
		if err != nil {
			return nil, nil, err
		}

		nodeClientProject, err := nodeService.NewNodeService(s.server, mkpRagent.ChainId)
		if err != nil {
			return nil, nil, err
		}

		count := 0
		quantityCheck := quantity
		// if mkpRagent.ChainId != int(define.BASE) {
		// 	quantityCheck = quantity * 2
		// }

		// get config amount of token in each account
		combinedAccountList := make(map[string]string)
		combinedAccountListBase := make(map[string]string)
		for account, tokenId := range mkpUser.AccountList {
			combinedAccountList[account] = tokenId
		}
		for tokenId, account := range mkpUser.AccountListBase {
			combinedAccountListBase[tokenId] = account
		}
		checkTokenIds := make(map[string]bool)

		for account, tokenId := range combinedAccountList {
			if count >= int(quantityCheck) {
				break
			}
			count++

			tokenIdInt, err := strconv.ParseInt(tokenId, 10, 64)
			if err != nil {
				return nil, nil, err
			}
			if !checkTokenIds[tokenId] {
				tokenIds = append(tokenIds, tokenIdInt)
				checkTokenIds[tokenId] = true
			}

			if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "local" {
				s.server.ExtendConfig.RIZAddress = mkpRagent.TokenAddress
			}

			// get balance of account for project token and riz token
			var projectTokenBalance *big.Int
			if mkpRagent.ChainId != int(define.SOLANA) {
				projectTokenBalance, err = nodeClientProject.GetBalanceERC20(ctx, account, mkpRagent.TokenAddress)
				if err != nil {
					return nil, nil, err
				}
			}

			rizTokenBalance, err := nodeClient.GetBalanceERC20(ctx, account, s.server.ExtendConfig.RIZAddress)
			if err != nil {
				return nil, nil, err
			}

			accountBase := combinedAccountListBase[tokenId]
			if accountBase != "" {
				if mkpRagent.ChainId != int(define.SOLANA) {
					projectTokenBalance, err = nodeClientProject.GetBalanceERC20(ctx, accountBase, mkpRagent.TokenAddress)
					if err != nil {
						return nil, nil, err
					}
				} else {
					projectTokenBalanceTmp, err := nodeClientProject.GetSolanaTokenBalance(ctx, accountBase, mkpRagent.TokenAddress)
					if err != nil {
						return nil, nil, err
					}
					projectTokenBalance = big.NewInt(int64(projectTokenBalanceTmp))
				}
			}

			balanceProject := amountDstToken
			balanceRiz := amountRizToken
			if mkpRagent.TokenAddress != s.server.ExtendConfig.RIZAddress {
				if projectTokenBalance.Cmp(amountDstToken) != 0 {
					balanceProject = projectTokenBalance
					isNeedReCalculate = true
				}

				if rizTokenBalance.Cmp(amountRizToken) != 0 {
					balanceRiz = rizTokenBalance
					isNeedReCalculate = true
				}
			}

			mapAccountBalance[account+"_"+tokenId] = marketplaceDto.BalanceTokenRagent{
				ProjectToken: balanceProject.String(),
				RizToken:     balanceRiz.String(),
			}
		}

		if isNeedReCalculate {
			totalPrice = 0
			for _, balance := range mapAccountBalance {
				projectTokenFloat, _ := new(big.Float).SetString(balance.ProjectToken)
				priceProjectDec := new(big.Float).Mul(projectTokenFloat, big.NewFloat(priceProject))
				priceProject, _ := new(big.Float).Quo(priceProjectDec, decimalMultiplier).Float64()
				rizTokenFloat, _ := new(big.Float).SetString(balance.RizToken)
				priceRizDec := new(big.Float).Mul(rizTokenFloat, big.NewFloat(priceRiz))
				priceRiz, _ := new(big.Float).Quo(priceRizDec, decimalMultiplier).Float64()
				totalPrice += priceProject + priceRiz
			}
		}
	}

	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "local" {
		totalPrice = rAgentEstValue * float64(quantity) * fee
		mapAccountBalance = make(map[string]marketplaceDto.BalanceTokenRagent)
		for _, tokenId := range tokenIds {
			tokenIdStr := strconv.FormatInt(tokenId, 10)

			projectTokenFloat := new(big.Float).SetFloat64(mkpRagent.ProjectToken)
			projectTokenFloat.Mul(projectTokenFloat, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(mkpRagent.TokenDecimals)), nil)))
			projectTokenInt, _ := projectTokenFloat.Int(nil)
			rizTokenFloat := new(big.Float).SetFloat64(mkpRagent.RizToken)
			rizTokenFloat.Mul(rizTokenFloat, new(big.Float).SetInt(big.NewInt(1e8)))
			rizTokenInt, _ := rizTokenFloat.Int(nil)

			mapAccountBalance[wallet_address+"_"+tokenIdStr] = marketplaceDto.BalanceTokenRagent{
				ProjectToken: projectTokenInt.String(),
				RizToken:     rizTokenInt.String(),
			}
		}
	}

	return &marketplaceDto.GetQuantityTokenResponse{
		RAgentEstValue: rAgentEstValue,
		Slippage:       math.Round((slippage*100)*1000) / 1000,
		TotalPrice:     totalPrice,
		TokenIds:       tokenIds,
	}, mapAccountBalance, nil
}

func (s *MarketplaceService) GetQuantityTokenByTokenIds(ctx context.Context, ragent_id string, quantity int64, side string, wallet_address string, tokenIds []int64) (*marketplaceDto.GetQuantityTokenResponse, map[string]marketplaceDto.BalanceTokenRagent, error) {

	// get mkp setting
	mkpSetting, err := s.helperGetMkpSetting(ctx)
	if err != nil {
		return nil, nil, err
	}

	// get ragent
	agentId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return nil, nil, err
	}
	filter := bson.M{"_id": agentId}
	var mkpRagent *marketplaceDto.MkpRagent
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return nil, nil, err
	}

	// get ragent est value and price from 1inch
	rAgentEstValue, priceProject, priceRiz, priceImpact, err := s.helperGetRAgentEstValue(ctx, ragent_id)
	if err != nil {
		return nil, nil, err
	}

	slippage := math.Abs(priceImpact)
	if slippage < 1 {
		slippage = 1
	}
	slippage = slippage / 100

	fee := (1 + mkpSetting.BuyFee + mkpSetting.SwapFee + slippage)
	if side == "sell" {
		fee = (1 - mkpSetting.SellFee - mkpSetting.SwapFee - slippage)
	} else if side == "unwrap" {
		fee = (1 - mkpSetting.UnwrapFee)
	}

	totalPrice := rAgentEstValue * float64(quantity) * fee

	// if sell -> we need to calculate the amount of token need to sell of each account
	var mapAccountBalance = make(map[string]marketplaceDto.BalanceTokenRagent)

	if side == "sell" || side == "unwrap" {
		isNeedReCalculate := false

		// get mkp_user by wallet address + ragent id
		filter = bson.M{"wallet": wallet_address, "ragent_id": ragent_id}
		var mkpUser *marketplaceDto.MkpUser
		err = s.colMkpUser.FindOne(ctx, filter).Decode(&mkpUser)
		if err != nil {
			return nil, nil, err
		}

		mapTokenIdToAccount := make(map[int64]string)
		for account, tokenId := range mkpUser.AccountList {
			tokenIdInt, err := strconv.ParseInt(tokenId, 10, 64)
			if err != nil {
				return nil, nil, err
			}
			mapTokenIdToAccount[tokenIdInt] = account
		}

		mapTokenIdToAccountBase := make(map[int64]string)
		for tokenId, account := range mkpUser.AccountListBase {
			tokenIdInt, err := strconv.ParseInt(tokenId, 10, 64)
			if err != nil {
				return nil, nil, err
			}
			mapTokenIdToAccountBase[tokenIdInt] = account
		}

		// calculate config amount of token in each account
		projectTokenBigFloat := new(big.Float).SetFloat64(mkpRagent.ProjectToken)
		decimalMultiplier := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(mkpRagent.TokenDecimals)), nil))
		result := new(big.Float).Mul(projectTokenBigFloat, decimalMultiplier)
		amountDstToken, _ := result.Int(nil)
		amountRizToken := big.NewInt(int64(mkpRagent.RizToken * 1e8))

		nodeClient, err := nodeService.NewNodeService(s.server)
		if err != nil {
			return nil, nil, err
		}

		nodeClientProject, err := nodeService.NewNodeService(s.server, mkpRagent.ChainId)
		if err != nil {
			return nil, nil, err
		}

		// get config amount of token in each account
		for _, tokenId := range tokenIds {
			account := mapTokenIdToAccount[tokenId]
			if account == "" {
				return nil, nil, errors.New("account not found")
			}

			// get balance of account for project token and riz token
			var projectTokenBalance *big.Int
			if mkpRagent.ChainId != int(define.SOLANA) {
				projectTokenBalance, err = nodeClientProject.GetBalanceERC20(ctx, account, mkpRagent.TokenAddress)
				if err != nil {
					return nil, nil, err
				}
			}
			rizTokenBalance, err := nodeClient.GetBalanceERC20(ctx, account, s.server.ExtendConfig.RIZAddress)
			if err != nil {
				return nil, nil, err
			}

			accountBase := mapTokenIdToAccountBase[tokenId]
			if accountBase != "" {
				if mkpRagent.ChainId != int(define.SOLANA) {
					projectTokenBalance, err = nodeClientProject.GetBalanceERC20(ctx, accountBase, mkpRagent.TokenAddress)
					if err != nil {
						return nil, nil, err
					}
				} else {
					projectTokenBalanceTmp, err := nodeClientProject.GetSolanaTokenBalance(ctx, accountBase, mkpRagent.TokenAddress)
					if err != nil {
						return nil, nil, err
					}
					projectTokenBalance = big.NewInt(int64(projectTokenBalanceTmp))
				}
			}

			balanceProject := amountDstToken
			balanceRiz := amountRizToken
			if mkpRagent.TokenAddress != s.server.ExtendConfig.RIZAddress {
				if projectTokenBalance.Cmp(amountDstToken) != 0 {
					balanceProject = projectTokenBalance
					isNeedReCalculate = true
				}

				if rizTokenBalance.Cmp(amountRizToken) != 0 {
					balanceRiz = rizTokenBalance
					isNeedReCalculate = true
				}
			}

			tokenIdStr := strconv.FormatInt(tokenId, 10)
			mapAccountBalance[account+"_"+tokenIdStr] = marketplaceDto.BalanceTokenRagent{
				ProjectToken: balanceProject.String(),
				RizToken:     balanceRiz.String(),
			}
		}

		if isNeedReCalculate {
			totalPrice = 0
			for _, balance := range mapAccountBalance {
				projectTokenFloat, _ := new(big.Float).SetString(balance.ProjectToken)
				priceProjectDec := new(big.Float).Mul(projectTokenFloat, big.NewFloat(priceProject))
				priceProject, _ := new(big.Float).Quo(priceProjectDec, decimalMultiplier).Float64()
				rizTokenFloat, _ := new(big.Float).SetString(balance.RizToken)
				priceRizDec := new(big.Float).Mul(rizTokenFloat, big.NewFloat(priceRiz))
				priceRiz, _ := new(big.Float).Quo(priceRizDec, decimalMultiplier).Float64()
				totalPrice += priceProject + priceRiz
			}
		}
	}

	return &marketplaceDto.GetQuantityTokenResponse{
		RAgentEstValue: rAgentEstValue,
		Slippage:       math.Round((slippage*100)*1000) / 1000,
		TotalPrice:     totalPrice,
		TokenIds:       tokenIds,
	}, mapAccountBalance, nil
}

func (s *MarketplaceService) GetRagentHistory(ctx context.Context, wallet_address string) ([]*marketplaceDto.MkpRagentUserBuyResponse, error) {
	wallet_address = strings.ToLower(wallet_address)
	filter := bson.M{
		"wallet_address": wallet_address,
		"$or": []bson.M{
			{
				"side":   "buy",
				"status": bson.M{"$nin": []string{"pending", "success"}},
			},
			{
				"side":   bson.M{"$ne": "buy"},
				"status": bson.M{"$ne": "success"},
			},
		},
	}
	var mkpUsers []*marketplaceDto.MkpRagentUserBuy
	cursor, err := s.colMkpMyAgent.Find(ctx, filter, options.Find().SetSort(bson.M{"createdAt": -1}))
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &mkpUsers)
	if err != nil {
		return nil, err
	}

	mkpSetting, err := s.helperGetMkpSetting(ctx)
	if err != nil {
		return nil, err
	}

	mapRagentId := make(map[string][]*marketplaceDto.MkpRagentUserBuy)
	for _, mkpUser := range mkpUsers {
		mapRagentId[mkpUser.MkpRagentId] = append(mapRagentId[mkpUser.MkpRagentId], mkpUser)
	}

	mkpMyAgentsResponse := make([]*marketplaceDto.MkpRagentUserBuyResponse, 0, len(mkpUsers))
	for agentID, mkpUsers := range mapRagentId {
		// get ragent info
		agentId, err := primitive.ObjectIDFromHex(agentID)
		if err != nil {
			return nil, err
		}
		filter := bson.M{"_id": agentId}
		var ragent *marketplaceDto.MkpRagent
		err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&ragent)
		if err != nil {
			return nil, err
		}

		rAgentEstValue, _, _, priceImpact, err := s.helperGetRAgentEstValue(ctx, ragent.ID.Hex())
		if err != nil {
			return nil, err
		}

		slippage := math.Abs(priceImpact)
		if slippage < 1 {
			slippage = 1
		}
		slippage = slippage / 100

		for _, mkpUser := range mkpUsers {
			amountInt, err := strconv.ParseInt(mkpUser.Amount, 10, 64)
			if err != nil {
				return nil, err
			}

			mkpMyAgentsResponse = append(mkpMyAgentsResponse, &marketplaceDto.MkpRagentUserBuyResponse{
				WalletAddress: mkpUser.WalletAddress,
				MkpRagentId:   mkpUser.MkpRagentId,
				MarketCap:     ragent.MarketCap,
				SellPrice:     rAgentEstValue * (1 - slippage) * (1 - mkpSetting.SellFee),
				X:             ragent.X,
				Coingecko:     ragent.Coingecko,
				CookieDAO:     ragent.CookieDAO,
				DEXScreener:   ragent.DEXScreener,
				USDValue:      rAgentEstValue * float64(amountInt),
				RagentName:    ragent.Company,
				Company:       ragent.Company,
				CompanyLogo:   ragent.Logo,
				TokenLogo:     ragent.TokenLogo,
				Quantity:      amountInt,
				Side:          mkpUser.Side,
				Status:        mkpUser.Status,
				TrackingId:    mkpUser.ID.Hex(),
				TxHash:        mkpUser.TxHash,
			})
		}
	}

	return mkpMyAgentsResponse, nil
}
func (s *MarketplaceService) RemoveUserHoldAgent(ctx context.Context, wallet_address string, ragent_id, token_id string) error {
	/*filter := bson.M{"wallet_address": wallet_address, "ragent_id": ragent_id, "token_id": token_id}
	_, err := s.colUserHoldAgent.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}*/
	return nil
}

func (s *MarketplaceService) GetCurrentLevelUser(ctx context.Context, wallet string) (*marketplaceDto.UserLevel, error) {
	wallet = strings.ToLower(wallet)
	//get level of user wallet
	filter := bson.M{}
	result, err := s.colSessionLevel.Find(ctx, filter)
	var level []marketplaceDto.SessionLevel
	if err != nil {
		log.Error("error when find level: "+err.Error(), "GetCurrentLevelUser", wallet)
		return nil, err
	}
	if err = result.All(ctx, &level); err != nil {
		log.Error("error when find level: "+err.Error(), "GetCurrentLevelUser", wallet)
		return nil, err
	}

	//get user current reward by wallet
	filter = bson.M{"wallet_address": wallet}
	userPointReward := s.colSessionUserPointReward.FindOne(ctx, filter)
	var userPointRewardData *marketplaceDto.SessionUserReward
	if err := userPointReward.Decode(&userPointRewardData); err != nil {
		if err == mongo.ErrNoDocuments {
			userPointRewardData = &marketplaceDto.SessionUserReward{
				Point:  0,
				Reward: 0,
			}
		} else {
			log.Error("error when find user point reward: "+err.Error(), "GetCurrentLevelUser", wallet)
			return nil, err
		}
	}
	//find level of user base on reward and level.reward
	level_name := "0"
	session_level := &marketplaceDto.SessionLevel{}
	for _, level := range level {
		if userPointRewardData.Point >= level.PointLimit {
			level_name = level.Name
			session_level = &level
		} else {
			break
		}
	}
	userLevel := &marketplaceDto.UserLevel{
		Level:        level_name,
		SessionLevel: session_level,
	}
	return userLevel, nil
}
func getAPYBootRateInt(level int) int {
	switch level {
	case 1:
		return 6 //+1
	case 2:
		return 7 //+2
	case 3:
		return 8 //+3
	case 4:
		return 10 //+5
	case 5:
		return 15 //+15
	default:
		return 5 //5%
	}
}
