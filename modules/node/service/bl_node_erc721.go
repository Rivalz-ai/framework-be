package service

import (
	"context"
	"errors"
	"math/big"
)

func (b *NodeBalancer) TransferERC721(ctx context.Context, secretKey string, tokenAddress string, toAddress string, tokenId *big.Int) (string, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := n.TransferERC721(ctx, secretKey, tokenAddress, toAddress, tokenId)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) OwnerOfERC721(ctx context.Context, tokenAddress string, tokenId *big.Int) (string, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("internal error")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			errCatch = err
			return "", err
		}
		result, err := n.OwnerOfERC721(ctx, tokenAddress, tokenId)
		if err != nil {
			errCatch = err
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed: " + errCatch.Error())
}

func (b *NodeBalancer) CheckBalanceOfWallet(ctx context.Context, tokenAddress string, walletAddress string) (int64, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("internal error")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		balance, err := n.CheckBalanceOfWallet(ctx, tokenAddress, walletAddress)
		if err != nil {
			excluded[n.Key] = true
			errCatch = err
			continue
		}
		return balance, nil
	}
	return 0, errors.New("All nodes provider are failed: " + errCatch.Error())
}

func (b *NodeBalancer) GetTokenIds(ctx context.Context, tokenAddress string, walletAddress string, quantity int64) ([]*big.Int, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("internal error")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			errCatch = err
			return nil, err
		}
		tokenIds, err := n.GetTokenIds(ctx, tokenAddress, walletAddress, quantity)
		if err != nil {
			excluded[n.Key] = true
			errCatch = err
			continue
		}
		return tokenIds, nil
	}
	return nil, errors.New("All nodes provider are failed: " + errCatch.Error())
}
