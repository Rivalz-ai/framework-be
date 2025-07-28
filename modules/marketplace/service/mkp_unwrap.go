package service

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/framework/log"
	marketplaceDto "github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	httpClient "github.com/Rivalz-ai/framework-be/modules/marketplace/service/http"
	nodeService "github.com/Rivalz-ai/framework-be/modules/node/service"
	"github.com/coinbase/cdp-sdk/go/openapi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gagliardetto/solana-go"
	associatedtokenaccount "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	abiRagent "github.com/Rivalz-ai/framework-be/define/abi"
)

func (s *MarketplaceService) SubmitSellUnwrapRagent(ctx context.Context, wallet_address, ragent_id, hash string, isSell bool) (int64, string, error) {
	// check if hash is already in db
	filter := bson.M{"tx_hash": hash}
	var mkpUserBuy *marketplaceDto.MkpRagentUserBuy
	err := s.colMkpMyAgent.FindOne(ctx, filter).Decode(&mkpUserBuy)
	if err == nil {
		return 0, "", errors.New("Transaction already verified")
	}

	// get ragent
	ragentObjectId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return 0, "", err
	}
	filter = bson.M{"_id": ragentObjectId}
	var mkpRagent *marketplaceDto.MkpRagent
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return 0, "", err
	}

	side := "unwrap"
	if isSell {
		side = "sell"
	}

	// get token id from transaction
	nodeClient, err := nodeService.NewNodeService(s.server)
	if err != nil {
		return 0, "", err
	}
	//get transaction
	_, txReceipt, err := nodeClient.GetTransactionReceiptByHash(ctx, hash)
	if err != nil {
		//retry
		return 0, "", err
	}
	if txReceipt == nil || txReceipt.Status == types.ReceiptStatusFailed {
		log.Error("Transaction not found,retry: ", "MarketPlaceVerify-GetTransactionReceiptByHash", hash)
		return 0, "", errors.New("NO_RECEIPT")
	}
	if len(txReceipt.Logs) == 0 {
		log.Error("Transaction logs not found, retry: ", "MarketPlaceVerify-GetTransactionReceiptByHash", hash)
		return 0, "", errors.New("NO_LOGS")
	}

	var tokenSellLogs []*types.Log
	eventName := "0xa51cd2e68b51ff3d84bb3a07da68d53f29690e7e01f6c5edcf291fc8165b7e4b"
	for _, ll := range txReceipt.Logs {
		if ll.Topics[0].Hex() == eventName {
			tokenSellLogs = append(tokenSellLogs, ll)
		}
	}

	if len(tokenSellLogs) == 0 {
		return 0, "", errors.New("no unwrap event found")
	}

	tokenIds := make([]int64, 0)
	for _, log := range tokenSellLogs {
		logData := log.Data
		logDataStr := hex.EncodeToString(logData)
		if len(logDataStr) < 128 {
			return 0, "", errors.New("invalid log data")
		}
		for i := 128; i < len(logDataStr); i += 64 {
			tokenId := logDataStr[i : i+64]
			tokenIdInt, ok := new(big.Int).SetString(tokenId, 16)
			if !ok {
				return 0, "", errors.New("invalid tokenId")
			}
			tokenIds = append(tokenIds, tokenIdInt.Int64())
		}
	}

	quantity := int64(len(tokenIds))
	// get mkp_user
	filter = bson.M{"wallet": wallet_address, "ragent_id": ragent_id}
	var mkpUser *marketplaceDto.MkpUser
	err = s.colMkpUser.FindOne(ctx, filter).Decode(&mkpUser)
	if err != nil {
		return 0, "", err
	}

	ragentWalletKeeperAddress, err := s.helperGetAddressFromPrivateKey(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey)
	if err != nil {
		return 0, "", err
	}

	// check if owner is our wallet for all tokenIds
	for _, tokenId := range tokenIds {
		owner, err := nodeClient.OwnerOfERC721(ctx, mkpRagent.RagentTokenAddress, big.NewInt(tokenId))
		if err != nil {
			return 0, "", err
		}
		if owner != ragentWalletKeeperAddress {
			return 0, "", errors.New("not all tokens are owned by our wallet")
		}
	}

	quantityInfo, mapAccountBalance, err := s.GetQuantityTokenByTokenIds(ctx, ragent_id, quantity, side, wallet_address, tokenIds)
	if err != nil {
		return 0, "", err
	}

	var (
		mapTokenId          = make(map[string]struct{})
		remaingTokenIds     = make(map[string]string)
		remaingTokenIdsBase = make(map[string]string)
		removedTokenMap     = make(map[string]string)
		removedTokenMapBase = make(map[string]string)
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

	for tokenId, account := range mkpUser.AccountListBase {
		if _, ok := mapTokenId[tokenId]; !ok {
			remaingTokenIdsBase[tokenId] = account
		} else {
			removedTokenMapBase[tokenId] = account
		}
	}

	if len(removedTokenMap) != len(tokenIds) {
		for tokenId, _ := range mapTokenId {
			if _, ok := removedTokenMap[tokenId]; !ok {
				tokenIdInt, err := strconv.ParseInt(tokenId, 10, 64)
				if err != nil {
					return 0, "", err
				}
				account, err := s.helperGetAccountFromTokenIds(ctx, mkpRagent.TokenAddress, mkpRagent.RagentTokenAddress, tokenIdInt)
				if err != nil {
					return 0, "", err
				}
				removedTokenMap[account] = tokenId
				//remove user hold agent
				go func() {
					err = s.RemoveUserHoldAgent(ctx, wallet_address, ragent_id, tokenId)
					if err != nil {
						log.Error("Error remove user hold agent: "+err.Error(), "MarketPlaceVerify-RemoveUserHoldAgent", tokenId)
					}
				}()
			}
		}
	}

	mkpSetting, err := s.helperGetMkpSetting(ctx)
	if err != nil {
		return 0, "", err
	}
	fee := quantityInfo.TotalPrice * mkpSetting.SellFee

	myAgentId := primitive.NewObjectID()
	record := &marketplaceDto.MkpRagentUserBuy{
		ID:                myAgentId,
		MkpRagentId:       ragent_id,
		WalletAddress:     wallet_address,
		Amount:            strconv.FormatInt(quantity, 10),
		Status:            "verify",
		Side:              side,
		TotalPrice:        strconv.FormatFloat(quantityInfo.TotalPrice, 'f', -1, 64),
		RemovedTokens:     removedTokenMap,
		RemovedTokensBase: removedTokenMapBase,
		AccountList:       remaingTokenIds,
		AccountListBase:   remaingTokenIdsBase,
		TxHash:            hash,
		BalanceToken:      mapAccountBalance,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		TotalFee:          fee,
	}
	_, err = s.colMkpMyAgent.InsertOne(ctx, record)
	if err != nil {
		return 0, "", err
	}

	//update back to mkp_user to make sure they don't sell again
	filter = bson.M{"wallet": wallet_address, "ragent_id": ragent_id}
	_, err = s.colMkpUser.UpdateOne(ctx, filter, bson.M{
		"$set": bson.M{
			"account_list":      remaingTokenIds,
			"account_list_base": remaingTokenIdsBase,
			"updated_at":        time.Now(),
		},
		"$inc": bson.M{
			"quantity": -quantity,
		},
	})
	if err != nil {
		return 0, "", err
	}
	/*
		/////////////resward session 2, sell agent, many time
		filter_user := bson.M{"walletAddress": wallet_address}
		user := userDto.User{}
		err = s.colUser.FindOne(ctx, filter_user).Decode(&user)
		if err == nil {
			rewardSVC, err := rewardService.NewRewardService(s.server)
			logData := map[string]interface{}{
				"wallet_address": wallet_address,
				"ragent_id":      ragent_id,
				"my_agent_id":    myAgentId,
			}
			if err != nil {
				log.Error("Error create reward service: "+err.Error(), define.SESSION2_MKP_SELL_AGENT, logData)
			} else {
				err = rewardSVC.AddUserReward(ctx, user.ID.Hex(), define.SESSION2_MKP_SELL_AGENT, hash, quantity)
				if err != nil {
					log.Error("Error add user reward: "+err.Error(), define.SESSION2_MKP_SELL_AGENT, logData)
				}
			}
		}
	*/
	return quantity, myAgentId.Hex(), nil
}

