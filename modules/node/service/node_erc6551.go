package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"time"

	defineabi "github.com/Rivalz-ai/framework-be/define/abi"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (n *Node) TransferERC6551(ctx context.Context, secretKey string, contractAddress string, tokenAddress string, callData []byte) (string, error) {
	erc6551Svc, err := defineabi.NewERC6551(common.HexToAddress(contractAddress), n.Client)
	if err != nil {
		log.Error("error when new erc6551 service: "+err.Error(), "TransferERC6551-Erc6551Service")
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
	auth.Value = big.NewInt(0)     // No ETH value needed for ERC20 transfers
	auth.GasLimit = uint64(200000) // Standard ERC20 transfer uses ~65,000 gas
	auth.GasPrice = gasPrice

	tx, err := erc6551Svc.Execute(auth, common.HexToAddress(tokenAddress), big.NewInt(0), callData, 0)
	if err != nil {
		log.Error("error when transfer erc6551: "+err.Error(), "TransferERC6551-Erc6551Service")
		return "", err
	}

	receipt, err := bind.WaitMined(ctx, n.Client, tx)
	if err != nil {
		return "", err
	}

	log.Info("TRANSFER_ERC_6551", "TIME_PROCESS", map[string]interface{}{
		"process_time": time.Since(start).Seconds(),
		"tx_hash":      tx.Hash().Hex(),
		"status":       receipt.Status,
		"node":         n.URL,
	})

	// fmt.Printf("Transaction status: %d\n", receipt.Status)
	if receipt.Status == 1 {
		// fmt.Println("Transaction succeeded")
	} else {
		return "", errors.New("transaction failed")
	}

	return tx.Hash().Hex(), nil
}
