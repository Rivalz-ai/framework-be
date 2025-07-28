package service

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	nodeService "github.com/Rivalz-ai/framework-be/modules/node/service"
	protjectDto "github.com/Rivalz-ai/framework-be/modules/project/dto"
	httpClientCGK "github.com/Rivalz-ai/framework-be/modules/project/service/httpclient"
	rewardService "github.com/Rivalz-ai/framework-be/modules/reward/service"
	userDto "github.com/Rivalz-ai/framework-be/modules/user/dto"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BidCreated(address indexed bidder, BidInfo bid);
//
//	struct BidInfo {
//		uint256 bidId;
//		address bidder;
//		uint256 quantity;
//		uint256 totalPrice;
//		uint256 vestingPeriod;
//		address paymentToken;
//		address rAgentToken;
//		uint256 createdAt;
//		bool isReleased;
//		bool isAgreed;
//		address seller;
//		uint256 agreedAt;
//	}
func (s *MarketplaceService) VerifyCreateBid(ctx context.Context, wallet_address, user_id, ragent_id string, hash string, is_retry bool, is_premium bool) (error, bool) {
	wallet_address = strings.ToLower(wallet_address)
	logData := map[string]interface{}{
		"hash":           hash,
		"wallet_address": wallet_address,
		"ragent_id":      ragent_id,
	}
	log.Info("MarketPlaceVerify: ", "Verify", logData)

	// get mkp ragent
	ragentId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return err, true
	}
	filter := bson.M{"_id": ragentId}
	mkpRagent := dto.MkpRagent{}
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return err, false
	}

	nodeClient, err := nodeService.NewNodeService(s.server)
	if err != nil {
		return err, true
	}

	// check if this txn is already verified
	filter = bson.M{"created_bid_tx_hash": hash}
	mkpBid := dto.Bid{}
	err = s.colMkpBid.FindOne(ctx, filter).Decode(&mkpBid)
	if err == nil {
		return errors.New("Transaction already verified"), false
	}

	//get transaction
	_, txReceipt, err := nodeClient.GetTransactionReceiptByHash(ctx, hash)
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

	var tokenLendingLogs []*types.Log
	eventName := "0xec51c76e6e4f403ac15e18304060b7f42a37bf30e6187cf4b388b57f077dba06" // BidCreated
	for _, ll := range txReceipt.Logs {
		if ll.Topics[0].Hex() == eventName {
			tokenLendingLogs = append(tokenLendingLogs, ll)
		}
	}
	if len(tokenLendingLogs) == 0 {
		return errors.New("no bid created event found"), true
	}
	for _, logEvent := range tokenLendingLogs {
		bidder := logEvent.Topics[1].Hex()
		bidder = bidder[:2] + bidder[26:]
		if bidder != wallet_address {
			continue
		}

		logData := logEvent.Data
		logDataStr := hex.EncodeToString(logData)
		log.Info("MarketPlaceVerify: ", "BidInfo", logDataStr)
		fmt.Println(len(logDataStr))

		bidId := logDataStr[:64]
		quantity := logDataStr[128:192]
		totalPrice := logDataStr[192:256]
		vestingPeriod := logDataStr[256:320]
		paymentToken := logDataStr[320:384]
		rAgentToken := logDataStr[384:448]
		createdAt := logDataStr[448:512]
		// isReleased := logData[512:576]
		// isAgreed := logData[576:640]
		// seller := logData[640:]

		bidOnchainId, ok := new(big.Int).SetString(bidId, 16)
		if !ok {
			return errors.New("invalid bidId"), true
		}
		quantityInt, ok := new(big.Int).SetString(quantity, 16)
		if !ok {
			return errors.New("invalid quantity"), true
		}
		totalPriceInt, ok := new(big.Int).SetString(totalPrice, 16)
		if !ok {
			return errors.New("invalid totalPrice"), true
		}
		vestingPeriodInt, ok := new(big.Int).SetString(vestingPeriod, 16)
		if !ok {
			return errors.New("invalid vestingPeriod"), true
		}
		createdAtInt, ok := new(big.Int).SetString(createdAt, 16)
		if !ok {
			return errors.New("invalid createdAt"), true
		}
		createdAtTime := time.Unix(createdAtInt.Int64(), 0)
		paymentToken = "0x" + paymentToken[24:]
		rAgentToken = "0x" + rAgentToken[24:]
		if rAgentToken != mkpRagent.RagentTokenAddress {
			return errors.New("invalid ragent token"), false
		}

		// log.Info("MarketPlaceVerify: ", "BidInfo", bidId, quantity, totalPrice, vestingPeriod, paymentToken, rAgentToken, createdAt, isReleased, isAgreed, seller)

		// get payment token decimals
		paymentTokenInfo, err := s.helperGetTokenInfo(ctx, paymentToken)
		if err != nil || paymentTokenInfo == nil || paymentTokenInfo.Decimals == 0 {
			paymentTokenInfo = &protjectDto.TokenInfo{
				Decimals: 18,
			}

			if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "local" {
				paymentTokenInfo.Decimals = 6
			}
		}

		// create bid
		bid := dto.Bid{
			ID:                   primitive.NewObjectID(),
			RagentId:             ragent_id,
			BidderWallet:         wallet_address,
			BidAmount:            quantityInt.Int64(),
			BidPrice:             float64(totalPriceInt.Int64()) / float64(quantityInt.Int64()),
			IsPremium:            is_premium,
			VestingType:          vestingPeriodInt.Int64() / 86400,
			Status:               "active",
			CreatedAt:            createdAtTime,
			UpdatedAt:            createdAtTime,
			CreatedBidTxHash:     hash,
			BidOnchainId:         bidOnchainId.String(),
			PaymentToken:         paymentToken,
			PaymentTokenDecimals: paymentTokenInfo.Decimals,
		}

		_, err = s.colMkpBid.InsertOne(ctx, bid)
		if err != nil {
			return err, true
		}
		/////////////resward session 2, bidder, many time
		filter_user := bson.M{"walletAddress": wallet_address}
		user := userDto.User{}
		err = s.colUser.FindOne(ctx, filter_user).Decode(&user)
		if err != nil {
			log.Error("Error find user: "+err.Error(), "MarketPlaceVerify-FindUser", logData)
		} else {
			rewardSVC, err := rewardService.NewRewardService(s.server)
			if err != nil {
				log.Error("Error create reward service: "+err.Error(), define.SESSION2_MKP_BIDDER, logData)
			} else {
				err = rewardSVC.AddUserReward(ctx, user.ID.Hex(), define.SESSION2_MKP_BIDDER, hash)
				if err != nil {
					log.Error("Error add user reward: "+err.Error(), define.SESSION2_MKP_BIDDER, logData)
				}
			}
		}

	}

	return nil, false
}

