package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	defineabi "github.com/Rivalz-ai/framework-be/define/abi"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	httpClient "github.com/Rivalz-ai/framework-be/modules/marketplace/service/http"
	"github.com/Rivalz-ai/framework-be/modules/node/dto"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// GetRewardTokenInfo lấy thông tin token (Decimals + Name)
func (n *Node) GetERC20TokenInfo(ctx context.Context, tokenAddress string) (*dto.Erc20Info, error) {
	contractAddress := common.HexToAddress(tokenAddress)
	var logData = utils.Dictionary()
	logData["token_address"] = tokenAddress
	// Load ABI
	parsedABI, err := abi.JSON(strings.NewReader(defineabi.ERC20_ABI))
	if err != nil {
		log.Error("Failed to load the contract ABI: "+err.Error(), "GetERC20TokenInfo", logData)
		return nil, err
	}

	// Call decimals()
	data, err := parsedABI.Pack("decimals")
	if err != nil {
		log.Error("Failed to pack decimals(): "+err.Error(), "GetERC20TokenInfo", logData)
		return nil, fmt.Errorf("failed to pack decimals(): %v", err)
	}
	result, err := n.Client.CallContract(ctx, ethereum.CallMsg{To: &contractAddress, Data: data}, nil)
	if err != nil {
		log.Error("Failed to call decimals(): "+err.Error(), "GetERC20TokenInfo", logData)
		return nil, fmt.Errorf("failed to call decimals(): %v", err)
	}

	var decimals uint8
	err = parsedABI.UnpackIntoInterface(&decimals, "decimals", result)
	if err != nil {
		log.Error("Failed to unpack decimals: "+err.Error(), "GetERC20TokenInfo", logData)
		return nil, fmt.Errorf("failed to unpack decimals: %v", err)
	}

	// Call name()
	data, err = parsedABI.Pack("name")
	if err != nil {
		return nil, fmt.Errorf("failed to pack name(): %v", err)
	}
	result, err = n.Client.CallContract(ctx, ethereum.CallMsg{To: &contractAddress, Data: data}, nil)
	if err != nil {
		log.Error("Failed to call name(): "+err.Error(), "GetERC20TokenInfo", logData)
		return nil, fmt.Errorf("failed to call name(): %v", err)
	}

	var name string
	err = parsedABI.UnpackIntoInterface(&name, "name", result)
	if err != nil {
		log.Error("Failed to unpack name: "+err.Error(), "GetERC20TokenInfo", logData)
		return nil, fmt.Errorf("failed to unpack name: %v", err)
	}

	data, err = parsedABI.Pack("symbol")
	if err != nil {
		return nil, fmt.Errorf("failed to pack symbol(): %v", err)
	}
	result, err = n.Client.CallContract(ctx, ethereum.CallMsg{To: &contractAddress, Data: data}, nil)
	if err != nil {
		log.Error("Failed to call symbol(): "+err.Error(), "GetERC20TokenInfo", logData)
		return nil, fmt.Errorf("failed to call name(): %v", err)
	}

	var symbol string
	err = parsedABI.UnpackIntoInterface(&symbol, "symbol", result)
	if err != nil {
		log.Error("Failed to unpack symbol: "+err.Error(), "GetERC20TokenInfo", logData)
		return nil, fmt.Errorf("failed to unpack symbol: %v", err)
	}

	return &dto.Erc20Info{
		Address:  tokenAddress,
		Decimals: decimals,
		Name:     name,
		Symbol:   symbol,
	}, nil
}

// EstimateGasERC20Transfer estimates gas limit for ERC20 transfer
func (n *Node) EstimateGasERC20Transfer(ctx context.Context, tokenAddress string, fromAddress string, toAddress string, amount *big.Int) (uint64, error) {
	// Load ABI
	parsedABI, err := abi.JSON(strings.NewReader(defineabi.ERC20_ABI))
	if err != nil {
		return 0, fmt.Errorf("failed to load the contract ABI: %v", err)
	}

	// Pack transfer function call data
	data, err := parsedABI.Pack("transfer", common.HexToAddress(toAddress), amount)
	if err != nil {
		return 0, fmt.Errorf("failed to pack transfer(): %v", err)
	}

	// Estimate gas
	contractAddress := common.HexToAddress(tokenAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(fromAddress),
		To:   &contractAddress,
		Data: data,
	}

	gasLimit, err := n.Client.EstimateGas(ctx, msg)
	if err != nil {
		// Return default gas limit if estimation fails
		log.Error("Failed to estimate gas for ERC20 transfer, using default: "+err.Error(), "EstimateGasERC20Transfer", map[string]interface{}{
			"token_address": tokenAddress,
			"from_address":  fromAddress,
			"to_address":    toAddress,
		})
		return 100000, nil // Default fallback
	}

	// Add 20% buffer to the estimated gas
	gasLimit = gasLimit * 120 / 100
	return gasLimit, nil
}

