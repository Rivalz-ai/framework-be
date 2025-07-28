package service

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/framework/base/event"
	"github.com/Rivalz-ai/framework-be/framework/log"
	marketplaceDto "github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	nodeService "github.com/Rivalz-ai/framework-be/modules/node/service"
	"go.mongodb.org/mongo-driver/bson"
)

func (sv *MarketplaceService) CheckBalanceOfWallet(wallet, consumed_token string, consumed_token_decimals int) error {
	tokenPrice := 1.0

	if consumed_token != sv.server.ExtendConfig.USDCAddress {
		tokenInfo, err := sv.helperTokenPrice(context.Background(), consumed_token, int(define.BASE))
		if err != nil {
			return err
		}

		tokenPrice = tokenInfo["usd"]
	}

	decimalMultiplierInt := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(consumed_token_decimals)), nil)
	tokenPriceFloat := new(big.Float).SetFloat64(tokenPrice)
	tokenPriceFloat = tokenPriceFloat.Mul(tokenPriceFloat, new(big.Float).SetInt(decimalMultiplierInt))
	tokenPriceInt, _ := tokenPriceFloat.Int(nil)
	LIMIT_BALANCE := new(big.Int).Mul(big.NewInt(int64(sv.server.ExtendConfig.LimitBalanceSwapWallet)), tokenPriceInt)
	VALUE_TOPUP := new(big.Int).Mul(big.NewInt(int64(sv.server.ExtendConfig.ValueTopupSwapWallet)), tokenPriceInt)

	nodeClient, err := nodeService.NewNodeService(sv.server)
	if err != nil {
		return err
	}
	balance, err := nodeClient.GetBalanceERC20(context.Background(), wallet, consumed_token)
	if err != nil {
		log.Error(err.Error(), "ConsumeSwapWallet")
		return err
	}

	// if below limit then topup wallet from keeper wallet
	if balance.Cmp(LIMIT_BALANCE) < 0 {
		// topup wallet from keeper wallet
		txHash, err := nodeClient.TransferERC20(context.Background(), sv.server.ExtendConfig.RagentWalletKeeperPrivateKey, consumed_token, wallet, VALUE_TOPUP)
		if err != nil {
			log.Error(err.Error(), "ConsumeSwapWallet")
			return err
		}
		fmt.Println("txHash: ", txHash)

		// alert telegram
		title := "Marketplace Consume Swap Wallet Topup"
		event2 := event.Event{
			EventName: "alert-internal-tele",
			EventData: map[string]interface{}{
				"level": "info",
				"title": title,
				"in": map[string]interface{}{
					"wallet_address": wallet,
					"token_address":  consumed_token,
					"value_topup":    VALUE_TOPUP.String(),
					"topup_tx_hash":  txHash,
				},
			},
		}
		if sv.server.Pub["tele-internal-alert"] != nil {
			errp := sv.server.Pub["tele-internal-alert"].Publish(event2)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceInternalSwapToken-pushToKafka", event2)
			}
		}

		return nil
	}

	err = sv.CheckBalanceOfWalletEth(wallet)
	if err != nil {
		log.Error(err.Error(), "ConsumeSwapWallet")
	}

	return nil
}