// BidAgreeded(address indexed bidder, uint256 indexed bidId, address indexed seller);
func (s *MarketplaceService) VerifyAgreeBid(ctx context.Context, wallet_address string, ragent_id string, hash string) (error, bool) {
	wallet_address = strings.ToLower(wallet_address)
	logData := map[string]interface{}{
		"hash":           hash,
		"wallet_address": wallet_address,
		"ragent_id":      ragent_id,
	}
	log.Info("MarketPlaceVerifyAgreeBid: ", "Verify", logData)
	nodeClient, err := nodeService.NewNodeService(s.server)
	if err != nil {
		return err, true
	}

	// check if this txn is already verified
	filter := bson.M{"agree_bid_tx_hash": hash}
	mkpBid := dto.UserBid{}
	err = s.colMkpUserBid.FindOne(ctx, filter).Decode(&mkpBid)
	if err == nil {
		return errors.New("Transaction already verified"), false
	}

	//get transaction
	_, txReceipt, err := nodeClient.GetTransactionReceiptByHash(ctx, hash)
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

	var (
		tokenLendingLogs  []*types.Log
		tokenTransferLogs []*types.Log
		eventName         = "0x12ca40d201bbdcab4c301edfeef15880a259c72a2e6eb60a4d595051b7ab155d" // AgreedBid
		transferEvent     = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	)
	for _, ll := range txReceipt.Logs {
		if ll.Topics[0].Hex() == eventName {
			tokenLendingLogs = append(tokenLendingLogs, ll)
		}
		if ll.Topics[0].Hex() == transferEvent {
			tokenTransferLogs = append(tokenTransferLogs, ll)
		}
	}
	if len(tokenLendingLogs) == 0 {
		return errors.New("no bid agreed event found"), true
	}

	tokenIdsMap := map[string]struct{}{}
	for _, logEvent := range tokenTransferLogs {
		tokenId := logEvent.Topics[3].Hex()
		tokenIdInt, ok := new(big.Int).SetString(tokenId[2:], 16)
		if !ok {
			return errors.New("invalid tokenId"), true
		}
		tokenIdsMap[tokenIdInt.String()] = struct{}{}
	}
	bidderTemp := dto.Bid{}
	for _, logEvent := range tokenLendingLogs {
		bidder := logEvent.Topics[1].Hex()
		bidder = bidder[:2] + bidder[26:]

		bidOnChainId := logEvent.Topics[2].Hex()
		bidOnChainIdInt, ok := new(big.Int).SetString(bidOnChainId[2:], 16)
		if !ok {
			return errors.New("invalid bidOnChainId"), true
		}

		seller := logEvent.Topics[3].Hex()
		seller = seller[:2] + seller[26:]
		if seller != wallet_address {
			continue
		}

		// find bid by bidOnChainId
		filter = bson.M{"bid_onchain_id": bidOnChainIdInt.String(), "ragent_id": ragent_id}
		mkpBid := dto.Bid{}
		err = s.colMkpBid.FindOne(ctx, filter).Decode(&mkpBid)
		if err != nil {
			return errors.New("bid not found"), true
		}

		if bidder != mkpBid.BidderWallet {
			continue
		}

		// create bid
		bid := dto.UserBid{
			ID:             primitive.NewObjectID(),
			BidAmount:      mkpBid.BidAmount,
			BidPrice:       mkpBid.BidPrice,
			IsPremium:      mkpBid.IsPremium,
			VestingType:    mkpBid.VestingType,
			VestingTime:    time.Now().Add(time.Duration(mkpBid.VestingType) * time.Hour * 24),
			Status:         "vesting",
			CreatedAt:      mkpBid.CreatedAt,
			UpdatedAt:      mkpBid.UpdatedAt,
			AgreeBidTxHash: hash,
			BidId:          mkpBid.ID.Hex(),
			BidOnchainId:   bidOnChainIdInt.String(),
			WalletAddress:  seller,
			AgreeBidWallet: wallet_address,
		}

		_, err = s.colMkpUserBid.InsertOne(ctx, bid)
		if err != nil {
			return err, true
		}

		// update bid status to done
		filter = bson.M{"bid_onchain_id": bidOnChainIdInt.String(), "ragent_id": ragent_id}
		update := bson.M{"$set": bson.M{"status": "done", "agree_bid_tx_hash": hash, "updated_at": time.Now()}}
		_, err = s.colMkpBid.UpdateOne(ctx, filter, update)
		if err != nil {
			return err, true
		}

		// update mkp user quantity
		// get list token id of
		filter = bson.M{"wallet": wallet_address, "ragent_id": ragent_id}
		mkpUser := dto.MkpUser{}
		err = s.colMkpUser.FindOne(ctx, filter).Decode(&mkpUser)
		if err != nil {
			return err, true
		}

		var (
			remainingTokenIds = map[string]string{}
			removedTokenIds   = map[string]string{}
		)
		for accountId, tokenId := range mkpUser.AccountList {
			if _, ok := tokenIdsMap[tokenId]; ok {
				removedTokenIds[accountId] = tokenId
				//remove user hold agent
				go func() {
					err = s.RemoveUserHoldAgent(ctx, wallet_address, ragent_id, tokenId)
					if err != nil {
						log.Error("Error remove user hold agent: "+err.Error(), "MarketPlaceVerify-RemoveUserHoldAgent", tokenId)
					}
				}()
			} else {
				remainingTokenIds[accountId] = tokenId
			}
		}

		// update mkp user quantity
		update = bson.M{"$set": bson.M{"account_list": remainingTokenIds, "quantity": len(remainingTokenIds), "removed_account_list": removedTokenIds}}
		_, err = s.colMkpUser.UpdateOne(ctx, filter, update)
		if err != nil {
			return err, true
		}

		// update mkp user to bid owner
		filter = bson.M{"wallet": mkpBid.BidderWallet, "ragent_id": ragent_id}
		mkpUser = dto.MkpUser{}
		err = s.colMkpUser.FindOne(ctx, filter).Decode(&mkpUser)
		if err != nil {
			return err, true
		}
		for acc, tokenId := range removedTokenIds {
			mkpUser.AccountList[acc] = tokenId
		}
		update = bson.M{"$set": bson.M{"account_list": mkpUser.AccountList, "quantity": len(mkpUser.AccountList)}}
		_, err = s.colMkpUser.UpdateOne(ctx, filter, update)
		if err != nil {
			return err, true
		}
		if bidderTemp.BidderWallet == "" {
			bidderTemp = mkpBid
		}
	}
	//
	/////////////resward session 2, first month bid was accepted, one-time
	//find user whom accepte bid will receive reward
	filter_user := bson.M{"walletAddress": wallet_address}
	user := userDto.User{}
	err = s.colUser.FindOne(ctx, filter_user).Decode(&user)
	if err != nil {
		log.Error("Error find user: "+err.Error(), "MarketPlaceVerify-FindUser", logData)
	} else {
		err, isAcceptBIDFirstMonth := s.IsBidAcceptAgentFirstMonth(ctx, user.ID.Hex(), ragent_id)
		if err != nil {
			log.Error("Error check is buy agent first month: "+err.Error(), "MarketPlaceVerify-IsBuyAgentFirstMonth", logData)
		} else {
			if isAcceptBIDFirstMonth {
				rewardSVC, err := rewardService.NewRewardService(s.server)
				if err != nil {
					log.Error("Error create reward service: "+err.Error(), define.SESSION2_MKP_BID_ACCEPT_ANY_AGENT_FIRST_MONTH, logData)
				} else {
					err = rewardSVC.AddUserReward(ctx, user.ID.Hex(), define.SESSION2_MKP_BID_ACCEPT_ANY_AGENT_FIRST_MONTH, hash)
					if err != nil {
						log.Error("Error add user reward: "+err.Error(), define.SESSION2_MKP_BID_ACCEPT_ANY_AGENT_FIRST_MONTH, logData)
					}
				}

			}
		}
	}

	/////////////resward session 2, bidder was accepted by someone, then they will receive reward, many time
	filter_user = bson.M{"walletAddress": bidderTemp.BidderWallet}
	user = userDto.User{}
	err = s.colUser.FindOne(ctx, filter_user).Decode(&user)
	if err != nil {
		log.Error("Error find user: "+err.Error(), "MarketPlaceVerify-FindUser", logData)
	} else {
		rewardSVC, err := rewardService.NewRewardService(s.server)
		if err != nil {
			log.Error("Error create reward service: "+err.Error(), define.SESSION2_MKP_BIDDER_WAS_ACCEPT, logData)
		} else {
			err = rewardSVC.AddUserReward(ctx, user.ID.Hex(), define.SESSION2_MKP_BIDDER_WAS_ACCEPT, hash)
			if err != nil {
				log.Error("Error add user reward: "+err.Error(), define.SESSION2_MKP_BIDDER_WAS_ACCEPT, logData)
			}
		}
	}

	return nil, false
}

