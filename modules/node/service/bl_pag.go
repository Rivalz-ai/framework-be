package service

import (
	"context"
	"errors"
	//"fmt"
	//"github.com/Rivalz-ai/framework-be/modules/node/dto"
)

func (b *NodeBalancer) TokenOfOwnerByIndex(ctx context.Context, walletAddress, token_address string, index int) (int64, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.TokenOfOwnerByIndex(ctx, walletAddress, token_address, index)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errors.New("All nodes provider are failed")
}