func (s *MarketplaceService) InternalUnwrapRagent(ctx context.Context, ragent_id, quantity, wallet_address, trackingId string, isSell bool) (error, bool, *big.Int, *big.Int, *marketplaceDto.MkpSwapWallet, string) {
	var (
		filter, updater bson.M
		txHash          string

		mkpUserBuy      *marketplaceDto.MkpRagentUserBuy
		mkpUser         *marketplaceDto.MkpUser
		mkpRagent       *marketplaceDto.MkpRagent
		accountList     = make(map[string]string)
		accountListBase = make(map[string]string)

		wg       sync.WaitGroup
		errCatch error
	)

	trackingIdInt, err := primitive.ObjectIDFromHex(trackingId)
	if err != nil {
		return err, true, nil, nil, nil, ""
	}

	// get ragent info
	agentId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return err, true, nil, nil, nil, ""
	}

	wg.Add(3)

	go func() {
		defer wg.Done()

		// get mkpuser buy
		filter := bson.M{"_id": trackingIdInt}
		err = s.colMkpMyAgent.FindOne(ctx, filter).Decode(&mkpUserBuy)
		if err != nil {
			errCatch = err
			return
		}
		if mkpUserBuy.TransferTokenTxHash == nil {
			mkpUserBuy.TransferTokenTxHash = make(map[string]string)
		}

		if mkpUserBuy.WalletAddress != wallet_address {
			errCatch = errors.New("wallet address not match")
			return
		}

		// get account list from mkp user
		accountList = mkpUserBuy.RemovedTokens
		accountListBase = mkpUserBuy.RemovedTokensBase
	}()

	go func() {
		defer wg.Done()

		// get mkp user
		filter := bson.M{"wallet": wallet_address, "ragent_id": ragent_id}
		err = s.colMkpUser.FindOne(ctx, filter).Decode(&mkpUser)
		if err != nil {
			errCatch = err
			return
		}

	}()

	go func() {
		defer wg.Done()

		filter := bson.M{"_id": agentId}
		err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
		if err != nil {
			errCatch = err
			return
		}
	}()

	wg.Wait()
	if errCatch != nil {
		return errCatch, true, nil, nil, nil, ""
	}

	nodeClientMkp, err := nodeService.NewNodeService(s.server, mkpRagent.ChainId)
	if err != nil {
		return err, true, nil, nil, nil, ""
	}

	nodeClient, err := nodeService.NewNodeService(s.server)
	if err != nil {
		return err, true, nil, nil, nil, ""
	}

	var walletSwapUsed *marketplaceDto.MkpSwapWallet
	if mkpUserBuy.ProcessByWallet == "" {
		walletSwapUsed, err = s.helperGetSwapWallet(ctx, mkpRagent.ChainId)
	} else {
		walletSwapUsed, err = s.helperGetSwapWalletByAddress(ctx, mkpUserBuy.ProcessByWallet, mkpRagent.ChainId)
	}
	if err != nil {
		return err, true, nil, nil, nil, ""
	}

	// update process by wallet
	filter = bson.M{"_id": trackingIdInt}
	updater = bson.M{"$set": bson.M{"process_by_wallet": walletSwapUsed.Wallet, "updatedAt": time.Now()}}
	_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
	if err != nil {
		return err, true, nil, nil, walletSwapUsed, ""
	}

	// if unwrap then need to release swap wallet when this function is done
	// if sell, it will be released in InternalSellRagent
	if !isSell {
		defer s.helperReleaseSwapWallet(ctx, walletSwapUsed.Wallet, mkpRagent.ChainId)
	}

	walletAddress := walletSwapUsed.Wallet
	if mkpRagent.ChainId == int(define.SOLANA) {
		if mkpUserBuy.WalletBase != "" {
			walletAddress = mkpUserBuy.WalletBase
		} else {
			// get swap wallet address
			filter = bson.M{"chain_id": int(define.BASE)}
			var swapWalletBase *marketplaceDto.MkpSwapWallet
			err = s.colMkpSwapWallet.FindOne(ctx, filter).Decode(&swapWalletBase)
			if err != nil {
				return err, true, nil, nil, walletSwapUsed, ""
			}
			walletAddress = swapWalletBase.Wallet
		}
	}

	// convert to big int
	// projectTokenBigFloat := new(big.Float).SetFloat64(mkpRagent.ProjectToken)
	// decimalMultiplier := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(mkpRagent.TokenDecimals)), nil))
	// result := new(big.Float).Mul(projectTokenBigFloat, decimalMultiplier)
	// amountTokenInt, _ := result.Int(nil)
	// amountRizInt := new(big.Int).SetInt64(int64(mkpRagent.RizToken * 1e8))
	amountTokenInt := big.NewInt(0)
	amountRizInt := big.NewInt(0)

	parsedABI, err := abi.JSON(strings.NewReader(abiRagent.ERC20ABI))
	if err != nil {
		return err, true, nil, nil, walletSwapUsed, ""
	}

	start := time.Now()

	var (
		contractAddressList = make([]string, 0, len(accountList))
		toAddressList       = make([]string, 0, len(accountList))
		datumList           = make([][]byte, 0, len(accountList))
	)

	// // skip swap token for dev and local
	// if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "local" {
	// 	goto skip_swap_token
	// }

	if mkpUserBuy.TranferTokenStartTime == nil {
		filter = bson.M{"_id": trackingIdInt}
		updater = bson.M{"$set": bson.M{"transfer_token_start_time": start, "updatedAt": time.Now()}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, true, nil, nil, walletSwapUsed, ""
		}
	}

	for account, tokenId := range accountList {
		key := account + "_" + mkpRagent.TokenAddress
		amountTokenStr := mkpUserBuy.BalanceToken[account+"_"+tokenId].ProjectToken
		amountTokenToSwap, _ := new(big.Int).SetString(amountTokenStr, 10)

		// add to total amount to swap
		amountTokenInt.Add(amountTokenInt, amountTokenToSwap)

		data, err := parsedABI.Pack("transfer", common.HexToAddress(walletAddress), amountTokenToSwap)
		if err != nil {
			return err, true, nil, nil, walletSwapUsed, ""
		}
		// fmt.Println(hex.EncodeToString(data))

		if mkpRagent.ChainId != int(define.BASE) {

			if _, ok := mkpUserBuy.TransferTokenTxHash[key]; !ok {

				// transfer to wallet by coinbase wallet
				coinbaseClient, err := httpClient.NewCoinbaseClient(s.server.ExtendConfig.CoinbaseAPIKey, s.server.ExtendConfig.CoinbaseAPIKeySecret, s.server.ExtendConfig.CoinbaseWalletSecret)
				if err != nil {
					return err, true, nil, nil, walletSwapUsed, ""
				}

				if mkpRagent.ChainId == int(define.SOLANA) {
					nodeClientSolana, err := nodeService.NewNodeService(s.server, int(define.SOLANA))
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}

					_, err = s.helperTransferGasFee(ctx, walletSwapUsed.Secret, accountListBase[tokenId], mkpRagent.ChainId, nil)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}

					fromPub, err := solana.PublicKeyFromBase58(accountListBase[tokenId])
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}
					toPub, err := solana.PublicKeyFromBase58(walletSwapUsed.Wallet)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}

					mint := solana.MustPublicKeyFromBase58(mkpRagent.TokenAddress)
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
						return err, true, nil, nil, walletSwapUsed, ""
					}

					ix := token.NewTransferCheckedInstruction(
						uint64(amountTokenToSwap.Int64()),
						uint8(mkpRagent.TokenDecimals),
						senderTokenAccount,
						mint,
						recipientTokenAccount,
						fromPub,
						nil,
					).Build()
					instructions = append(instructions, ix)

					recent, err := nodeClientSolana.GetLatestBlockhashSolana(ctx)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}

					tx, err := solana.NewTransaction(
						instructions,
						recent.Value.Blockhash,
						solana.TransactionPayer(fromPub),
					)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}

					// Serialize unsigned tx to base64
					txBytes, err := tx.MarshalBinary()
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}
					txBase64 := base64.StdEncoding.EncodeToString(txBytes)

					// Sign transaction with CDP
					signResp, err := coinbaseClient.CdpClient.SignSolanaTransactionWithResponse(ctx, accountListBase[tokenId], nil, openapi.SignSolanaTransactionJSONRequestBody{
						Transaction: txBase64,
					})
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}
					if signResp.JSON200 == nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}

					// Decode signed tx and send to Solana
					signedTxBytes, err := base64.StdEncoding.DecodeString(signResp.JSON200.SignedTransaction)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}

					txHash, err := nodeClientSolana.SendSolanaTransaction(ctx, signedTxBytes)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}

					mkpUserBuy.TransferTokenTxHash[key] = txHash

					filter = bson.M{"_id": trackingIdInt}
					updater = bson.M{"$set": bson.M{"transfer_token_tx_hash": mkpUserBuy.TransferTokenTxHash, "updatedAt": time.Now()}}
					_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}
				} else {
					excluded := make(map[string]bool)
					node, err := nodeClientMkp.GetNode(excluded)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}

					signedTransaction, gasLimit, gasPrice, err := coinbaseClient.CreateAndSignEVMTransaction(ctx, s.server.Redis.Client, node.URL, accountListBase[tokenId], mkpRagent.TokenAddress, int(mkpRagent.ChainId), data)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}
					amountETH := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))

					_, err = s.helperTransferGasFee(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey, accountListBase[tokenId], mkpRagent.ChainId, amountETH)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}

					txHash, err := nodeClientMkp.BroadcastSignedTransaction(ctx, signedTransaction)
					if err != nil {
						if strings.Contains(err.Error(), "insufficient funds for gas * price + value") {
							re := regexp.MustCompile(`tx cost (\d+)`)
							matches := re.FindStringSubmatch(err.Error())
							if len(matches) > 1 {
								txCostStr := matches[1]
								txCost, parseErr := new(big.Int).SetString(txCostStr, 10)
								if parseErr {
									// Transfer additional gas fee
									_, transferErr := s.helperTransferGasFee(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey, accountListBase[tokenId], mkpRagent.ChainId, txCost)
									if transferErr != nil {
										return transferErr, true, nil, nil, walletSwapUsed, ""
									}

									// Retry broadcasting the transaction
									txHash, err = nodeClientMkp.BroadcastSignedTransaction(ctx, signedTransaction)
									if err != nil {
										return err, true, nil, nil, walletSwapUsed, ""
									}
								}
							}
						} else {
							return err, true, nil, nil, walletSwapUsed, ""
						}
					}

					mkpUserBuy.TransferTokenTxHash[key] = txHash

					filter = bson.M{"_id": trackingIdInt}
					updater = bson.M{"$set": bson.M{"transfer_token_tx_hash": mkpUserBuy.TransferTokenTxHash, "updatedAt": time.Now()}}
					_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
					if err != nil {
						return err, true, nil, nil, walletSwapUsed, ""
					}
				}
			}

		} else {
			contractAddressList = append(contractAddressList, account)
			toAddressList = append(toAddressList, mkpRagent.TokenAddress)
			datumList = append(datumList, data)
		}

		// key = account + "_" + s.server.ExtendConfig.RIZAddress
		amountRizStr := mkpUserBuy.BalanceToken[account+"_"+tokenId].RizToken
		amountRizToSwap, _ := new(big.Int).SetString(amountRizStr, 10)

		// add to total amount to swap
		amountRizInt.Add(amountRizInt, amountRizToSwap)

		data, err = parsedABI.Pack("transfer", common.HexToAddress(walletAddress), amountRizToSwap)
		if err != nil {
			return err, true, nil, nil, walletSwapUsed, ""
		}

		contractAddressList = append(contractAddressList, account)
		toAddressList = append(toAddressList, s.server.ExtendConfig.RIZAddress)
		datumList = append(datumList, data)
	}

	// transfer to wallet by coinbase wallet
	if mkpUserBuy.BatchUnwrapTxHash == "" {
		txHash, err = nodeClient.BatchUnwrapFromErc6551(ctx, s.server.ExtendConfig.RagentWalletKeeperPrivateKey, s.server.ExtendConfig.BatchTransferContract, contractAddressList, toAddressList, datumList)
		if err != nil {
			return err, true, nil, nil, walletSwapUsed, ""
		}
		filter = bson.M{"_id": trackingIdInt}
		updater = bson.M{"$set": bson.M{"batch_unwrap_tx_hash": txHash, "batch_unwrap_time": time.Now(), "wallet_base": walletAddress, "updatedAt": time.Now()}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, true, nil, nil, walletSwapUsed, ""
		}
	}

	// transfer money back to user
	// reduce amount to transfer buy unwrap fee
	if !isSell {
		mkpSetting, err := s.helperGetMkpSetting(ctx)
		if err != nil {
			return err, true, nil, nil, nil, ""
		}
		amountTokenInt.Mul(amountTokenInt, big.NewInt(int64((1-mkpSetting.UnwrapFee)*100)))
		amountRizInt.Mul(amountRizInt, big.NewInt(int64((1-mkpSetting.UnwrapFee)*100)))
		amountTokenInt.Div(amountTokenInt, big.NewInt(100))
		amountRizInt.Div(amountRizInt, big.NewInt(100))

		// get final number by mul with quantity
		// quantityInt, err := strconv.ParseInt(quantity, 10, 64)
		// if err != nil {
		// 	return err, true, nil, nil, nil
		// }
		// amountTokenInt.Mul(amountTokenInt, big.NewInt(quantityInt))
		// amountRizInt.Mul(amountRizInt, big.NewInt(quantityInt))

		// transfer project token back to user
		if mkpUserBuy.TransferUserProjectTokenTxHash == "" {
			start := time.Now()
			txHash, err := nodeClientMkp.TransferERC20(ctx, walletSwapUsed.Secret, mkpRagent.TokenAddress, wallet_address, amountTokenInt)
			if err != nil {
				return err, true, nil, nil, nil, ""
			}

			filter = bson.M{"_id": trackingIdInt}
			updater = bson.M{"$set": bson.M{"status": "transfer_project", "transfer_user_project_token_tx_hash": txHash, "transfer_user_project_token_time": time.Since(start).Seconds(), "updatedAt": time.Now()}}
			_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
			if err != nil {
				return err, true, nil, nil, nil, ""
			}
		}

		if mkpUserBuy.TransferUserRizTokenTxHash == "" {
			start := time.Now()
			// transfer riz token back to user
			txHash, err := nodeClient.TransferERC20(ctx, walletSwapUsed.Secret, s.server.ExtendConfig.RIZAddress, wallet_address, amountRizInt)
			if err != nil {
				return err, true, nil, nil, nil, ""
			}
			filter = bson.M{"_id": trackingIdInt}
			updater = bson.M{"$set": bson.M{"status": "transfer_riz", "transfer_user_riz_token_tx_hash": txHash, "transfer_user_riz_token_time": time.Since(start).Seconds(), "updatedAt": time.Now()}}
			_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
			if err != nil {
				return err, true, nil, nil, nil, ""
			}
		}
	}

	// skip_swap_token:
	// update status to transferred_token
	// filter = bson.M{"_id": trackingIdInt}
	// updater = bson.M{"$set": bson.M{"status": "transferred_token", "transfer_token_tx_hash": mkpUserBuy.TransferTokenTxHash, "updatedAt": time.Now()}}
	// _, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
	// if err != nil {
	// 	return err, true, nil, nil, walletSwapUsed
	// }

	//Burn the nft
	if !isSell {
		// update status to success
		trackingIdInt, err := primitive.ObjectIDFromHex(trackingId)
		if err != nil {
			return err, true, nil, nil, nil, ""
		}
		filter = bson.M{"_id": trackingIdInt}
		updater = bson.M{"$set": bson.M{"status": "success", "updatedAt": time.Now()}}
		_, err = s.colMkpMyAgent.UpdateOne(ctx, filter, updater)
		if err != nil {
			return err, true, nil, nil, nil, ""
		}
	}

	return nil, false, amountTokenInt, amountRizInt, walletSwapUsed, walletAddress
}

