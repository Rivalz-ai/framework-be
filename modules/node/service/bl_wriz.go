package service

import (
	"context"
	"errors"
)

func (b *NodeBalancer) WRizCheckNonceInUsed(ctx context.Context, wrizContractAddress string, nonce int64) (bool, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return false, err
		}
		result, err := n.CheckNonceInUsed(ctx, wrizContractAddress, nonce)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return false, errors.New("All nodes provider are failed")
}