// BidReleased(address indexed bidder, uint256 indexed bidId, address indexed seller);
func (s *MarketplaceService) VerifyReleaseBid(ctx context.Context, wallet_address string, ragent_id string, hash string) (error, bool) {
	wallet_address = strings.ToLower(wallet_address)
	logData := map[string]interface{}{
		"hash":           hash,
		"wallet_address": wallet_address,
		"ragent_id":      ragent_id,
	}
	log.Info("MarketPlaceVerifyAgreeBid: ", "Verify", logData)
	nodeClient, err := nodeService.NewNodeService(s.server)
	if err != nil {
		return err, true
	}

	// check if this txn is already verified
	filter := bson.M{"release_bid_tx_hash": hash}
	mkpBid := dto.UserBid{}
	err = s.colMkpUserBid.FindOne(ctx, filter).Decode(&mkpBid)
	if err == nil {
		return errors.New("Transaction already verified"), false
	}

	//get transaction
	_, txReceipt, err := nodeClient.GetTransactionReceiptByHash(ctx, hash)
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

	var tokenLendingLogs []*types.Log
	eventName := "0xbb6507c5d47bbab2eb31717fa9b641f289e7091aa3de7298f3d85da411071bbb" // ReleasedBid
	for _, ll := range txReceipt.Logs {
		if ll.Topics[0].Hex() == eventName {
			tokenLendingLogs = append(tokenLendingLogs, ll)
		}
	}
	if len(tokenLendingLogs) == 0 {
		return errors.New("no bid released event found"), true
	}
	bidderTemp := dto.Bid{}
	for _, logEvent := range tokenLendingLogs {
		bidder := logEvent.Topics[1].Hex()
		bidder = bidder[:2] + bidder[26:]

		bidOnChainId := logEvent.Topics[2].Hex()
		bidOnChainIdInt, ok := new(big.Int).SetString(bidOnChainId[2:], 16)
		if !ok {
			return errors.New("invalid bidOnChainId"), true
		}

		// find bid by bidOnChainId
		filter = bson.M{"bid_onchain_id": bidOnChainIdInt.String(), "ragent_id": ragent_id}
		mkpBid := dto.Bid{}
		err = s.colMkpBid.FindOne(ctx, filter).Decode(&mkpBid)
		if err != nil {
			return errors.New("bid not found"), true
		}

		if bidder != mkpBid.BidderWallet {
			continue
		}

		seller := logEvent.Topics[3].Hex()
		seller = seller[:2] + seller[26:]
		if seller != wallet_address {
			continue
		}

		// update bid status to done
		filter = bson.M{"bid_onchain_id": bidOnChainIdInt.String(), "bid_id": mkpBid.ID.Hex()}
		update := bson.M{"$set": bson.M{"status": "done"}}
		_, err = s.colMkpUserBid.UpdateOne(ctx, filter, update)
		if err != nil {
			return err, true
		}
		if bidderTemp.BidderWallet == "" {
			bidderTemp = mkpBid
		}
	}
	/////////////resward session 2, first month bid was accepted, one-time
	//find user whom accepte bid will receive reward
	/*
		filter_user := bson.M{"walletAddress": wallet_address}
		user := userDto.User{}
		err = s.colUser.FindOne(ctx, filter_user).Decode(&user)
		if err != nil {
			log.Error("Error find user: "+err.Error(), "MarketPlaceVerify-FindUser", logData)
		} else {
			err, isAcceptBIDFirstMonth := s.IsBidAcceptAgentFirstMonth(ctx, user.ID.Hex(), ragent_id)
			if err != nil {
				log.Error("Error check is buy agent first month: "+err.Error(), "MarketPlaceVerify-IsBuyAgentFirstMonth", logData)
			} else {
				if isAcceptBIDFirstMonth {
					rewardSVC, err := rewardService.NewRewardService(s.server)
					if err != nil {
						log.Error("Error create reward service: "+err.Error(), define.SESSION2_MKP_BID_ACCEPT_ANY_AGENT_FIRST_MONTH, logData)
					} else {
						err = rewardSVC.AddUserReward(ctx, user.ID.Hex(), define.SESSION2_MKP_BID_ACCEPT_ANY_AGENT_FIRST_MONTH, hash)
						if err != nil {
							log.Error("Error add user reward: "+err.Error(), define.SESSION2_MKP_BID_ACCEPT_ANY_AGENT_FIRST_MONTH, logData)
						}
					}

				}
			}
		}
	*/

	/////////////resward session 2, bidder was accepted by someone, then they will receive reward, many time
	//ex: A create BID, B accept bid, A will receive reward
	filter_user := bson.M{"walletAddress": bidderTemp.BidderWallet}
	user := userDto.User{}
	err = s.colUser.FindOne(ctx, filter_user).Decode(&user)
	if err != nil {
		log.Error("Error find user: "+err.Error(), "MarketPlaceVerify-FindUser", logData)
	} else {
		rewardSVC, err := rewardService.NewRewardService(s.server)
		if err != nil {
			log.Error("Error create reward service: "+err.Error(), define.SESSION2_MKP_BIDDER_WAS_ACCEPT, logData)
		} else {
			err = rewardSVC.AddUserReward(ctx, user.ID.Hex(), define.SESSION2_MKP_BIDDER_WAS_ACCEPT, hash)
			if err != nil {
				log.Error("Error add user reward: "+err.Error(), define.SESSION2_MKP_BIDDER_WAS_ACCEPT, logData)
			}
		}
	}
	//reward session 2, accepter will receive reward, manytime
	//ex: A create BID, B accept BID, B will receive reward
	filter_user = bson.M{"walletAddress": wallet_address}
	user = userDto.User{}
	err = s.colUser.FindOne(ctx, filter_user).Decode(&user)
	if err != nil {
		log.Error("Error find user: "+err.Error(), "MarketPlaceVerify-FindUser", logData)
	} else {
		rewardSVC, err := rewardService.NewRewardService(s.server)
		if err != nil {
			log.Error("Error create reward service: "+err.Error(), define.SESSION2_MKP_ACCEPT_BID, logData)
		} else {
			err = rewardSVC.AddUserReward(ctx, user.ID.Hex(), define.SESSION2_MKP_ACCEPT_BID, hash)
			if err != nil {
				log.Error("Error add user reward: "+err.Error(), define.SESSION2_MKP_ACCEPT_BID, logData)
			}
		}
	}

	return nil, false
}

