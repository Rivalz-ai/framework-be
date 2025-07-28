package service

import (
	"context"
	"errors"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/framework/log"
	marketplaceDto "github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	nodeService "github.com/Rivalz-ai/framework-be/modules/node/service"
	rewardService "github.com/Rivalz-ai/framework-be/modules/reward/service"
	userDto "github.com/Rivalz-ai/framework-be/modules/user/dto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// deprecated
func (s *MarketplaceService) SellRagent(ctx context.Context, wallet_address, ragent_id string, tokenIds []int64) (string, error) {
	quantity := int64(len(tokenIds))

	// get mkp_user
	filter := bson.M{"wallet": wallet_address, "ragent_id": ragent_id}
	var mkpUser *marketplaceDto.MkpUser
	err := s.colMkpUser.FindOne(ctx, filter).Decode(&mkpUser)
	if err != nil {
		return "", err
	}

	quantityInfo, _, err := s.GetQuantityToken(ctx, ragent_id, quantity, "sell", wallet_address)
	if err != nil {
		return "", err
	}

	var (
		mapTokenId      = make(map[string]struct{})
		remaingTokenIds = make(map[string]string)
		removedTokenMap = make(map[string]string)
	)
	for _, tokenId := range tokenIds {
		mapTokenId[strconv.FormatInt(tokenId, 10)] = struct{}{}
	}

	for acc, tokenid := range mkpUser.AccountList {
		if _, ok := mapTokenId[tokenid]; !ok {
			remaingTokenIds[acc] = tokenid
		} else {
			removedTokenMap[acc] = tokenid
		}
	}

	myAgentId := primitive.NewObjectID()
	record := &marketplaceDto.MkpRagentUserBuy{
		ID:            myAgentId,
		MkpRagentId:   ragent_id,
		WalletAddress: wallet_address,
		Amount:        strconv.FormatInt(quantity, 10),
		Status:        "pending",
		Side:          "sell",
		TotalPrice:    strconv.FormatFloat(quantityInfo.TotalPrice, 'f', -1, 64),
		RemovedTokens: removedTokenMap,
		AccountList:   remaingTokenIds,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	_, err = s.colMkpMyAgent.InsertOne(ctx, record)
	if err != nil {
		return "", err
	}

	//update back to mkp_user to make sure they don't sell again
	_, err = s.colMkpUser.UpdateOne(ctx, filter, bson.M{
		"$set": bson.M{
			"account_list": remaingTokenIds,
			"updated_at":   time.Now(),
		},
		"$inc": bson.M{
			"quantity": -quantity,
		},
	})
	if err != nil {
		return "", err
	}

	return myAgentId.Hex(), nil
}

func (s *MarketplaceService) InternalSellRagent(ctx context.Context, ragent_id, quantity, wallet_address, trackingId string) (error, bool) {
	trackingIdInt, err := primitive.ObjectIDFromHex(trackingId)
	if err != nil {
		return err, true
	}

	// get ragent info
	agentId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return err, true
	}
	filter := bson.M{"_id": agentId}
	var mkpRagent *marketplaceDto.MkpRagent
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return err, true
	}

	// get mkp user
	err, _, amountTokenInt, amountRizInt, walletSwapUsed, walletBase := s.InternalUnwrapRagent(ctx, ragent_id, quantity, wallet_address, trackingId, true)
	// release swap wallet
	if walletSwapUsed != nil {
		defer s.helperReleaseSwapWallet(ctx, walletSwapUsed.Wallet, mkpRagent.ChainId)
	}

	if err != nil {
		return err, true
	}

	// get mkpuser buy
	filter = bson.M{"_id": trackingIdInt}
	var mkpUserBuy *marketplaceDto.MkpRagentUserBuy
	err = s.colMkpMyAgent.FindOne(ctx, filter).Decode(&mkpUserBuy)
	if err != nil {
		return err, true
	}
	if mkpUserBuy.SellSwapTokenTxHash == nil {
		mkpUserBuy.SellSwapTokenTxHash = make(map[string]string)
	}
	if mkpUserBuy.SellAmountToken == nil {
		mkpUserBuy.SellAmountToken = make(map[string]string)
	}

	// get total amount of token
	// quantityInt, err := strconv.ParseInt(quantity, 10, 64)
	// if err != nil {
	// 	return err, true
	// }
	// amountTokenInt.Mul(amountTokenInt, big.NewInt(quantityInt))
	// amountRizInt.Mul(amountRizInt, big.NewInt(quantityInt))

	nodeClient, err := nodeService.NewNodeService(s.server)
	if err != nil {
		return err, true
	}

	walletAddress, err := s.helperGetAddressFromPrivateKey(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey)
	if err != nil {
		return err, true
	}

	var (
		remainingAmountInt = big.NewInt(0)
		key                = mkpRagent.TokenAddress + "_" + walletAddress
		ok                 bool
		tmp                *big.Int
		updater            bson.M
	)

	mkpSetting, err := s.helperGetMkpSetting(ctx)
	if err != nil {
		return err, true
	}

	// _, slippageFloat, err := s.helperGetTokenPriceOkx(ctx, mkpRagent.TokenAddress, mkpRagent.TokenDecimals)
	// if err != nil {
	// 	slippageFloat = mkpSetting.SlippageRate
	// }

	_, priceProjectToken, _, slippageFloat, err := s.helperGetRAgentEstValue(ctx, ragent_id)
	if err != nil {
		return err, true
	}
	slippageFloat = slippageFloat / 100
	slippageStr := strconv.FormatFloat(slippageFloat, 'f', -1, 64)

	// skip swap token for dev and local
	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "local" {
		goto skip_swap_token
	}

	key = mkpRagent.TokenAddress
	if key == s.server.ExtendConfig.RIZAddress {
		key = s.server.ExtendConfig.USDCAddress + "_project"
	}
	if _, ok = mkpUserBuy.SellSwapTokenTxHash[key]; !ok {
		filter = bson.M{"_id": trackingIdInt}
		updater = bson.M{"$set": bson.M{"status": "find_routing"}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, true
		}

		start := time.Now()

		// fmt.Println("priceProjectToken", priceProjectToken)
		// fmt.Println("amountTokenInt", amountTokenInt)

		decimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(mkpRagent.TokenDecimals)), nil)
		requiredAmount := new(big.Float).Mul(big.NewFloat(priceProjectToken), big.NewFloat(0).SetInt(amountTokenInt))
		requiredAmount.Quo(requiredAmount, big.NewFloat(0).SetInt(decimal))
		// fmt.Println("requiredAmount", requiredAmount)
		processWallet, toAddress, txHash, swapInfo, err := s.helperSwapToken(ctx, "sell", mkpRagent.TokenAddress, int(mkpRagent.ChainId), s.server.ExtendConfig.USDCAddress, int(define.BASE), slippageStr, amountTokenInt.String(), mkpRagent.TokenDecimals, 6, requiredAmount, false, walletSwapUsed)
		if err != nil {
			return err, true
		}
		if mkpRagent.ChainId == int(define.SOLANA) {
			// get wallet from swap wallet
			var swapWallet *marketplaceDto.MkpSwapWallet
			err = s.colMkpSwapWallet.FindOne(ctx, bson.M{"wallet": toAddress}).Decode(&swapWallet)
			if err != nil {
				return err, true
			}
			walletSwapUsed.Wallet = toAddress
			walletSwapUsed.Secret = swapWallet.Secret
			processWallet = toAddress
		}

		mkpUserBuy.SellSwapTokenTxHash[key] = txHash
		tmp, ok = new(big.Int).SetString(swapInfo.RouterResult.ToTokenAmount, 10)
		if !ok {
			return errors.New("invalid remaining amount"), true
		}
		remainingAmountInt.Add(remainingAmountInt, tmp)
		mkpUserBuy.SellAmountToken[key] = swapInfo.RouterResult.ToTokenAmount

		filter = bson.M{"_id": trackingIdInt}
		updater = bson.M{"$set": bson.M{"status": "swap_project", "sell_swap_token_tx_hash": mkpUserBuy.SellSwapTokenTxHash, "sell_amount_token": mkpUserBuy.SellAmountToken, "swap_project_token_time": time.Since(start).Seconds(), "updatedAt": time.Now(), "swap_project_token_process_by_wallet": processWallet}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, true
		}
	} else {
		tmp, ok = new(big.Int).SetString(mkpUserBuy.SellAmountToken[key], 10)
		if !ok {
			return errors.New("invalid remaining amount"), true
		}
		remainingAmountInt.Add(remainingAmountInt, tmp)
	}

	key = s.server.ExtendConfig.RIZAddress
	if _, ok := mkpUserBuy.SellSwapTokenTxHash[key]; !ok {
		start := time.Now()

		var swapWalletBase *marketplaceDto.MkpSwapWallet
		if mkpRagent.ChainId == int(define.SOLANA) {
			swapWalletBase, err = s.helperGetSwapWalletByAddress(ctx, walletBase, int(define.BASE))
			if err != nil {
				return err, true
			}
			s.helperReleaseSwapWallet(ctx, swapWalletBase.Wallet, int(define.BASE))
		} else {
			swapWalletBase = walletSwapUsed
		}
		processWallet, _, txHash, swapInfo, err := s.helperSwapToken(ctx, "sell", s.server.ExtendConfig.RIZAddress, int(define.BASE), s.server.ExtendConfig.USDCAddress, int(define.BASE), slippageStr, amountRizInt.String(), 8, 6, big.NewFloat(0), false, swapWalletBase)
		if err != nil {
			return err, true
		}

		mkpUserBuy.SellSwapTokenTxHash[key] = txHash
		tmp, ok = new(big.Int).SetString(swapInfo.RouterResult.ToTokenAmount, 10)
		if !ok {
			return errors.New("invalid remaining amount"), true
		}
		remainingAmountInt.Add(remainingAmountInt, tmp)
		mkpUserBuy.SellAmountToken[key] = swapInfo.RouterResult.ToTokenAmount
		filter = bson.M{"_id": trackingIdInt}
		updater = bson.M{"$set": bson.M{"status": "swap_riz", "sell_swap_token_tx_hash": mkpUserBuy.SellSwapTokenTxHash, "sell_amount_token": mkpUserBuy.SellAmountToken, "swap_riz_token_time": time.Since(start).Seconds(), "updatedAt": time.Now(), "swap_riz_token_process_by_wallet": processWallet}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, true
		}

	} else {
		tmp, ok = new(big.Int).SetString(mkpUserBuy.SellAmountToken[key], 10)
		if !ok {
			return errors.New("invalid remaining amount"), true
		}
		remainingAmountInt.Add(remainingAmountInt, tmp)
	}

	// We send the USDC to the owner reduce 2% fee
	if mkpUserBuy.SellTransferToUserTxHash == "" {
		start := time.Now()
		remainingAmountInt.Mul(remainingAmountInt, big.NewInt(int64((1-mkpSetting.SellFee)*100)))
		remainingAmountInt.Div(remainingAmountInt, big.NewInt(100))

		var swapWalletBase *marketplaceDto.MkpSwapWallet
		if mkpRagent.ChainId == int(define.SOLANA) {
			swapWalletBase, err = s.helperGetSwapWalletByAddress(ctx, walletBase, int(define.BASE))
			if err != nil {
				return err, true
			}
			s.helperReleaseSwapWallet(ctx, swapWalletBase.Wallet, int(define.BASE))
		} else {
			swapWalletBase = walletSwapUsed
		}

		txHash, err := nodeClient.TransferERC20(ctx, swapWalletBase.Secret, s.server.ExtendConfig.USDCAddress, wallet_address, remainingAmountInt)
		if err != nil {
			return err, true
		}
		mkpUserBuy.SellTransferToUserTxHash = txHash
		updater = bson.M{"$set": bson.M{"status": "transfer_usdc", "sell_transfer_to_user_tx_hash": mkpUserBuy.SellTransferToUserTxHash, "transfer_to_user_time": time.Since(start).Seconds(), "updatedAt": time.Now()}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, true
		}
	}