func (sv *MarketplaceService) CheckBalanceOfWalletEth(wallet string) error {
	decimalMultiplierInt := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(12)), nil)
	LIMIT_BALANCE := new(big.Int).Mul(big.NewInt(int64(sv.server.ExtendConfig.LimitBalanceSwapWalletEth)), decimalMultiplierInt)
	VALUE_TOPUP := new(big.Int).Mul(big.NewInt(int64(sv.server.ExtendConfig.ValueTopupSwapWalletEth)), decimalMultiplierInt)

	nodeClient, err := nodeService.NewNodeService(sv.server)
	if err != nil {
		return err
	}
	balance, err := nodeClient.GetBalanceETH(context.Background(), wallet)
	if err != nil {
		log.Error(err.Error(), "ConsumeSwapWallet")
		return err
	}

	// if below limit then topup wallet from keeper wallet
	if balance.Cmp(LIMIT_BALANCE) < 0 {
		// topup wallet from keeper wallet
		txHash, err := nodeClient.TransferETH(context.Background(), sv.server.ExtendConfig.RagentWalletKeeperPrivateKey, wallet, VALUE_TOPUP)
		if err != nil {
			log.Error(err.Error(), "ConsumeSwapWallet")
			return err
		}
		fmt.Println("txHash: ", txHash)

		// alert telegram
		title := "Marketplace Consume Swap Wallet Topup ETH"
		event2 := event.Event{
			EventName: "alert-internal-tele",
			EventData: map[string]interface{}{
				"level": "info",
				"title": title,
				"in": map[string]interface{}{
					"wallet_address": wallet,
					"value_topup":    VALUE_TOPUP.String(),
					"topup_tx_hash":  txHash,
					"token_address":  "ETH",
				},
			},
		}
		if sv.server.Pub["tele-internal-alert"] != nil {
			errp := sv.server.Pub["tele-internal-alert"].Publish(event2)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceInternalSwapToken-pushToKafka", event2)
			}
		}

		return nil
	}

	return nil
}

func (sv *MarketplaceService) CronjobCheckBalanceOfWallet() {

	// use redis to lock cronjob so that just run on one pod
	lockKey := "cronjob_check_balance_of_wallet"
	lockValue := "1"
	lockDuration := 5 * time.Minute

	// Try to acquire lock with NX (only set if key doesn't exist)
	acquired, err := sv.server.Redis.Client.SetNX(context.Background(), lockKey, lockValue, lockDuration).Result()
	if err != nil {
		log.Error("Failed to acquire lock: "+err.Error(), "CronjobCheckBalanceOfWallet")
		return
	}

	// If lock not acquired, another pod is running the job
	if !acquired {
		log.Info("Another pod is running the job", "CronjobCheckBalanceOfWallet")
		return
	}

	// get list token
	var mkpRAgentList []*marketplaceDto.MkpRagent

	filter := bson.M{}
	cursor, err := sv.colMkpRAgent.Find(context.Background(), filter)
	if err != nil {
		log.Error(err.Error(), "CronjobCheckBalanceOfWallet")
		return
	}

	for cursor.Next(context.Background()) {
		var mkpRAgent marketplaceDto.MkpRagent
		err := cursor.Decode(&mkpRAgent)
		if err != nil {
			log.Error(err.Error(), "CronjobCheckBalanceOfWallet")
			continue
		}
		mkpRAgentList = append(mkpRAgentList, &mkpRAgent)
	}

	// add USDC to list
	mkpRAgentList = append(mkpRAgentList, &marketplaceDto.MkpRagent{
		TokenAddress:  sv.server.ExtendConfig.USDCAddress,
		TokenDecimals: 6,
	})

	// get all swap wallet
	filter = bson.M{}
	cursor, err = sv.colMkpSwapWallet.Find(context.Background(), filter)
	if err != nil {
		log.Error(err.Error(), "CronjobCheckBalanceOfWallet")
		return
	}

	for cursor.Next(context.Background()) {
		var mkpSwapWallet marketplaceDto.MkpSwapWallet
		err := cursor.Decode(&mkpSwapWallet)
		if err != nil {
			log.Error(err.Error(), "CronjobCheckBalanceOfWallet")
			continue
		}

		for _, mkpRAgent := range mkpRAgentList {
			err = sv.CheckBalanceOfWallet(mkpSwapWallet.Wallet, mkpRAgent.TokenAddress, mkpRAgent.TokenDecimals)
			if err != nil {
				log.Error(err.Error(), "CronjobCheckBalanceOfWallet")
			}
		}
	}
}