// BidRemoved(address indexed bidder, uint256 indexed bidId);
func (s *MarketplaceService) VerifyRemoveBid(ctx context.Context, wallet_address string, ragent_id string, hash string) (error, bool) {
	wallet_address = strings.ToLower(wallet_address)
	logData := map[string]interface{}{
		"hash":           hash,
		"wallet_address": wallet_address,
		"ragent_id":      ragent_id,
	}
	log.Info("MarketPlaceVerify: ", "Verify", logData)
	nodeClient, err := nodeService.NewNodeService(s.server)
	if err != nil {
		return err, true
	}

	// check if this txn is already verified
	filter := bson.M{"removed_bid_tx_hash": hash}
	mkpBid := dto.Bid{}
	err = s.colMkpBid.FindOne(ctx, filter).Decode(&mkpBid)
	if err == nil {
		return errors.New("Transaction already verified"), false
	}

	//get transaction
	_, txReceipt, err := nodeClient.GetTransactionReceiptByHash(ctx, hash)
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

	var tokenLendingLogs []*types.Log
	eventName := "0x7fd2c6472e4a5cf47dd45d6e616062bb6b2a25c19206eaa78f5ac8ea80cd5f97" // BidCreated
	for _, ll := range txReceipt.Logs {
		if ll.Topics[0].Hex() == eventName {
			tokenLendingLogs = append(tokenLendingLogs, ll)
		}
	}
	if len(tokenLendingLogs) == 0 {
		return errors.New("no bid created event found"), true
	}

	for _, logEvent := range tokenLendingLogs {
		bidder := logEvent.Topics[1].Hex()
		bidder = bidder[:2] + bidder[26:]

		bidOnChainId := logEvent.Topics[2].Hex()
		bidOnChainIdInt, ok := new(big.Int).SetString(bidOnChainId[2:], 16)
		if !ok {
			return errors.New("invalid bidOnChainId"), true
		}

		// find bid by bidOnChainId
		filter = bson.M{"bid_onchain_id": bidOnChainIdInt.String(), "ragent_id": ragent_id}
		mkpBid := dto.Bid{}
		err = s.colMkpBid.FindOne(ctx, filter).Decode(&mkpBid)
		if err != nil {
			return errors.New("bid not found"), true
		}

		if bidder != mkpBid.BidderWallet {
			continue
		}

		// update status of bid to removed and update removed_bid_tx_hash
		filter = bson.M{"bid_onchain_id": bidOnChainIdInt.String(), "ragent_id": ragent_id}
		update := bson.M{"$set": bson.M{"status": "removed", "removed_bid_tx_hash": hash}}
		_, err = s.colMkpBid.UpdateOne(ctx, filter, update)
		if err != nil {
			return err, true
		}
	}

	return nil, false
}

