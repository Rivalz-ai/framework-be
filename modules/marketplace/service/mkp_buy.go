package service

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/framework/base/event"
	"github.com/Rivalz-ai/framework-be/framework/log"
	marketplaceDto "github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	httpClient "github.com/Rivalz-ai/framework-be/modules/marketplace/service/http"
	nodeService "github.com/Rivalz-ai/framework-be/modules/node/service"
	rewardService "github.com/Rivalz-ai/framework-be/modules/reward/service"
	userService "github.com/Rivalz-ai/framework-be/modules/user/service"
	commonTypes "github.com/Rivalz-ai/framework-be/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	associatedtokenaccount "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MarketplaceService) GenerateSignature(ctx context.Context, walletAddress string, quantity int64, ragentId string) (*marketplaceDto.SignBuyRagentMarketResponse, error) {
	walletAddress = strings.ToLower(walletAddress)
	ragentMarketContract := s.server.ExtendConfig.RagentMarketContract
	if ragentMarketContract == "" {
		return nil, errors.New("ragent market contract not found")
	}

	logData := map[string]interface{}{
		"wallet_address": walletAddress,
		"quantity":       quantity,
	}
	log.Info("Sign signature payload: ", "SignSinatureBuyMarketplace-Payload", logData)

	// get ragent info
	agentId, err := primitive.ObjectIDFromHex(ragentId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": agentId}
	var mkpRagent *marketplaceDto.MkpRagent
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return nil, err
	}

	mkpSetting, err := s.helperGetMkpSetting(ctx)
	if err != nil {
		return nil, err
	}

	var totalPrice, price, priceImpact float64

	for i := 0; i < 5; i++ {
		totalPrice, price, _, priceImpact, err = s.helperGetRAgentEstValue(ctx, ragentId)
		if err != nil {
			continue
		}

		if price == 0 {
			time.Sleep(time.Millisecond * 200)
			continue
		}
	}

	if err != nil {
		return nil, err
	}

	if price == 0 {
		return nil, errors.New("internal error: can not get price")
	}

	slippageNumber := math.Abs(priceImpact * 1.25)
	if slippageNumber < 1 {
		slippageNumber = 1
	}
	slippageNumber = slippageNumber / 100
	totalPriceBigInt := big.NewInt(int64(float64(quantity) * totalPrice * (1 + slippageNumber) * (1 + mkpSetting.BuyFee + mkpSetting.SwapFee) * 1e6)) // usdc decimals

	// get nonce
	nonce, err := s.server.Redis.Client.Incr(ctx, "marketplace:buy:nonce:"+strings.ToLower(ragentMarketContract)).Result()
	if err != nil {
		return nil, err
	}

	rAgentTokenContract := mkpRagent.RagentTokenAddress //nft contract

	signature, err := signMessage(walletAddress, big.NewInt(quantity), totalPriceBigInt, ragentMarketContract, rAgentTokenContract, big.NewInt(nonce), s.server.ExtendConfig.RagentMarketContractPrivateKey)
	if err != nil {
		return nil, err
	}

	totalPriceDivide := new(big.Float).Quo(new(big.Float).SetInt(totalPriceBigInt), new(big.Float).SetInt(big.NewInt(1e6)))
	fee := new(big.Float).Mul(totalPriceDivide, new(big.Float).SetFloat64(mkpSetting.BuyFee))
	feeFloat, _ := fee.Float64()

	ragentUserId := primitive.NewObjectID()
	mkpRagentUserBuy := &marketplaceDto.MkpRagentUserBuy{
		ID:                 ragentUserId,
		MkpRagentId:        ragentId,
		WalletAddress:      walletAddress,
		Amount:             strconv.FormatInt(quantity, 10),
		TotalPrice:         totalPriceDivide.String(),
		Status:             "pending",
		Signature:          signature,
		Nonce:              nonce,
		RAgentTokenAddress: rAgentTokenContract,
		CreatedAt:          time.Now().UTC(),
		UpdatedAt:          time.Now().UTC(),
		Side:               "buy",
		TotalFee:           feeFloat,
	}
	_, err = s.colMkpMyAgent.InsertOne(ctx, mkpRagentUserBuy)
	if err != nil {
		log.Error("Error insert mkp_ragent_user_buy: "+err.Error(), "MarketPlaceVerify-InsertMkpRagentUserBuy", logData)
		return nil, err
	}

	tokenAddress := mkpRagent.TokenAddress
	if os.Getenv("ENV") == "dev" {
		tokenAddress = "0xD12eB8776D44507F12045b6931460BAA1401e1D6"
	}

	return &marketplaceDto.SignBuyRagentMarketResponse{
		TrackingId:      ragentUserId.Hex(),
		Signature:       signature,
		Nonce:           nonce,
		WalletAddress:   walletAddress,
		Quantity:        quantity,
		TotalPrice:      totalPriceBigInt.String(),
		TokenAddress:    tokenAddress,
		ContractAddress: ragentMarketContract,
		RagentToken:     rAgentTokenContract,
	}, nil
}

func signMessage(walletAddress string, quantity *big.Int, totalPrice *big.Int, ragentMarketContract, rAgentTokenContract string, nonce *big.Int, privateKeyStr string) (string, error) {
	// Convert private key string to ECDSA private key
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %v", err)
	}

	totalPriceBytes := make([]byte, 32)
	totalPrice.FillBytes(totalPriceBytes)

	quantityBytes := make([]byte, 32)
	quantity.FillBytes(quantityBytes)

	nonceBytes := make([]byte, 32)
	nonce.FillBytes(nonceBytes)

	// Convert string addresses to common.Address type
	walletAddr := common.HexToAddress(walletAddress)
	contractAddr := common.HexToAddress(ragentMarketContract)
	ragentAddr := common.HexToAddress(rAgentTokenContract)

	//- sign 1 chữ ký: (user walletAddress, quantity(số lượng mua), totalPrice(tổng tiền), rAgentmaket address)

	var packedData []byte

	packedData = append(packedData, ragentAddr.Bytes()...)

	// Pack address (20 bytes)
	packedData = append(packedData, walletAddr.Bytes()...)

	// Pack uint256 (quantity)
	packedData = append(packedData, quantityBytes...)

	// Pack uint256 (totalPrice)
	packedData = append(packedData, totalPriceBytes...)

	// Pack contract address
	packedData = append(packedData, contractAddr.Bytes()...)

	// Pack nonce
	packedData = append(packedData, nonceBytes...)

	// Step 3: Hash the packed data with keccak256
	// This is equivalent to keccak256(solidityPacked(...))
	messageHash := crypto.Keccak256(packedData)

	// Step 4: Sign the message with EIP-191 prefix
	// In ethers.js, wallet.signMessage(hashedMessage) adds the EIP-191 prefix automatically
	prefixedMessage := accounts.TextHash(messageHash)

	signature, err := crypto.Sign(prefixedMessage, privateKey)
	if err != nil {
		return "", fmt.Errorf("signing error: %v", err)
	}

	// Adjust v value for Ethereum compatibility (27/28 instead of 0/1)
	signature[64] += 27

	// Format as hex string
	signatureHex := "0x" + common.Bytes2Hex(signature)

	return signatureHex, nil
}

