package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"time"

	defineabi "github.com/Rivalz-ai/framework-be/define/abi"
	"github.com/Rivalz-ai/framework-be/framework/log"
	httpClient "github.com/Rivalz-ai/framework-be/modules/marketplace/service/http"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (n *Node) BatchTransferERC20(ctx context.Context, secretKey string, contractAddress string, tokenAddress string, toAddressMap map[string]*big.Int, retryCount uint64) (string, error) {
	erc20Svc, err := defineabi.NewBatchTransfer(common.HexToAddress(contractAddress), n.Client)
	if err != nil {
		log.Error("error when new erc20 service: "+err.Error(), "TransferERC20-Erc20Service")
		return "", err
	}

	start := time.Now()

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
	nonce, err := n.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	// get chain id
	chainID, err := n.Client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	gasPrice, err := n.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(120)) // 20% more to speed up
	gasPrice = gasPrice.Div(gasPrice, big.NewInt(100))

	etherClient := httpClient.NewEtherscanClient(nil)
	gasPriceEther, err := etherClient.GetGasPrice(ctx, int(chainID.Int64()))
	if err == nil {
		gasPriceFl, ok := new(big.Float).SetString(gasPriceEther.FastGasPrice)
		if ok {
			gasPriceFl.Mul(gasPriceFl, big.NewFloat(1e9))
			gasPriceTmp, _ := gasPriceFl.Int(nil)

			if gasPriceTmp.Cmp(gasPrice) > 0 {
				gasPrice = gasPriceTmp
			}
		}
	}

	gasTip, err := n.Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // No ETH value needed for ERC20 transfers
	auth.GasLimit = uint64(1000000) // Standard ERC20 transfer uses ~65,000 gas
	auth.GasTipCap = gasTip
	auth.GasFeeCap = gasPrice

	if len(toAddressMap) > 20 {
		auth.GasLimit = uint64(65000) * uint64(len(toAddressMap))
	}

	toAddressList := make([]common.Address, 0, len(toAddressMap))
	amountList := make([]*big.Int, 0, len(toAddressMap))
	for toAddress, amount := range toAddressMap {
		toAddressList = append(toAddressList, common.HexToAddress(toAddress))
		amountList = append(amountList, amount)
	}

	tx, err := erc20Svc.BatchTransfer(auth, common.HexToAddress(tokenAddress), toAddressList, amountList)
	if err != nil {
		log.Error("error when transfer erc20: "+err.Error(), "TransferERC20-Erc20Service")
		return "", err
	}

	receipt, err := bind.WaitMined(ctx, n.Client, tx)
	if err != nil {
		return "", err
	}

	log.Info("BATCH_TRANSFER_ERC20", "TIME_PROCESS", map[string]interface{}{
		"time_process":    time.Since(start).Seconds(),
		"tx_hash":         tx.Hash().Hex(),
		"token_address":   tokenAddress,
		"to_address_list": toAddressList,
		"amount_list":     amountList,
		"node":            n.URL,
	})

	if receipt.Status == 1 {

	} else {
		return "", errors.New("transaction failed")
	}

	return tx.Hash().Hex(), nil
}

func (n *Node) BatchUnwrapFromErc6551(ctx context.Context, secretKey string, contractAddress string, contractAddressList []string, toAddressMap []string, datumList [][]byte, retryCount uint64) (string, error) {
	erc20Svc, err := defineabi.NewBatchTransfer(common.HexToAddress(contractAddress), n.Client)
	if err != nil {
		log.Error("error when new erc20 service: "+err.Error(), "TransferERC20-Erc20Service")
		return "", err
	}

	start := time.Now()

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
	nonce, err := n.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := n.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	// get chain id
	chainID, err := n.Client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // No ETH value needed for ERC20 transfers
	auth.GasLimit = uint64(1000000) // Standard ERC20 transfer uses ~65,000 gas
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(110)).Div(gasPrice, big.NewInt(100))
	if len(toAddressMap) > 20 {
		auth.GasLimit = uint64(65000) * uint64(len(toAddressMap))
	}

	contractAddresses := make([]common.Address, 0, len(contractAddressList))
	for _, tokenAddress := range contractAddressList {
		contractAddresses = append(contractAddresses, common.HexToAddress(tokenAddress))
	}

	toAddressList := make([]common.Address, 0, len(toAddressMap))
	for _, toAddress := range toAddressMap {
		toAddressList = append(toAddressList, common.HexToAddress(toAddress))
	}

	tx, err := erc20Svc.UnwrapFromErc6551(auth, contractAddresses, toAddressList, datumList)
	if err != nil {
		log.Error("error when transfer erc20: "+err.Error(), "TransferERC20-Erc20Service")
		return "", err
	}

	receipt, err := bind.WaitMined(ctx, n.Client, tx)
	if err != nil {
		return "", err
	}

	log.Info("BATCH_TRANSFER_ERC20", "TIME_PROCESS", map[string]interface{}{
		"time_process":          time.Since(start).Seconds(),
		"tx_hash":               tx.Hash().Hex(),
		"contract_address_list": contractAddressList,
		"to_address_list":       toAddressList,
		"datum_list":            datumList,
		"node":                  n.URL,
	})

	if receipt.Status == 1 {

	} else {
		return "", errors.New("transaction failed")
	}

	return tx.Hash().Hex(), nil
}