func (s *MarketplaceService) GetBids(ctx context.Context, ragent_id string, wallet_address string) ([]*dto.BidResponse, error) {
	filter := bson.M{"ragent_id": ragent_id, "status": "active"}
	cur, err := s.colMkpBid.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var bids []*dto.Bid
	if err := cur.All(ctx, &bids); err != nil {
		return nil, err
	}

	var result []*dto.BidResponse
	for _, bid := range bids {
		premiumPercent := 0.0
		if bid.IsPremium {
			premiumPercent = 20
		}

		bidderInfo, err := s.helperGetBidderInfoFromRagent(ctx, bid.RagentId, bid.BidderWallet)
		if err != nil {
			return nil, err
		}

		fdv, err := s.helperGetFDV(ctx, bidderInfo.BidderToken)
		if err != nil {
			return nil, err
		}
		// Get token price
		tokenPrice, err := s.helperTokenPrice(ctx, bid.PaymentToken, utils.ItoInt(bid.BidOnchainId))
		if err != nil {
			tokenPrice = make(map[string]float64)
		}
		bidResponse := &dto.BidResponse{
			ID:             bid.ID.Hex(),
			BidId:          bid.BidOnchainId,
			BidderInfo:     bidderInfo,
			RagentsToSell:  bid.BidAmount,
			EstValue:       bid.BidPrice * float64(bid.BidAmount) / math.Pow10(int(bid.PaymentTokenDecimals)),
			PremiumPercent: premiumPercent,
			FDV:            fdv,
			Vesting:        bid.VestingType,
			RagentId:       bid.RagentId,
			MC:             tokenPrice["usd_market_cap"],
			TokenAmount:    bid.BidAmount,
		}

		// if wallet_address != "" {
		// 	isOwner := false
		// 	if bid.BidderWallet == wallet_address {
		// 		isOwner = true
		// 	}
		// 	bidResponse.IsOwner = isOwner
		// }
		if wallet_address != "" {
			// check if this wallet has enough token id to buy this bid
			filter := bson.M{"wallet": wallet_address, "ragent_id": bid.RagentId}
			user := dto.MkpUser{}
			err := s.colMkpUser.FindOne(ctx, filter).Decode(&user)
			if err == nil {
				if user.Quantity >= bid.BidAmount {
					bidResponse.IsOwner = true
				}
			}
		}

		result = append(result, bidResponse)
	}

	return result, nil
}