func (s *MarketplaceService) VerifyBuyTransaction(ctx context.Context, ragent_id, user_id, hash, wallet_address string, is_retry bool, quantity int64, nonce int64) (error, bool) {
	wallet_address = strings.ToLower(wallet_address)
	logData := map[string]interface{}{
		"hash":           hash,
		"user_id":        user_id,
		"wallet_address": wallet_address,
	}
	log.Info("MarketPlaceVerify: ", "Verify", logData)
	nodeClient, err := nodeService.NewNodeService(s.server)
	if err != nil {
		return err, true
	}

	//
	filter := bson.M{"hash": hash}
	var txDB *commonTypes.Transaction
	err = s.colTransaction.FindOne(ctx, filter).Decode(&txDB)
	tx_pending := false
	if !is_retry { //nếu call tạo mới
		if err == mongo.ErrNoDocuments {
			//save transaction
			txData := &commonTypes.Transaction{
				Hash:         hash,
				Type:         define.TX_TYPE_BUY_RAGENT_MARKETPLACE,
				ChainId:      8453,
				From:         "",
				Recipient:    "",
				Status:       define.TX_PENDING,
				Value:        0,
				TokenAddress: "",
				Amount:       1,
				CreatedAt:    time.Now().UTC(),
				Source:       "go",
			}
			//insert transaction
			_, erri := s.colTransaction.InsertOne(ctx, txData)
			if erri != nil {
				log.Error("error when insert transaction: "+erri.Error(), "MarketPlaceVerify-InsertTransaction", logData)
				return errors.New("INSERT_TX"), true
			}
			tx_pending = true
		} else if err != nil {
			log.Error("Error read result:"+err.Error(), "MarketPlaceVerify-scanCursor", logData)
			//retry job, db error
			return errors.New("FIND_TX"), true
		} else {
			//just ignore, and process other
			if txDB.Status == define.TX_PENDING {
				tx_pending = true
			} else {
				log.Warn("Transaction success: ", "MarketPlaceVerify-TransactionSuccess", logData)
				return nil, false
			}
		}
	} else {
		//check phải có tx trong transaction collection mới retry, ngược lại ko cần retry
		if err == mongo.ErrNoDocuments {
			log.Error("Retry job, Transaction not found: ", "MarketPlaceVerify-TransactionNotFound", logData)
			return errors.New("Retry job, Transaction not found"), false
		}
		//nếu có tx và đang là retry cần check tx status nếu thành công thì thoát, trả lỗi
		if txDB.Status == define.TX_SUCCESS {
			log.Warn("Retry job, Transaction success: ", "MarketPlaceVerify-TransactionSuccess", logData)
			//return true để api resoponse success, worker ko retry lại
			return nil, false
		}

		tx_pending = true
	}

	//get transaction
	tx, txReceipt, err := nodeClient.GetTransactionReceiptByHash(ctx, hash)
	if err != nil {
		//retry
		return err, true
	}
	if txReceipt == nil || txReceipt.Status == types.ReceiptStatusFailed {
		log.Error("Transaction not found,retry: ", "MarketPlaceVerify-GetTransactionReceiptByHash", hash)
		return errors.New("NO_RECEIPT"), true
	}
	if len(txReceipt.Logs) == 0 {
		log.Error("Transaction logs not found, retry: ", "MarketPlaceVerify-GetTransactionReceiptByHash", hash)
		return errors.New("NO_LOGS"), true
	}

	//prepare ABI parser
	// parsedABI, err := abi.JSON(strings.NewReader(marketplaceABI.RAgentMarketABI))
	// if err != nil {
	// 	log.Error("error when parse ABI: "+err.Error(), "MarketPlaceVerify-ABI", logData)
	// 	return errors.New(define.INTERNAL_SERVER_ERROR), false
	// }

	//find user by user_id
	userSVC, err := userService.NewUserService(s.server)
	if err != nil {
		return err, true
	}
	_, err = userSVC.GetUserByID(ctx, user_id)
	if err != nil {
		if err.Error() == define.NOT_FOUND {
			log.Error("User not found: ", "MarketPlaceVerify-User", logData)
			return errors.New("User not found"), false
		}
		return err, true
	}

	var tokenLendingLogs []*types.Log
	eventName := "0x79f19b3655ee38b1ce526556b7731a20c8f218fbda4a3990b6cc4172fdf88722"
	for _, ll := range txReceipt.Logs {
		if ll.Topics[0].Hex() == eventName {
			tokenLendingLogs = append(tokenLendingLogs, ll)
		}
	}

	// get ragent info
	agentId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return err, true
	}
	filter = bson.M{"_id": agentId}
	var mkpRagent *marketplaceDto.MkpRagent
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return err, true
	}

	mapAccountTokenId := make(map[string]string)
	mapAccountTokenIdBase := make(map[string]string)
	for _, logEvent := range tokenLendingLogs {
		if len(logEvent.Topics) < 3 {
			log.Error("ERC6551AccountCreated log topics not match: ", "MarketPlaceVerify-ERC6551AccountCreated-LogTopics", logData)
			return errors.New("ERC6551AccountCreated log topics not match: " + hash), false
		}

		tokenId := logEvent.Topics[3].Hex()
		tokenIdBigInt := new(big.Int)
		_, ok := tokenIdBigInt.SetString(tokenId[2:], 16)
		if !ok {
			log.Error("Invalid token id: ", "MarketPlaceVerify-InvalidTokenId", logData)
			return errors.New("Invalid token id: " + tokenId), false
		}

		// get account in data
		accountStr := ""
		account := logEvent.Data
		accountStr = hex.EncodeToString(account)
		accountStr = "0x" + accountStr[24:64]
		mapAccountTokenId[accountStr] = tokenIdBigInt.String()
		if mkpRagent.ChainId != int(define.BASE) {
			coinbaseClient, err := httpClient.NewCoinbaseClient(s.server.ExtendConfig.CoinbaseAPIKey, s.server.ExtendConfig.CoinbaseAPIKeySecret, s.server.ExtendConfig.CoinbaseWalletSecret)
			if err != nil {
				return err, true
			}

			var accountStr string
			if mkpRagent.ChainId == int(define.SOLANA) {
				accountStr, err = coinbaseClient.CreateSolanaAccount(ctx)
				if err != nil {
					return err, true
				}

			} else {
				accountStr, err = coinbaseClient.CreateEVMAccount(ctx)
				if err != nil {
					return err, true
				}

			}
			mapAccountTokenIdBase[tokenIdBigInt.String()] = accountStr
		}
	}

	if tx_pending {
		filter := bson.M{"hash": hash}
		update := bson.M{"$set": bson.M{"status": define.TX_SUCCESS, "from": tx.From, "recipient": tx.To}}
		_, err := s.colTransaction.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Error("Error update transaction: "+err.Error(), "MarketPlaceVerify-UpdateTransaction", logData)
			return err, true
		}
	}

	//update mkp_ragent_user_buy
	filter = bson.M{"wallet_address": wallet_address, "nonce": nonce, "mkp_ragent_id": ragent_id}
	update := bson.M{"$set": bson.M{
		"status":            "verify",
		"account_list":      mapAccountTokenId,
		"account_list_base": mapAccountTokenIdBase,
		"user_id":           user_id,
		"updated_at":        time.Now().UTC(),
		"tx_hash":           hash,
	}}
	_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("Error update mkp_ragent_user_buy: "+err.Error(), "MarketPlaceVerify-UpdateMkpRagentUserBuy", logData)
		return err, true
	}
	//resward session 2, first month buy, one-time
	/*err, isBuyAgentFirstMonth := s.IsBuyAgentFirstMonth(ctx, user_id, ragent_id)
	if err != nil {
		log.Error("Error check is buy agent first month: "+err.Error(), "MarketPlaceVerify-IsBuyAgentFirstMonth", logData)
	} else {
		if isBuyAgentFirstMonth {
			rewardSVC, err := rewardService.NewRewardService(s.server)
			if err != nil {
				log.Error("Error create reward service: "+err.Error(), "MarketPlaceVerify-CreateRewardService", logData)
			} else {
				err = rewardSVC.AddUserReward(ctx, user_id, define.SESSION2_MKP_BUY_ANY_AGENT_FRIST_MONTH, hash)
				if err != nil {
					log.Error("Error add user reward: "+err.Error(), "MarketPlaceVerify-AddUserReward", logData)
				}
			}
		}
	}
	*/
	/////////////resward session 2, buy agent, many time
	rewardSVC, err := rewardService.NewRewardService(s.server)
	if err != nil {
		log.Error("Error create reward service: "+err.Error(), define.SESSION2_MKP_BUY_AGENT, logData)
	} else {
		err = rewardSVC.AddUserReward(ctx, user_id, define.SESSION2_MKP_BUY_AGENT, hash, quantity)
		if err != nil {
			log.Error("Error add user reward: "+err.Error(), define.SESSION2_MKP_BUY_AGENT, logData)
		}
	}
	/////////////resward session 2, buy agent highlight, one-time
	/*
		if mkpRagent.IsHighLight {
			err = rewardSVC.AddUserReward(ctx, user_id, define.SESSION2_MKP_BUY_AGENT_HIGHLIGHT, hash)
			if err != nil {
				log.Error("Error add user reward: "+err.Error(), define.SESSION2_MKP_BUY_AGENT_HIGHLIGHT, logData)
			}
		}
	*/
	return nil, false
}

