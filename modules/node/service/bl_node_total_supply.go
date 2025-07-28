package service

import (
	"context"
	"errors"
)

func (b *NodeBalancer) GetTotalSupply(ctx context.Context, tokenAddress string, ABI string) (string, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ {
		n, err := b.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := n.GetTotalSupply(ctx, tokenAddress, ABI)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) GetDecimals(ctx context.Context, tokenAddress string, ABI string) (int, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ {
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.GetDecimals(ctx, tokenAddress, ABI)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errors.New("All nodes provider are failed")
}