func (s *MarketplaceService) GetUnwrapTokenQuantity(ctx context.Context, ragent_id string, quantity int64, wallet_address string) (marketplaceDto.UnwrapTokenQuantityResponse, error) {

	quantityInfo, mapAccountBalance, err := s.GetQuantityToken(ctx, ragent_id, quantity, "sell", wallet_address)
	if err != nil {
		return marketplaceDto.UnwrapTokenQuantityResponse{}, err
	}

	agentId, err := primitive.ObjectIDFromHex(ragent_id)
	if err != nil {
		return marketplaceDto.UnwrapTokenQuantityResponse{}, err
	}
	filter := bson.M{"_id": agentId, "status": "active"}
	var mkpRagent *marketplaceDto.MkpRagent
	err = s.colMkpRAgent.FindOne(ctx, filter).Decode(&mkpRagent)
	if err != nil {
		return marketplaceDto.UnwrapTokenQuantityResponse{}, err
	}

	var amountProjectTokenIntDecimal = big.NewInt(0)
	var amountRizIntDecimal = big.NewInt(0)
	for _, balance := range mapAccountBalance {
		amountProjectInt, _ := new(big.Int).SetString(balance.ProjectToken, 10)
		amountProjectTokenIntDecimal.Add(amountProjectTokenIntDecimal, amountProjectInt)
		amountRizInt, _ := new(big.Int).SetString(balance.RizToken, 10)
		amountRizIntDecimal.Add(amountRizIntDecimal, amountRizInt)
	}
	decimalMultiplier := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(mkpRagent.TokenDecimals)), nil))
	amountProjectTokenIntDecimalFloat := new(big.Float).SetInt(amountProjectTokenIntDecimal)
	amountProjectTokenIntDecimalFloat.Quo(amountProjectTokenIntDecimalFloat, decimalMultiplier)
	amountProjectTokenFloat, _ := amountProjectTokenIntDecimalFloat.Float64()

	amountRizIntDecimalFloat := new(big.Float).SetInt(amountRizIntDecimal)
	amountRizIntDecimalFloat.Quo(amountRizIntDecimalFloat, new(big.Float).SetInt(big.NewInt(1e8)))
	amountRizFloat, _ := amountRizIntDecimalFloat.Float64()

	mkpSetting, err := s.helperGetMkpSetting(ctx)
	if err != nil {
		return marketplaceDto.UnwrapTokenQuantityResponse{}, err
	}
	amountProjectToken := math.Round(amountProjectTokenFloat*(1-mkpSetting.UnwrapFee)*1000) / 1000
	amountRiz := math.Round(amountRizFloat*(1-mkpSetting.UnwrapFee)*1000) / 1000

	return marketplaceDto.UnwrapTokenQuantityResponse{
		AmountProjectToken: amountProjectToken,
		AmountRiz:          amountRiz,
		TokenIds:           quantityInfo.TokenIds,
	}, nil
}