func (s *MarketplaceService) InternalSwapToken(ctx context.Context, ragent_id, quantity, walletAddress, originWallet string, hash string) (error, bool) {
	var (
		amountDstToken  *big.Int
		amountRizToken  *big.Int
		filter, updater bson.M

		err                error
		quantitySwap       float64
		quantitySwapStr    string
		quantityInt        int64
		originRizSwap      float64
		quantityRizSwap    float64
		quantityRizSwapStr string
		slippage           = "0.02"

		originSwap, tokenPrice, rizTokenPrice, priceImpact float64
		txHash                                             string

		approveAmount         *big.Int
		approveAmountRizToken *big.Int
		decimalMultiplierInt  *big.Int
		wg                    sync.WaitGroup
		errCatch              error
		mkpRagentUserBuy      = &marketplaceDto.MkpRagentUserBuy{}
		// mkpSetting            *marketplaceDto.MkpSetting
		mkpUserBuy *marketplaceDto.MkpRagentUserBuy
		mkpRagent  = &marketplaceDto.MkpRagent{}
		mkpSetting = &marketplaceDto.MkpSetting{}

		accountToAmountProjectToken = make(map[string]*big.Int)
		accountToAmountRizToken     = make(map[string]*big.Int)

		processByWallet = make(map[string]string)
	)

	// get ragent
	agentId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return err, true
	}

	quantityInt, err = strconv.ParseInt(quantity, 10, 64)
	if err != nil {
		return err, true
	}

	wg.Add(5)

	go func() {
		defer wg.Done()

		// split amount to all account
		err = s.colMkpMyAgent.FindOne(ctx, bson.M{"tx_hash": hash}).Decode(&mkpRagentUserBuy)
		if err != nil {
			errCatch = err
			return
		}
	}()

	go func() {
		defer wg.Done()

		// get mkp_user_bid to check which step user is in
		filter := bson.M{"tx_hash": hash}
		err := s.colMkpMyAgent.FindOne(ctx, filter).Decode(&mkpUserBuy)
		if err != nil {
			errCatch = err
			return
		}
		if mkpUserBuy.TransferUserProjectTokenTxHashMap == nil {
			mkpUserBuy.TransferUserProjectTokenTxHashMap = make(map[string]string)
		}
	}()

	go func() {
		defer wg.Done()

		err := s.colMkpRAgent.FindOne(ctx, bson.M{"_id": agentId}).Decode(&mkpRagent)
		if err != nil {
			errCatch = err
			return
		}
	}()

	go func() {
		defer wg.Done()

		// get token price
		_, tokenPrice, rizTokenPrice, priceImpact, err = s.helperGetRAgentEstValue(ctx, ragent_id)
		if err != nil {
			errCatch = err
			return
		}
	}()

	go func() {
		defer wg.Done()

		mkpSetting, err = s.helperGetMkpSetting(ctx)
		if err != nil {
			errCatch = err
			return
		}
	}()

	wg.Wait()
	if errCatch != nil {
		return errCatch, true
	}

	slippageNumber := math.Abs(priceImpact)
	if slippageNumber < 1 {
		slippageNumber = 1
	}
	slippageNumber = slippageNumber / 100
	slippage = strconv.FormatFloat(slippageNumber, 'f', -1, 64)

	// update to db all info
	filter = bson.M{"tx_hash": hash}
	updater = bson.M{"$push": bson.M{
		"slippage":        slippageNumber,
		"price_impact":    priceImpact,
		"token_price":     tokenPrice,
		"riz_token_price": rizTokenPrice,
	}}
	_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
	if err != nil {
		return err, true
	}

	// get amount need to transfer for each account by config to make sure each ragent have same amount of token
	projectTokenBigFloat := new(big.Float).SetFloat64(mkpRagent.ProjectToken)
	decimalMultiplier := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(mkpRagent.TokenDecimals)), nil))
	result := new(big.Float).Mul(projectTokenBigFloat, decimalMultiplier)
	amountDstToken, _ = result.Int(nil)
	amountRizToken = big.NewInt(int64(mkpRagent.RizToken * 1e8))

	// TODO: for testing
	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "local" {
		filter = bson.M{"tx_hash": hash}
		updater = bson.M{"$set": bson.M{"status": "swap_token"}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, true
		}

		// transfer riz to account
		for acc, _ := range mkpRagentUserBuy.AccountList {
			nodeClient, err := nodeService.NewNodeService(s.server)
			if err != nil {
				return err, true
			}

			_, err = nodeClient.TransferERC20(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey, s.server.ExtendConfig.RIZAddress, acc, amountDstToken)
			if err != nil {
				return err, true
			}

			_, err = nodeClient.TransferERC20(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey, s.server.ExtendConfig.RIZAddress, acc, amountRizToken)
			if err != nil {
				return err, true
			}
		}

		goto skip_swap_token
	}

	///// process project token swap
	if mkpUserBuy.SwapProjectTokenHash == "" {
		filter = bson.M{"tx_hash": hash}
		updater = bson.M{"$set": bson.M{"status": "find_routing"}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, false
		}

		start := time.Now()
		additionalSlippage := 0.005
		isSuccess := false
		for i := 0; i < 6; i++ {
			// generate should be the quantity of token need to swap = amount token * number of nft token
			originSwap = mkpRagent.ProjectToken * tokenPrice * 1e6                                                                                          // swap 1 token
			quantitySwap = float64(quantityInt) * originSwap                                                                                                // swap quantity token
			quantitySwapStr = strconv.FormatFloat(math.Ceil(quantitySwap*(1+slippageNumber+mkpSetting.SwapFee+additionalSlippage*float64(i))), 'f', -1, 64) // add slippage
			fromWallet, toWallet, projectSwapHash, _, err := s.helperSwapToken(ctx, "buy", s.server.ExtendConfig.USDCAddress, int(define.BASE), mkpRagent.TokenAddress, int(mkpRagent.ChainId), slippage, quantitySwapStr, 6, mkpRagent.TokenDecimals, big.NewFloat(mkpRagent.ProjectToken*float64(quantityInt)), true, nil)
			if err != nil {
				isGotRefund := false
				if err.Error() == "REFUNDED" {
					isGotRefund = true
				}
				isSuccess = false
				filter = bson.M{"tx_hash": hash}
				updater = bson.M{"$push": bson.M{"quantity_swap_project_token": quantitySwapStr}, "$set": bson.M{"is_got_refund": isGotRefund}}
				_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
				if err != nil {
					return err, false
				}
				continue
			}

			filter = bson.M{"tx_hash": hash}
			updater = bson.M{"$set": bson.M{"status": "swap_project", "swap_project_token_hash": projectSwapHash, "swap_project_token_time": time.Since(start).Seconds(), "swap_project_token_process_by_wallet": fromWallet, "swap_project_token_to_wallet": toWallet}, "$push": bson.M{"quantity_swap_project_token": quantitySwapStr}}
			_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
			if err != nil {
				return err, false
			}

			processByWallet["project_token"] = fromWallet
			processByWallet["project_token_to_wallet"] = toWallet

			isSuccess = true
			break
		}

		if !isSuccess {
			return errors.New("swap project token failed"), true
		}
	} else {
		processByWallet["project_token"] = mkpRagentUserBuy.SwapProjectTokenProcessByWallet
		processByWallet["project_token_to_wallet"] = mkpRagentUserBuy.SwapProjectTokenToWallet
	}

	///// process riz token swap
	if mkpUserBuy.SwapRizTokenHash == "" {
		start := time.Now()
		additionalSlippage := 0.005
		isSuccess := false
		for i := 0; i < 6; i++ {
			// generate should be the quantity of token need to swap = amount token * number of nft token
			originRizSwap = mkpRagent.RizToken * rizTokenPrice * 1e6                                                                                              // swap 1 token from usdc
			quantityRizSwap = float64(quantityInt) * originRizSwap                                                                                                // swap quantity token
			quantityRizSwapStr = strconv.FormatFloat(math.Ceil(quantityRizSwap*(1+slippageNumber+mkpSetting.SwapFee+additionalSlippage*float64(i))), 'f', -1, 64) // add slippage

			fromWallet, _, rizSwapHash, _, err := s.helperSwapToken(ctx, "buy", s.server.ExtendConfig.USDCAddress, int(define.BASE), s.server.ExtendConfig.RIZAddress, int(define.BASE), slippage, quantityRizSwapStr, 6, 8, big.NewFloat(mkpRagent.RizToken*float64(quantityInt)), true, nil)
			if err != nil {
				isSuccess = false
				filter = bson.M{"tx_hash": hash}
				updater = bson.M{"$push": bson.M{"quantity_swap_riz_token": quantityRizSwapStr}}
				_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
				if err != nil {
					return err, false
				}
				continue
			}

			processByWallet["riz_token"] = fromWallet

			filter = bson.M{"tx_hash": hash}
			updater = bson.M{"$set": bson.M{"status": "swap_riz", "swap_riz_token_hash": rizSwapHash, "swap_riz_token_time": time.Since(start).Seconds(), "swap_riz_token_process_by_wallet": fromWallet}, "$push": bson.M{"quantity_swap_riz_token": quantityRizSwapStr}}
			_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
			if err != nil {
				return err, false
			}

			isSuccess = true
			break
		}

		if !isSuccess {
			return errors.New("swap riz token failed"), true
		}
	} else {
		processByWallet["riz_token"] = mkpRagentUserBuy.SwapRizTokenProcessByWallet
	}

	// get mapping account to amount
	for account, tokenId := range mkpRagentUserBuy.AccountList {
		if mkpRagent.ChainId == int(define.BASE) {
			accountToAmountProjectToken[account] = amountDstToken
		} else {
			accountToAmountProjectToken[mkpRagentUserBuy.AccountListBase[tokenId]] = amountDstToken
		}
		accountToAmountRizToken[account] = amountRizToken
	}

	// approve token to batch transfer contract
	if mkpUserBuy.TransferProjectTokenTxHash == "" {
		approveAmount = new(big.Int).Mul(amountDstToken, big.NewInt(int64(len(mkpRagentUserBuy.AccountList))))
		decimalMultiplierInt = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(mkpRagent.TokenDecimals)), nil) // add extra 10^decimals to approve amount to make sure approve amount is enough
		approveAmount = approveAmount.Add(approveAmount, decimalMultiplierInt)

		start := time.Now()

		txHash, err = s.helperBatchTransferERC20(ctx, processByWallet["project_token"], processByWallet["project_token_to_wallet"], mkpRagent.TokenAddress, mkpRagent.TokenDecimals, int(mkpRagent.ChainId), approveAmount, accountToAmountProjectToken, mkpUserBuy.TransferUserProjectTokenTxHashMap)
		if err != nil {
			return err, true
		}

		// update mkp_user_buy
		filter = bson.M{"tx_hash": hash}
		updater = bson.M{"$set": bson.M{"status": "transfer_project", "transfer_project_token_tx_hash": txHash, "transfer_user_project_token_tx_hash_map": mkpUserBuy.TransferUserProjectTokenTxHashMap, "updatedAt": time.Now().UTC(), "batch_transfer_project_token_time": time.Since(start).Seconds()}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, true
		}
	}

	// approve riz token to batch transfer contract
	if mkpUserBuy.TransferRizTokenTxHash == "" {
		approveAmountRizToken = new(big.Int).Mul(amountRizToken, big.NewInt(int64(len(mkpRagentUserBuy.AccountList))))
		decimalMultiplierInt = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(8)), nil) // add extra 10^8 to approve amount to make sure approve amount is enough
		approveAmountRizToken = approveAmountRizToken.Add(approveAmountRizToken, decimalMultiplierInt)

		start := time.Now()

		txHash, err = s.helperBatchTransferERC20(ctx, processByWallet["riz_token"], "", s.server.ExtendConfig.RIZAddress, 8, int(define.BASE), approveAmountRizToken, accountToAmountRizToken, mkpUserBuy.TransferUserProjectTokenTxHashMap)
		if err != nil {
			return err, true
		}

		// update mkp_user_buy
		filter = bson.M{"tx_hash": hash}
		updater = bson.M{"$set": bson.M{"status": "transfer_riz", "transfer_riz_token_tx_hash": txHash, "updatedAt": time.Now().UTC(), "batch_transfer_riz_token_time": time.Since(start).Seconds()}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, true
		}
	}

