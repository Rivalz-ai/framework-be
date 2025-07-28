package service

import (
	"context"
	"errors"

	commonType "github.com/Rivalz-ai/framework-be/types"
	"github.com/ethereum/go-ethereum/core/types"
	//"fmt"
)

func (b *NodeBalancer) GetTransactionReceiptByHash(ctx context.Context, txHash string) (*commonType.TxWrap, *types.Receipt, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, nil, err
		}
		result1, result2, err := n.GetTransactionReceiptByHash(ctx, txHash)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result1, result2, nil
	}
	return nil, nil, errors.New("All nodes provider are failed")
}
func (b *NodeBalancer) ParseLogs(ctx context.Context, method, ABI string, logs []*types.Log) ([]*types.Log, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.ParseLogs(ctx, method, ABI, logs)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}
