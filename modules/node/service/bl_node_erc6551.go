package service

import (
	"context"
	"errors"
)

func (b *NodeBalancer) TransferERC6551(ctx context.Context, secretKey string, contractAddress string, tokenAddress string, callData []byte) (string, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("internal error")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			errCatch = err
			return "", err
		}
		result, err := n.TransferERC6551(ctx, secretKey, contractAddress, tokenAddress, callData)
		if err != nil {
			excluded[n.Key] = true
			errCatch = err
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed: " + errCatch.Error())
}