skip_swap_token:

	// update status
	filter = bson.M{"tx_hash": hash}
	updater = bson.M{"$set": bson.M{"status": "success", "updatedAt": time.Now().UTC()}}
	_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
	if err != nil {
		return err, true
	}

	///// update mkp_user
	// find mkp_user
	filter = bson.M{"wallet": walletAddress, "ragent_id": ragent_id}
	var mkpUser *marketplaceDto.MkpUser
	err = s.colMkpUser.FindOne(ctx, filter).Decode(&mkpUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			mkpUser = &marketplaceDto.MkpUser{
				ID:              primitive.NewObjectID(),
				Wallet:          walletAddress,
				RagentId:        ragent_id,
				AccountList:     mkpRagentUserBuy.AccountList,
				AccountListBase: mkpRagentUserBuy.AccountListBase,
				Quantity:        quantityInt,
				CreatedAt:       time.Now().UTC(),
				UpdatedAt:       time.Now().UTC(),
			}
			_, err = s.colMkpUser.InsertOne(ctx, mkpUser)
			if err != nil {
				return err, true
			}
			return nil, false
		}
		return err, true
	} else {
		// get new account list
		var newAccountList = make(map[string]string)
		if mkpUser.AccountList != nil {
			newAccountList = mkpUser.AccountList
		}
		for acc, tokenId := range mkpRagentUserBuy.AccountList {
			newAccountList[acc] = tokenId
		}

		var newAccountListBase = make(map[string]string)
		if mkpUser.AccountListBase != nil {
			newAccountListBase = mkpUser.AccountListBase
		}
		for tokenId, acc := range mkpRagentUserBuy.AccountListBase {
			newAccountListBase[tokenId] = acc
		}
		// increse quantity
		filter = bson.M{"wallet": walletAddress, "ragent_id": ragent_id}
		update := bson.M{"$set": bson.M{"account_list": newAccountList, "account_list_base": newAccountListBase, "quantity": len(newAccountList), "updatedAt": time.Now().UTC()}}
		_, err = s.colMkpUser.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Error("Error update mkp_user: "+err.Error(), "MarketPlaceVerify-UpdateMkpUser")
			return err, true
		}
	}

	return nil, false
}