// EstimateGasERC20Approve estimates gas limit for ERC20 approve
func (n *Node) EstimateGasERC20Approve(ctx context.Context, tokenAddress string, fromAddress string, spenderAddress string, amount *big.Int) (uint64, error) {
	// Load ABI
	parsedABI, err := abi.JSON(strings.NewReader(defineabi.ERC20_ABI))
	if err != nil {
		return 0, fmt.Errorf("failed to load the contract ABI: %v", err)
	}

	// Pack approve function call data
	data, err := parsedABI.Pack("approve", common.HexToAddress(spenderAddress), amount)
	if err != nil {
		return 0, fmt.Errorf("failed to pack approve(): %v", err)
	}

	// Estimate gas
	contractAddress := common.HexToAddress(tokenAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(fromAddress),
		To:   &contractAddress,
		Data: data,
	}

	gasLimit, err := n.Client.EstimateGas(ctx, msg)
	if err != nil {
		// Return default gas limit if estimation fails
		log.Error("Failed to estimate gas for ERC20 approve, using default: "+err.Error(), "EstimateGasERC20Approve", map[string]interface{}{
			"token_address":   tokenAddress,
			"from_address":    fromAddress,
			"spender_address": spenderAddress,
		})
		return 100000, nil // Default fallback
	}

	// Add 20% buffer to the estimated gas
	gasLimit = gasLimit * 120 / 100
	return gasLimit, nil
}