func (s *MarketplaceService) helperGetAccountFromTokenIds(ctx context.Context, tokenAddress string, ragentTokenAddress string, tokenId int64) (string, error) {
	nodeClient, err := nodeService.NewNodeService(s.server)
	if err != nil {
		return "", err
	}

	excluded := make(map[string]bool)
	node, err := nodeClient.GetNode(excluded)
	if err != nil {
		return "", err
	}

	topic0 := common.HexToHash("0x79f19b3655ee38b1ce526556b7731a20c8f218fbda4a3990b6cc4172fdf88722")
	address := common.HexToAddress("0x000000006551c19487814612e58FE06813775758")
	ragentTokenAddressHash := common.HexToHash(ragentTokenAddress)
	tokenIdHash := common.HexToHash(strconv.FormatInt(tokenId, 16))

	step := big.NewInt(10000)
	startBlock := big.NewInt(30700000)
	endBlock := big.NewInt(0).Add(startBlock, step)

	for {
		startBlock.Add(startBlock, step)
		endBlock.Add(endBlock, step)
		filter := ethereum.FilterQuery{
			FromBlock: startBlock,
			ToBlock:   endBlock,
			Addresses: []common.Address{address},
			Topics:    [][]common.Hash{{topic0}, {}, {ragentTokenAddressHash}, {tokenIdHash}},
		}
		logs, err := node.Client.FilterLogs(ctx, filter)
		if err != nil {
			break
		}

		for _, log := range logs {
			if len(log.Data) < 64 {
				continue
			}
			account := "0x" + hex.EncodeToString(log.Data)[24:64]

			balance, err := nodeClient.GetBalanceERC20(ctx, account, tokenAddress)
			if err != nil {
				return "", err
			}
			if balance.Cmp(big.NewInt(0)) == 0 {
				return "", errors.New("token not found")
			}

			return account, nil
		}
	}

	return "", errors.New("no account found")
}