// TODO: if fromChainId and toChainId is different, need to use coinbase to swap on lifi
func (s *MarketplaceService) helperSwapToken(ctx context.Context, side string, fromTokenAddress string, fromChainId int, toTokenAddress string, toChainId int, slippage, quantityRizSwapStr string, fromTokenDecimals, toTokenDecimals int, requiredAmount *big.Float, isNeedCheckMinReceiveAmount bool, walletSwapUsed *marketplaceDto.MkpSwapWallet) (string, string, string, *marketplaceDto.OkxSwapResponse, error) {
	var (
		walletInfo        *marketplaceDto.MkpSwapWallet
		toWallet          string
		err               error
		isUseKeeperWallet = false
	)

	// check if need to use keeper wallet
	// if money reach limit balance swap wallet, use keeper wallet to swap
	if fromTokenAddress == s.server.ExtendConfig.USDCAddress {
		swapQuantity, ok := big.NewFloat(0).SetString(quantityRizSwapStr)
		if !ok {
			return "", "", "", nil, errors.New("invalid quantity")
		}

		if swapQuantity.Cmp(big.NewFloat(float64(s.server.ExtendConfig.LimitBalanceSwapWallet*1e6))) >= 0 {
			isUseKeeperWallet = true
		}
	}

	if isUseKeeperWallet {
		walletAddress, err := s.helperGetAddressFromPrivateKey(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey)
		if err != nil {
			return "", "", "", nil, err
		}
		walletInfo = &marketplaceDto.MkpSwapWallet{
			Secret: s.server.ExtendConfig.RagentWalletKeeperPrivateKey,
			Wallet: walletAddress,
		}

	} else {
		if walletSwapUsed != nil {
			walletInfo = walletSwapUsed
		} else {
			walletInfo, err = s.helperGetSwapWallet(ctx, fromChainId)
			if err != nil {
				return "", "", "", nil, err
			}
			defer s.helperReleaseSwapWallet(ctx, walletInfo.Wallet, fromChainId)
		}
	}

	fromWallet := walletInfo.Wallet
	toWallet = walletInfo.Wallet

	var (
		swapHash string
		callData *marketplaceDto.OkxSwapResponse
	)

	// update wallet status back to NONE
	if !isUseKeeperWallet && walletSwapUsed == nil {
		defer s.helperReleaseSwapWallet(ctx, fromWallet, fromChainId)
	}

	multiplier := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(toTokenDecimals)), nil))
	requiredAmountBigFloat := new(big.Float).Mul(requiredAmount, multiplier)
	requiredAmountBigInt, _ := requiredAmountBigFloat.Int(nil)
	requiredAmountStr := requiredAmountBigInt.String()
	if fromChainId != int(define.BASE) || toChainId != int(define.BASE) {
		var lifiResponse *marketplaceDto.LiFiQuoteResponse
		swapAmount, ok := new(big.Int).SetString(quantityRizSwapStr, 10)
		if !ok {
			return "", "", "", nil, errors.New("invalid from amount")
		}
		var isEnough = false
		for i := 0; i < 20; i++ {
			lifiRequest := &marketplaceDto.LiFiQuoteRequest{
				ToAmount:  requiredAmountStr,
				FromToken: fromTokenAddress,
				ToToken:   toTokenAddress,
				FromChain: fromChainId,
				ToChain:   toChainId,
				ContractCalls: []marketplaceDto.LiFiContractCall{
					{
						FromTokenAddress:   fromTokenAddress,
						ToContractAddress:  toTokenAddress,
						ToContractCallData: "0x",
						ToContractGasLimit: "200000",
					},
				},
				FromAddress: fromWallet,
				ToAddress:   toWallet,
			}
			if toChainId == int(define.SOLANA) {
				// get wallet from swap wallet
				toWalletInfo, err := s.helperGetSwapWallet(ctx, int(define.SOLANA))
				if err != nil {
					return "", "", "", nil, err
				}
				s.helperReleaseSwapWallet(ctx, toWalletInfo.Wallet, int(define.SOLANA))

				toWallet = toWalletInfo.Wallet
				lifiRequest.ToAddress = toWallet
				lifiRequest.ToChain = 1151111081099710
			}
			if fromChainId == int(define.SOLANA) {
				// get wallet from swap wallet
				toWalletInfo, err := s.helperGetSwapWallet(ctx, int(define.BASE))
				if err != nil {
					return "", "", "", nil, err
				}
				s.helperReleaseSwapWallet(ctx, toWalletInfo.Wallet, int(define.BASE))

				toWallet = toWalletInfo.Wallet
				lifiRequest.ToAddress = toWallet
				lifiRequest.FromChain = 1151111081099710
			}
			lifiResponse, err = httpClient.GetLiFiQuoteByToAmount(lifiRequest)
			if err != nil {
				log.Error("Error get lifi quote: "+err.Error(), "MarketPlaceVerify-helperSwapToken", map[string]interface{}{
					"fromTokenAddress": fromTokenAddress,
					"toTokenAddress":   toTokenAddress,
					"fromChainId":      fromChainId,
					"toChainId":        toChainId,
					"requiredAmount":   requiredAmountStr,
				})
				time.Sleep(time.Millisecond * 100)
				continue
			}

			fromAmount, ok := new(big.Int).SetString(lifiResponse.Estimate.FromAmount, 10)
			if !ok {
				return "", "", "", nil, errors.New("invalid from amount")
			}
			if side == "sell" && fromAmount.Cmp(swapAmount) > 0 {
				// Calculate adjustment based on token decimals
				// For 18 decimals: 500000, for 6 decimals: 500 (proportionally smaller)
				// decimalAdjustment := big.NewInt(500000)

				// percent = float64(fromAmount - swapAmount) / float64(swapAmount)
				// adjustment = requiredAmount * (1 - percent)

				subSwap := new(big.Int).Sub(fromAmount, swapAmount)
				percentDifference := new(big.Float).Quo(new(big.Float).SetInt(subSwap), new(big.Float).SetInt(swapAmount))
				percentDifference = percentDifference.Sub(big.NewFloat(1), percentDifference)
				adjustment := new(big.Float).Mul(new(big.Float).SetInt(requiredAmountBigInt), percentDifference)
				adjustment = adjustment.Sub(adjustment, big.NewFloat(500*float64(i)))
				requiredAmountBigInt, _ = adjustment.Int(nil)

				requiredAmountStr = requiredAmountBigInt.String()
				time.Sleep(time.Millisecond * 200)
				continue
			}

			isEnough = true
			break
		}
		// fmt.Println(swapAmount)
		if lifiResponse == nil || !isEnough {
			return "", "", "", nil, errors.New("lifi response is nil")
		}

		if fromChainId == int(define.SOLANA) {
			// map to this to use later on sell agent
			callData = &marketplaceDto.OkxSwapResponse{
				Tx: marketplaceDto.OkxTx{
					Data: lifiResponse.TransactionRequest.Data,
					From: fromWallet,
					To:   lifiResponse.TransactionRequest.To,
				},
				RouterResult: marketplaceDto.RouterResult{
					ToTokenAmount: lifiResponse.Estimate.ToAmountMin,
				},
			}

			// approveAddress := lifiResponse.IncludedSteps[0].Estimate.ApprovalAddress
			// approveAmount, ok := big.NewInt(0).SetString(lifiResponse.IncludedSteps[0].Estimate.FromAmount, 10)
			// if !ok {
			// 	return "", "", "", nil, errors.New("invalid approve amount")
			// }
			nodeClientSolana, err := nodeService.NewNodeService(s.server, int(define.SOLANA))
			if err != nil {
				return "", "", "", nil, err
			}
			// _, err = nodeClientSolana.ApproveSPLToken(ctx, fromWallet, fromTokenAddress, approveAddress, approveAmount.Uint64())
			// if err != nil {
			// 	return "", "", "", nil, err
			// }

			// Get transaction data from LiFi response
			txData := lifiResponse.TransactionRequest.Data
			if txData == "" {
				return "", "", "", nil, errors.New("empty transaction data from LiFi")
			}

			// Decode the base64 transaction data
			txBytes, err := base64.StdEncoding.DecodeString(txData)
			if err != nil {
				return "", "", "", nil, fmt.Errorf("failed to decode transaction data: %v", err)
			}

			// Parse the transaction from bytes
			tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(txBytes))
			if err != nil {
				return "", "", "", nil, fmt.Errorf("failed to parse transaction: %v", err)
			}

			// Get fresh blockhash to prevent "Blockhash not found" error
			recentBlockhash, err := nodeClientSolana.GetLatestBlockhashSolana(ctx)
			if err != nil {
				return "", "", "", nil, fmt.Errorf("failed to get latest blockhash: %v", err)
			}

			// Update transaction with fresh blockhash
			tx.Message.RecentBlockhash = recentBlockhash.Value.Blockhash

			// Get private key from wallet
			privateKey, err := solana.PrivateKeyFromBase58(walletInfo.Secret)
			if err != nil {
				return "", "", "", nil, fmt.Errorf("invalid private key: %v", err)
			}

			// Clear existing signatures since we're updating the blockhash
			tx.Signatures = make([]solana.Signature, int(tx.Message.Header.NumRequiredSignatures))

			// Sign the transaction with the fresh blockhash
			_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
				if key.Equals(privateKey.PublicKey()) {
					return &privateKey
				}
				return nil
			})
			if err != nil {
				return "", "", "", nil, fmt.Errorf("failed to sign transaction: %v", err)
			}

			// Serialize signed transaction to bytes
			signedTxBytes, err := tx.MarshalBinary()
			if err != nil {
				return "", "", "", nil, fmt.Errorf("marshal signed tx error: %v", err)
			}

			// Send transaction to Solana network
			swapHash, err = nodeClientSolana.SendSolanaTransaction(ctx, signedTxBytes)
			if err != nil {
				if strings.Contains(err.Error(), "This transaction has already been processed") {
					time.Sleep(time.Millisecond * 300)
				} else {
					return "", "", "", nil, fmt.Errorf("send solana transaction error: %v", err)
				}
			}

		} else {
			gasLimit, ok := big.NewInt(0).SetString(lifiResponse.TransactionRequest.GasLimit[2:], 16)
			if !ok {
				return "", "", "", nil, errors.New("invalid gas limit")
			}
			gasPrice, ok := big.NewInt(0).SetString(lifiResponse.TransactionRequest.GasPrice[2:], 16)
			if !ok {
				return "", "", "", nil, errors.New("invalid gas price")
			}
			value, ok := big.NewInt(0).SetString(lifiResponse.TransactionRequest.Value[2:], 16)
			if !ok {
				return "", "", "", nil, errors.New("invalid value")
			}
			gasLimitStr := gasLimit.String()
			gasPriceStr := gasPrice.String()
			valueStr := value.String()

			// map to this to use later on sell agent
			callData = &marketplaceDto.OkxSwapResponse{
				Tx: marketplaceDto.OkxTx{
					Data:     lifiResponse.TransactionRequest.Data,
					Gas:      gasLimitStr,
					GasPrice: gasPriceStr,
					From:     fromWallet,
					To:       lifiResponse.TransactionRequest.To,
					Value:    valueStr,
				},
				RouterResult: marketplaceDto.RouterResult{
					ToTokenAmount: lifiResponse.Estimate.ToAmountMin,
				},
			}

			minReceiveAmountFloat, ok := big.NewFloat(0).SetString(lifiResponse.Estimate.ToAmountMin)
			if !ok {
				return "", "", "", nil, err
			}
			toTokenDecimals := int64(lifiResponse.Action.ToToken.Decimals)
			decimalMultiplier := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(toTokenDecimals), nil))
			minReceiveAmountFloat = minReceiveAmountFloat.Quo(minReceiveAmountFloat, decimalMultiplier)
			if isNeedCheckMinReceiveAmount && minReceiveAmountFloat.Cmp(requiredAmount) < 0 {
				return "", "", "", nil, errors.New("min receive token is not enough")
			}

			// approve tx
			nodeClient, err := nodeService.NewNodeService(s.server, fromChainId)
			if err != nil {
				return "", "", "", nil, err
			}
			approveAmountInt, ok := big.NewInt(0).SetString(lifiResponse.Estimate.FromAmount, 10)
			if !ok {
				return "", "", "", nil, err
			}
			_, err = nodeClient.ApproveERC20(ctx, walletInfo.Secret, lifiResponse.Estimate.ApprovalAddress, fromTokenAddress, approveAmountInt)
			if err != nil {
				return "", "", "", nil, err
			}

			swapHash, err = nodeClient.BroadcastTransaction(ctx, fromWallet, callData, walletInfo.Secret)
			if err != nil {
				return "", "", "", nil, err
			}
		}

		// wait for transaction to be confirmed
		// TODO: add check for all cases can happen with lifi
		for i := 0; i < 20; i++ {
			fromChainCall := fromChainId
			toChainCall := toChainId
			if fromChainId == int(define.SOLANA) {
				fromChainCall = 1151111081099710
			}
			if toChainId == int(define.SOLANA) {
				toChainCall = 1151111081099710
			}
			lifiStatus, err := httpClient.GetLiFiStatus(ctx, swapHash, fromChainCall, toChainCall)
			if err == nil {
				if lifiStatus.Status == "DONE" {
					if lifiStatus.SubStatus != "COMPLETED" {
						return "", "", "", nil, errors.New(lifiStatus.SubStatus)
					}
					break
				}
			}

			time.Sleep(time.Millisecond * 300)
		}

	} else {
		callData, err = httpClient.GetOkxSwap(fromTokenAddress, toTokenAddress, quantityRizSwapStr, slippage, fromWallet)
		if err != nil {
			return "", "", "", nil, err
		}

		// calculate min receive amount
		minReceiveAmountFloat, ok := big.NewFloat(0).SetString(callData.Tx.MinReceiveAmount)
		if !ok {
			return "", "", "", nil, err
		}
		toTokenDecimals, err := strconv.ParseInt(callData.RouterResult.ToToken.Decimal, 10, 64)
		if err != nil {
			return "", "", "", nil, err
		}
		decimalMultiplier := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(toTokenDecimals), nil))
		minReceiveAmountFloat = minReceiveAmountFloat.Quo(minReceiveAmountFloat, decimalMultiplier)
		if isNeedCheckMinReceiveAmount && minReceiveAmountFloat.Cmp(requiredAmount) < 0 {
			return "", "", "", nil, errors.New("min receive token is not enough")
		}

		// approve tx
		approveData, err := httpClient.OkxApproveTransaction(fromTokenAddress, quantityRizSwapStr)
		if err != nil {
			fmt.Println("approveData: ", err)
			return "", "", "", nil, err
		}

		nodeClient, err := nodeService.NewNodeService(s.server)
		if err != nil {
			return "", "", "", nil, err
		}

		_, err = nodeClient.BroadcastTransaction(ctx, fromWallet, &marketplaceDto.OkxSwapResponse{
			Tx: marketplaceDto.OkxTx{
				Data:     approveData.Data,
				Gas:      approveData.GasLimit,
				GasPrice: approveData.GasPrice,
				From:     fromWallet,
				To:       fromTokenAddress,
			},
		}, walletInfo.Secret)
		if err != nil {
			return "", "", "", nil, err
		}

		swapHash, err = nodeClient.BroadcastTransaction(ctx, fromWallet, callData, walletInfo.Secret)
		if err != nil {
			return "", "", "", nil, err
		}
	}

	// push to kafka
	if fromTokenAddress == s.server.ExtendConfig.USDCAddress {
		event1 := event.Event{
			EventName: "marketplace|consume-swap-wallet",
			EventData: map[string]interface{}{
				"wallet":                  fromWallet,
				"consumed_token":          fromTokenAddress,
				"consumed_token_decimals": fromTokenDecimals,
			},
		}
		if s.server.Pub["task"] != nil {
			errp := s.server.Pub["task"].Publish(event1)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceConsumeSwapWallet-pushToKafka", event1)
			}
		}
	}

	return fromWallet, toWallet, swapHash, callData, nil
}

