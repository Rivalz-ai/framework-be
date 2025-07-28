package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"
	"time"

	defineabi "github.com/Rivalz-ai/framework-be/define/abi"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (n *Node) TransferERC721(ctx context.Context, secretKey string, tokenAddress string, toAddress string, tokenId *big.Int) (string, error) {
	erc20Svc, err := defineabi.NewERC721(common.HexToAddress(tokenAddress), n.Client)
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
	auth.Value = big.NewInt(0)     // No ETH value needed for ERC20 transfers
	auth.GasLimit = uint64(100000) // Standard ERC20 transfer uses ~65,000 gas
	auth.GasPrice = gasPrice

	tx, err := erc20Svc.TransferFrom(auth, common.HexToAddress(fromAddress.Hex()), common.HexToAddress(toAddress), tokenId)
	if err != nil {
		log.Error("error when transfer erc20: "+err.Error(), "TransferERC20-Erc20Service")
		return "", err
	}

	receipt, err := bind.WaitMined(ctx, n.Client, tx)
	if err != nil {
		return "", err
	}

	log.Info("TRANSFER_ERC_721", "TIME_PROCESS", map[string]interface{}{
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

func (n *Node) OwnerOfERC721(ctx context.Context, tokenAddress string, tokenId *big.Int) (string, error) {
	erc20Svc, err := defineabi.NewERC721(common.HexToAddress(tokenAddress), n.Client)
	if err != nil {
		return "", err
	}

	owner, err := erc20Svc.OwnerOf(nil, tokenId)
	if err != nil {
		return "", err
	}

	return strings.ToLower(owner.Hex()), nil
}

func (n *Node) CheckBalanceOfWallet(ctx context.Context, tokenAddress string, walletAddress string) (int64, error) {
	erc20Svc, err := defineabi.NewERC721(common.HexToAddress(tokenAddress), n.Client)
	if err != nil {
		return 0, err
	}

	balance, err := erc20Svc.BalanceOf(nil, common.HexToAddress(walletAddress))
	if err != nil {
		return 0, err
	}

	return balance.Int64(), nil
}

func (n *Node) GetTokenIds(ctx context.Context, tokenAddress string, walletAddress string, quantity int64) ([]*big.Int, error) {
	erc20Svc, err := defineabi.NewERC721(common.HexToAddress(tokenAddress), n.Client)
	if err != nil {
		return nil, err
	}

	tokenIds := make([]*big.Int, quantity)
	for i := int64(0); i < quantity; i++ {
		tokenId, err := erc20Svc.TokenOfOwnerByIndex(nil, common.HexToAddress(walletAddress), big.NewInt(i))
		if err != nil {
			return nil, err
		}
		tokenIds[i] = tokenId
	}

	return tokenIds, nil
}