skip_swap_token:
	filter = bson.M{"_id": trackingIdInt}
	updater = bson.M{"$set": bson.M{"status": "success", "updatedAt": time.Now()}}
	_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
	if err != nil {
		return err, true
	}
	/////////////resward session 2, sell agent, many time
	tx_hash := mkpUserBuy.TxHash
	if tx_hash == "" {
		//tx hash = uuid
		tx_hash = uuid.New().String()
	}
	filter_user := bson.M{"walletAddress": wallet_address}
	user := userDto.User{}
	err = s.colUser.FindOne(ctx, filter_user).Decode(&user)
	if err == nil {
		rewardSVC, err := rewardService.NewRewardService(s.server)
		logData := map[string]interface{}{
			"wallet_address": wallet_address,
			"ragent_id":      ragent_id,
			"my_agent_id":    ragent_id,
		}
		if err != nil {
			log.Error("Error create reward service: "+err.Error(), define.SESSION2_MKP_SELL_AGENT, logData)
		} else {
			err = rewardSVC.AddUserReward(ctx, user.ID.Hex(), define.SESSION2_MKP_SELL_AGENT, tx_hash, quantity)
			if err != nil {
				log.Error("Error add user reward: "+err.Error(), define.SESSION2_MKP_SELL_AGENT, logData)
			}
		}
	}
	//remove user hold history
	return nil, false
}