func (s *MarketplaceService) helperTransferGasFee(ctx context.Context, secret string, toAddress string, chainId int, amount *big.Int) (string, error) {
	if chainId == int(define.SOLANA) {
		nodeClientSolana, err := nodeService.NewNodeService(s.server, int(define.SOLANA))
		if err != nil {
			return "", err
		}

		from, err := solana.PrivateKeyFromBase58(secret)
		if err != nil {
			return "", err
		}
		fromPub := from.PublicKey()

		toPub, err := solana.PublicKeyFromBase58(toAddress)
		if err != nil {
			return "", err
		}
		amount := big.NewInt(4000000) // 0.005 SOL for transfer fee

		// check if recipient has enough balance
		balance, err := nodeClientSolana.GetTokenBalanceSolana(ctx, toPub.String())
		if err != nil {
			return "", err
		}
		if balance >= 0.004 {
			return "", nil
		}

		lamports := uint64(amount.Int64())
		ix := system.NewTransferInstruction(lamports, fromPub, toPub).Build()
		recent, err := nodeClientSolana.GetLatestBlockhashSolana(ctx)
		if err != nil {
			return "", err
		}

		tx, err := solana.NewTransaction(
			[]solana.Instruction{ix},
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

		txBytes, err := tx.MarshalBinary()
		if err != nil {
			return "", err
		}

		txHash, err := nodeClientSolana.SendSolanaTransaction(ctx, txBytes)
		if err != nil {
			return "", err
		}

		for i := 0; i < 50; i++ {
			balance, err = nodeClientSolana.GetTokenBalanceSolana(ctx, toPub.String())
			if err != nil {
				time.Sleep(time.Millisecond * 300)
				continue
			}
			if balance >= 0.004 {
				break
			}
			time.Sleep(time.Millisecond * 300)
		}

		return txHash, nil
	}

	nodeClient, err := nodeService.NewNodeService(s.server, chainId)
	if err != nil {
		return "", err
	}

	// check if amount is enough
	balance, err := nodeClient.GetBalanceETH(ctx, toAddress)
	if err != nil {
		return "", err
	}
	if balance.Cmp(amount) > 0 {
		return "", nil
	}

	txHash, err := nodeClient.TransferETH(ctx, secret, toAddress, amount)
	if err != nil {
		return "", err
	}

	return txHash, nil
}
