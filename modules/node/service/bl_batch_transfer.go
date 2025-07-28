package service

import (
	"context"
	"errors"
	"math/big"
)

func (b *NodeBalancer) BatchTransferERC20(ctx context.Context, secretKey string, contractAddress string, tokenAddress string, toAddressMap map[string]*big.Int) (string, error) {
	excluded := make(map[string]bool)
	retryCount := uint64(1)
	var errCatched error
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := n.BatchTransferERC20(ctx, secretKey, contractAddress, tokenAddress, toAddressMap, retryCount)
		if err != nil {
			retryCount++
			excluded[n.Key] = true
			errCatched = err
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed: " + errCatched.Error())
}

func (b *NodeBalancer) BatchUnwrapFromErc6551(ctx context.Context, secretKey string, contractAddress string, contractAddressList []string, toAddressMap []string, datumList [][]byte) (string, error) {
	excluded := make(map[string]bool)
	retryCount := uint64(1)
	var errCatched error
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := n.BatchUnwrapFromErc6551(ctx, secretKey, contractAddress, contractAddressList, toAddressMap, datumList, retryCount)
		if err != nil {
			retryCount++
			excluded[n.Key] = true
			errCatched = err
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed: " + errCatched.Error())
}