func (s *MarketplaceService) GetMyActiveBids(ctx context.Context, wallet_address string) ([]*dto.MyActiveBidResponse, error) {
	filter := bson.M{"bidder_wallet": wallet_address, "status": "active"}
	cur, err := s.colMkpBid.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var bids []*dto.Bid
	if err := cur.All(ctx, &bids); err != nil {
		return nil, err
	}

	var result []*dto.MyActiveBidResponse
	for _, bid := range bids {
		bidderInfo, err := s.helperGetBidderInfoFromRagent(ctx, bid.RagentId, wallet_address)
		if err != nil {
			return nil, err
		}

		fdv, err := s.helperGetFDV(ctx, bidderInfo.BidderToken)
		if err != nil {
			return nil, err
		}

		premiumPercent := 0.0
		if bid.IsPremium {
			premiumPercent = 20
		}

		estValue := bid.BidPrice * float64(bid.BidAmount) / math.Pow10(int(bid.PaymentTokenDecimals))

		// Map to the new response struct
		activeBid := &dto.MyActiveBidResponse{
			RagentId:        bid.RagentId,
			BidId:           bid.BidOnchainId,
			BidderInfo:      bidderInfo,
			RagentsToSell:   bid.BidAmount,
			EstValue:        estValue,
			PremiumPercent:  premiumPercent, // Calculate premium percentage if available
			FDV:             fdv,
			VestingPeriod:   bid.VestingType, // Helper function to determine vesting period in days
			Status:          bid.Status,
			CreatedAt:       bid.CreatedAt,
			FullBidEstValue: estValue * float64(bid.BidAmount),
		}

		// Add to result
		result = append(result, activeBid)
	}

	return result, nil
}

