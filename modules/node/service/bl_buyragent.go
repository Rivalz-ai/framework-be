package service

import (
	"context"
	"errors"
	"math/big"
)

func (n *NodeBalancer) BuyRagent(ctx context.Context, secretKey string, tokenAddress string, quantity *big.Int, totalPrice *big.Int, nonceSign *big.Int, signature string) (string, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(n.nodes); i++ {
		node, err := n.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := node.BuyRagent(ctx, secretKey, tokenAddress, quantity, totalPrice, nonceSign, signature)
		if err != nil {
			excluded[node.Key] = true
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed")
}
