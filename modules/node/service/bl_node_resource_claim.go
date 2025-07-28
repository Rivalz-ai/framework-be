package service

import (
	//"github.com/Rivalz-ai/framework-be/modules/node/dto"
	"context"
	"errors"
	//"fmt"
)

func (b *NodeBalancer) ClaimedResourceAmount(ctx context.Context, resource_claimer_address, wallet_address string) (int64, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.ClaimedResourceAmount(ctx, resource_claimer_address, wallet_address)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errors.New("All nodes provider are failed")
}