func (s *MarketplaceService) GetMyHistoricalBids(ctx context.Context, wallet_address string) ([]*dto.BidHistoryResponse, error) {
	filter := bson.M{"bidder_wallet": wallet_address, "status": "done"}
	cur, err := s.colMkpBid.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var bids []*dto.Bid
	if err := cur.All(ctx, &bids); err != nil {
		return nil, err
	}

	var result []*dto.BidHistoryResponse
	// compute total value
	for _, bid := range bids {
		// get all user bids
		filter = bson.M{"bid_id": bid.ID.Hex(), "status": "done"}
		cur, err = s.colMkpUserBid.Find(ctx, filter)
		if err != nil {
			return nil, err
		}
		defer cur.Close(ctx)

		var userBids []dto.UserBid
		if err := cur.All(ctx, &userBids); err != nil {
			return nil, err
		}

		totalValue := 0.0
		// compute total value
		for _, userBid := range userBids {
			totalValue += float64(userBid.BidAmount) * userBid.BidPrice
		}

		bidderInfo, err := s.helperGetBidderInfoFromRagent(ctx, bid.RagentId, wallet_address)
		if err != nil {
			return nil, err
		}

		fdv, err := s.helperGetFDV(ctx, bidderInfo.BidderToken)
		if err != nil {
			return nil, err
		}

		premiumPercent := 0.0
		if bid.IsPremium {
			premiumPercent = 20
		}

		// Map to the new response struct
		historyBid := &dto.BidHistoryResponse{
			BidderInfo:     bidderInfo,
			RagentsToSell:  bid.BidAmount,
			TotalValue:     totalValue,
			PremiumPercent: premiumPercent, // Calculate premium percentage if available
			FDV:            fdv,
			VestingPeriod:  bid.VestingType, // Helper function to determine vesting period in days
			Status:         bid.Status,
			ActionDate:     bid.UpdatedAt,
		}

		// Add to result
		result = append(result, historyBid)
	}

	return result, nil
}

func (s *MarketplaceService) GetMyVestingTokens(ctx context.Context, wallet_address string) ([]*dto.VestingTokenResponse, error) {
	filter := bson.M{"wallet_address": wallet_address, "status": "vesting"}
	cur, err := s.colMkpUserBid.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var bids []*dto.UserBid
	if err := cur.All(ctx, &bids); err != nil {
		return nil, err
	}

	var result []*dto.VestingTokenResponse
	for _, bid := range bids {
		bidId, err := primitive.ObjectIDFromHex(bid.BidId)
		if err != nil {
			return nil, err
		}
		filter := bson.M{"_id": bidId}
		mkpBid := dto.Bid{}
		err = s.colMkpBid.FindOne(ctx, filter).Decode(&mkpBid)
		if err != nil {
			return nil, err
		}

		bidderInfo, err := s.helperGetBidderInfoFromRagent(ctx, mkpBid.RagentId, wallet_address)
		if err != nil {
			return nil, err
		}

		// Determine if the token is claimable (based on vesting time)
		isClaimable := time.Now().After(bid.VestingTime)

		fdv, err := s.helperGetFDV(ctx, bidderInfo.BidderToken)
		if err != nil {
			return nil, err
		}

		premiumPercent := 0.0
		if bid.IsPremium {
			premiumPercent = 20
		}

		// Map to the new response struct
		vestingToken := &dto.VestingTokenResponse{
			BidId:          bid.BidOnchainId,
			BidderInfo:     bidderInfo,
			RagentsToSell:  bid.BidAmount,
			EstValue:       bid.BidPrice * float64(bid.BidAmount),
			PremiumPercent: premiumPercent, // Calculate premium percentage if available
			FDV:            fdv,
			VestingStatus:  bid.Status,
			VestingEndsAt:  bid.VestingTime, // Client will handle countdown display
			IsClaimable:    isClaimable,
			RagentId:       mkpBid.RagentId,
		}

		// Add to result
		result = append(result, vestingToken)
	}

	return result, nil
}

