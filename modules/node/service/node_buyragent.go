package service

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	defineabi "github.com/Rivalz-ai/framework-be/define/abi"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (n *Node) BuyRagent(ctx context.Context, secretKey string, tokenAddress string, quantity *big.Int, totalPrice *big.Int, nonceSign *big.Int, signature string) (string, error) {
	erc20Svc, err := defineabi.NewRAgentMarket(common.HexToAddress(tokenAddress), n.Client)
	if err != nil {
		log.Error("error when new erc20 service: "+err.Error(), "TransferERC20-Erc20Service")
		return "", err
	}

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
	auth.GasLimit = uint64(3000000) // Standard ERC20 transfer uses ~65,000 gas

	// Update gas price with a tip to prioritize the transaction
	tipCap, err := n.Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return "", err
	}
	feeCap := new(big.Int).Add(gasPrice, new(big.Int).Mul(tipCap, big.NewInt(2)))

	auth.GasFeeCap = feeCap
	auth.GasTipCap = tipCap

	signBytes, err := hex.DecodeString(signature[2:])
	if err != nil {
		return "", err
	}
	rAgentAddr := common.HexToAddress("0x73e875F602acfCa989b05b2A7D838c68F3ABeEE5")
	tx, err := erc20Svc.BuyRagent(auth, rAgentAddr, quantity, totalPrice, nonceSign, signBytes)
	if err != nil {
		log.Error("error when buy ragent: "+err.Error(), "BuyRagent-Erc20Service")
		return "", err
	}

	fmt.Println(tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, n.Client, tx)
	if err != nil {
		return "", err
	}

	// fmt.Printf("Transaction status: %d\n", receipt.Status)
	if receipt.Status == 1 {
		// fmt.Println("Transaction succeeded")
	} else {
		return tx.Hash().Hex(), errors.New("transaction failed")
	}

	return tx.Hash().Hex(), nil
}