func (s *MarketplaceService) helperGetSwapWallet(ctx context.Context, chainId int) (*marketplaceDto.MkpSwapWallet, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("get swap wallet timeout")
		default:
		}

		filter := bson.M{"process_by": "NONE", "status": "active", "chain_id": chainId}
		var mkpSwapWallet *marketplaceDto.MkpSwapWallet
		err := s.colMkpSwapWallet.FindOneAndUpdate(ctx, filter, bson.M{"$set": bson.M{"process_by": "BUSY"}}).Decode(&mkpSwapWallet)
		if err != nil {
			time.Sleep(time.Millisecond * 200)
			continue
		}

		return mkpSwapWallet, nil
	}
}

func (s *MarketplaceService) helperGetSwapWalletByAddress(ctx context.Context, walletAddress string, chainId int) (*marketplaceDto.MkpSwapWallet, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("get swap wallet timeout")
		default:
		}

		filter := bson.M{"wallet": walletAddress, "process_by": "NONE", "chain_id": chainId}
		var mkpSwapWallet *marketplaceDto.MkpSwapWallet
		err := s.colMkpSwapWallet.FindOneAndUpdate(ctx, filter, bson.M{"$set": bson.M{"process_by": "BUSY"}}).Decode(&mkpSwapWallet)
		if err != nil {
			time.Sleep(time.Millisecond * 200)
			continue
		}

		return mkpSwapWallet, nil
	}
}