func (s *MarketplaceService) helperGetBidderInfoFromRagent(ctx context.Context, ragent_id string, bidder_wallet string) (dto.BidderInfo, error) {
	ragentId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return dto.BidderInfo{}, err
	}
	filter := bson.M{"_id": ragentId}
	mkpRagent := dto.MkpRagent{}
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return dto.BidderInfo{}, err
	}

	rs := dto.BidderInfo{
		BidderName:   bidder_wallet,
		BidderToken:  mkpRagent.TokenAddress,
		BidderSymbol: mkpRagent.TokenSymbol,
		BidderXLink:  mkpRagent.X,
		BidderType:   "user",
	}

	// check if bidder wallet is a swarm owner
	swarmOwner, err := s.SwarmOwner(ctx, bidder_wallet)
	if err != nil {
		return dto.BidderInfo{}, err
	}
	if swarmOwner {
		rs.BidderType = "swarm"
		rs.BidderLogo = mkpRagent.Logo
		rs.BidderName = mkpRagent.Name
	}

	return rs, nil
}

func (s *MarketplaceService) helperGetFDV(ctx context.Context, tokenAddress string) (float64, error) {
	// Check cache first
	cacheKey := "fdv:" + tokenAddress
	rs := s.server.Redis.Client.Get(ctx, cacheKey)
	if rs.Err() == nil {
		fdv, err := rs.Float64()
		if err == nil {
			return fdv, nil
		}
		// If conversion fails, continue to fetch fresh data
	}

	// fdv = total_supply * price
	// get token price and total supply from coingecko
	fdv, err := httpClientCGK.GetCoinGecKoFDV(s.server.ExtendConfig.CoinGeckoURL, s.server.ExtendConfig.CoinGeckoToken, tokenAddress, map[string]interface{}{})
	if err != nil {
		return 0, err
	}

	// Cache the result for 10 minutes as FDV doesn't change frequently
	s.server.Redis.Client.Set(ctx, cacheKey, fdv, time.Minute*10)

	return fdv, nil
}

func (s *MarketplaceService) AcceptBidQuantity(ctx context.Context, bidId, ragentId string) (*dto.AcceptBidQuantityResponse, error) {
	filter := bson.M{"bid_onchain_id": bidId, "ragent_id": ragentId}
	mkpBid := dto.Bid{}
	err := s.colMkpBid.FindOne(ctx, filter).Decode(&mkpBid)
	if err != nil {
		return nil, err
	}

	amountAutoAgreed := mkpBid.BidAmount
	tokenReceive := mkpBid.BidPrice * float64(mkpBid.BidAmount) / math.Pow10(int(mkpBid.PaymentTokenDecimals))

	// get nft contract address
	ragentIdObj, err := primitive.ObjectIDFromHex(mkpBid.RagentId)
	if err != nil {
		return nil, err
	}
	filter = bson.M{"_id": ragentIdObj}
	mkpRagent := dto.MkpRagent{}
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return nil, err
	}

	// get price from coingecko
	var (
		price float64 = 1
	)
	if os.Getenv("ENV") != "dev" {
		priceTmp, err := s.helperTokenPrice(ctx, mkpBid.PaymentToken, mkpRagent.ChainId)
		if err != nil {
			return nil, err
		}
		price = priceTmp["usd"]
	}

	totalValue := tokenReceive * price

	return &dto.AcceptBidQuantityResponse{
		AmountAutoAgreed:   amountAutoAgreed,
		TokenReceive:       tokenReceive,
		TotalValue:         totalValue,
		NFTContractAddress: mkpRagent.RagentTokenAddress,
		// VestingTime:      vestingTime,
	}, nil
}
