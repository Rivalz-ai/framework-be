package service

import (
	"context"
	"errors"

	projectDto "github.com/Rivalz-ai/framework-be/modules/project/dto"
)

func (n *NodeBalancer) CreateResourceStakingBatch(ctx context.Context, factoryContract, privateKey string, data *projectDto.ProjectResponse) (map[string]string, bool, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(n.nodes); i++ {
		node, err := n.GetNode(excluded)
		if err != nil {
			return nil, true, err
		}
		result, isRetry, err := node.CreateResourceStakingBatch(ctx, factoryContract, privateKey, data)
		if err != nil {
			excluded[node.Key] = true
			continue
		}
		return result, isRetry, nil
	}
	return nil, false, errors.New("All nodes provider are failed")
}