func (s *MarketplaceService) helperReleaseSwapWallet(ctx context.Context, walletAddress string, chainId int) error {
	filter := bson.M{"wallet": walletAddress, "chain_id": chainId}
	updater := bson.M{"$set": bson.M{"process_by": "NONE"}}
	_, err := s.colMkpSwapWallet.UpdateOne(ctx, filter, updater)
	if err != nil {
		return err
	}
	return nil
}

// TODO: make this multichain
func (s *MarketplaceService) helperBatchTransferERC20(ctx context.Context, walletAddress string, toWallet string, tokenAddress string, tokenDecimals int, toChainId int, approveAmount *big.Int, accountToAmountToken map[string]*big.Int, mapAccountToTxHash map[string]string) (string, error) {
	nodeClient, err := nodeService.NewNodeService(s.server, toChainId)
	if err != nil {
		return "", err
	}

	var mkpSwapWallet *marketplaceDto.MkpSwapWallet
	keeperWalletAddress, err := s.helperGetAddressFromPrivateKey(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey)
	if err != nil {
		return "", err
	}
	if walletAddress == keeperWalletAddress && toChainId != int(define.SOLANA) {
		mkpSwapWallet = &marketplaceDto.MkpSwapWallet{
			Secret: s.server.ExtendConfig.RagentWalletKeeperPrivateKey,
			Wallet: walletAddress,
		}
	} else {
		if toWallet != "" {
			mkpSwapWallet, err = s.helperGetSwapWalletByAddress(ctx, toWallet, toChainId)
			if err != nil {
				return "", err
			}
			defer s.helperReleaseSwapWallet(ctx, toWallet, toChainId)
		} else {
			mkpSwapWallet, err = s.helperGetSwapWalletByAddress(ctx, walletAddress, toChainId)
			if err != nil {
				return "", err
			}
			defer s.helperReleaseSwapWallet(ctx, walletAddress, toChainId)
		}
	}

	batchTransferContract := s.server.ExtendConfig.BatchTransferContract
	if toChainId == int(define.BSC) {
		batchTransferContract = s.server.ExtendConfig.BatchTransferContractBsc
	} else if toChainId == int(define.ETHEREUM) {
		batchTransferContract = s.server.ExtendConfig.BatchTransferContractEth
	} else if toChainId == int(define.SOLANA) {

		nodeClientSolana, err := nodeService.NewNodeService(s.server, toChainId)
		if err != nil {
			return "", err
		}

		// transfer for each account
		for account, amount := range accountToAmountToken {

			if mapAccountToTxHash[account] != "" {
				continue
			}

			from, err := solana.PrivateKeyFromBase58(mkpSwapWallet.Secret)
			if err != nil {
				return "", err
			}
			fromPub := from.PublicKey()

			toPub, err := solana.PublicKeyFromBase58(account)
			if err != nil {
				return "", err
			}

			mint := solana.MustPublicKeyFromBase58(tokenAddress)

			// Convert amount properly - don't multiply by LAMPORTS_PER_SOL for token transfers
			// The amount should already be in the correct token decimals
			tokenAmount := uint64(amount.Int64())

			senderTokenAccount, _, _ := solana.FindAssociatedTokenAddress(fromPub, mint)
			recipientTokenAccount, _, _ := solana.FindAssociatedTokenAddress(toPub, mint)

			// Create instructions slice
			var instructions []solana.Instruction

			excludeChain := map[string]bool{}
			node, _ := nodeClientSolana.GetNode(excludeChain)
			rpcClient := rpc.New(node.URL)
			_, err = rpcClient.GetAccountInfo(ctx, recipientTokenAccount)
			if err != nil {
				// Create ATA if it doesn't exist
				createATAIx := associatedtokenaccount.NewCreateInstruction(
					fromPub, // payer
					toPub,   // wallet
					mint,    // mint
				).Build()
				instructions = append(instructions, createATAIx)
			}

			_, err = rpcClient.GetAccountInfo(ctx, senderTokenAccount)
			if err != nil {
				// Sender account doesn't exist, this is an error
				return "", fmt.Errorf("sender token account does not exist: %v", err)
			}

			// Create transfer instruction
			ix := token.NewTransferCheckedInstruction(
				tokenAmount,
				uint8(tokenDecimals),
				senderTokenAccount,
				mint,
				recipientTokenAccount,
				fromPub,
				nil,
			).Build()
			instructions = append(instructions, ix)

			recent, err := nodeClientSolana.GetLatestBlockhashSolana(ctx)
			if err != nil {
				return "", err
			}

			tx, err := solana.NewTransaction(
				instructions,
				recent.Value.Blockhash,
				solana.TransactionPayer(fromPub),
			)
			if err != nil {
				return "", err
			}

			// Sign transaction with private key
			_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
				if key.Equals(fromPub) {
					return &from
				}
				return nil
			})
			if err != nil {
				return "", err
			}

			// Serialize signed transaction to bytes
			txBytes, err := tx.MarshalBinary()
			if err != nil {
				return "", fmt.Errorf("marshal signed tx error: %v", err)
			}

			// Send transaction to Solana network
			txHash, err := nodeClientSolana.SendSolanaTransaction(ctx, txBytes)
			if err != nil {
				return "", fmt.Errorf("send solana transaction error: %v", err)
			}

			mapAccountToTxHash[account] = txHash
			fmt.Printf("Solana transaction sent successfully. TxHash: %s\n", txHash)
		}

		return "", nil
	}

	_, err = nodeClient.ApproveERC20(ctx, mkpSwapWallet.Secret, batchTransferContract, tokenAddress, approveAmount)
	if err != nil {
		return "", err
	}

	// transfer riz token to account
	txHash, err := nodeClient.BatchTransferERC20(ctx, mkpSwapWallet.Secret, batchTransferContract, tokenAddress, accountToAmountToken)
	if err != nil {
		return "", err
	}

	return txHash, nil
}