func (n *Node) TransferERC20(ctx context.Context, secretKey string, tokenAddress string, toAddress string, amount *big.Int) (string, error) {
	erc20Svc, err := defineabi.NewERC20(common.HexToAddress(tokenAddress), n.Client)
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

	// Estimate gas limit for the transfer
	gasLimit, err := n.EstimateGasERC20Transfer(ctx, tokenAddress, fromAddress.Hex(), toAddress, amount)
	if err != nil {
		log.Error("Failed to estimate gas limit: "+err.Error(), "TransferERC20", map[string]interface{}{
			"token_address": tokenAddress,
			"from_address":  fromAddress.Hex(),
			"to_address":    toAddress,
		})
		gasLimit = 100000 // Use default fallback
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // No ETH value needed for ERC20 transfers
	auth.GasLimit = gasLimit   // Use estimated gas limit
	// auth.GasPrice = gasPrice
	auth.GasFeeCap = gasPrice
	auth.GasTipCap = gasTip

	tx, err := erc20Svc.Transfer(auth, common.HexToAddress(toAddress), amount)
	if err != nil {
		log.Error("error when transfer erc20: "+err.Error(), "TransferERC20-Erc20Service")
		return "", err
	}

	receipt, err := bind.WaitMined(ctx, n.Client, tx)
	if err != nil {
		return "", err
	}

	log.Info("TRANSFER_ERC20", "TIME_PROCESS", map[string]interface{}{
		"process_time": time.Since(start).Seconds(),
		"tx_hash":      tx.Hash().Hex(),
		"status":       receipt.Status,
		"gas_used":     receipt.GasUsed,
		"gas_limit":    gasLimit,
		"node":         n.URL,
	})

	if receipt.Status == 1 {
	} else {
		return "", errors.New("transaction failed")
	}

	return tx.Hash().Hex(), nil
}

func (n *Node) ApproveERC20(ctx context.Context, secretKey string, spenderAddress string, tokenAddress string, amount *big.Int) (string, error) {
	erc20Svc, err := defineabi.NewERC20(common.HexToAddress(tokenAddress), n.Client)
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

	// check allowance
	allowance, err := erc20Svc.Allowance(nil, fromAddress, common.HexToAddress(spenderAddress))
	if err != nil {
		return "", err
	}

	if allowance.Cmp(amount) >= 0 {
		return "", nil
	}

	nonce, err := n.Client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", err
	}

	// get chain id
	chainID, err := n.Client.ChainID(ctx)
	if err != nil {
		return "", err
	}

	gasPrice, err := n.Client.SuggestGasPrice(ctx)
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

	// Estimate gas limit for the approve
	gasLimit, err := n.EstimateGasERC20Approve(ctx, tokenAddress, fromAddress.Hex(), spenderAddress, amount)
	if err != nil {
		log.Error("Failed to estimate gas limit: "+err.Error(), "ApproveERC20", map[string]interface{}{
			"token_address":   tokenAddress,
			"from_address":    fromAddress.Hex(),
			"spender_address": spenderAddress,
		})
		gasLimit = 100000 // Use default fallback
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // No ETH value needed for ERC20 transfers
	auth.GasLimit = gasLimit   // Use estimated gas limit
	// auth.GasPrice = gasPrice
	auth.GasTipCap = gasTip
	auth.GasFeeCap = gasPrice

	tx, err := erc20Svc.Approve(auth, common.HexToAddress(spenderAddress), amount)
	if err != nil {
		log.Error("error when transfer erc20: "+err.Error(), "TransferERC20-Erc20Service")
		return "", err
	}

	receipt, err := bind.WaitMined(ctx, n.Client, tx)
	if err != nil {
		return "", err
	}

	log.Info("APPROVE_ERC20", "TIME_PROCESS", map[string]interface{}{
		"time_process":    time.Since(start).Seconds(),
		"tx_hash":         tx.Hash().Hex(),
		"spender_address": spenderAddress,
		"token_address":   tokenAddress,
		"amount":          amount,
		"gas_used":        receipt.GasUsed,
		"gas_limit":       gasLimit,
		"node":            n.URL,
	})

	if receipt.Status == 1 {

	} else {
		return "", errors.New("transaction failed")
	}

	return tx.Hash().Hex(), nil

}

func (n *Node) GetBalanceERC20(ctx context.Context, walletAddress string, tokenAddress string) (*big.Int, error) {
	erc20Svc, err := defineabi.NewERC20(common.HexToAddress(tokenAddress), n.Client)
	if err != nil {
		return nil, err
	}

	balance, err := erc20Svc.BalanceOf(nil, common.HexToAddress(walletAddress))
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (n *Node) GetBalanceETH(ctx context.Context, walletAddress string) (*big.Int, error) {
	balance, err := n.Client.BalanceAt(ctx, common.HexToAddress(walletAddress), nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (n *Node) TransferETH(ctx context.Context, secretKey string, toAddress string, amount *big.Int) (string, error) {
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
	nonce, err := n.Client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := n.Client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(120)) // 20% more to speed up
	gasPrice = gasPrice.Div(gasPrice, big.NewInt(100))

	// get chain id
	chainID, err := n.Client.ChainID(ctx)
	if err != nil {
		return "", err
	}

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

	gasTip, err := n.Client.SuggestGasTipCap(ctx)
	if err != nil {
		return "", err
	}

	toAddressCommon := common.HexToAddress(toAddress)
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		To:        &toAddressCommon,
		Value:     amount,
		Gas:       uint64(22000),
		GasFeeCap: gasPrice,
		GasTipCap: gasTip,
	})

	// 6. Create the Transaction
	// tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), amount, uint64(22000), gasPrice, nil) // 'nil' for data in a simple ETH transfer

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return "", err
	}

	// 8. Broadcast the Transaction
	err = n.Client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", err
	}

	receipt, err := bind.WaitMined(ctx, n.Client, signedTx)
	if err != nil {
		return "", err
	}

	log.Info("TRANSFER_ETH", "TIME_PROCESS", map[string]interface{}{
		"process_time": time.Since(start).Seconds(),
		"tx_hash":      tx.Hash().Hex(),
		"status":       receipt.Status,
		"node":         n.URL,
	})

	if receipt.Status == 1 {
	} else {
		return tx.Hash().Hex(), errors.New("transaction failed")
	}

	return tx.Hash().Hex(), nil
}

func (n *Node) GetAllowanceERC20(ctx context.Context, walletAddress string, tokenAddress string, spenderAddress string) (*big.Int, error) {
	erc20Svc, err := defineabi.NewERC20(common.HexToAddress(tokenAddress), n.Client)
	if err != nil {
		return nil, err
	}

	allowance, err := erc20Svc.Allowance(nil, common.HexToAddress(walletAddress), common.HexToAddress(spenderAddress))
	if err != nil {
		return nil, err
	}

	return allowance, nil
}
