package service

import (
	"context"
)

func (n *Node) GetBlockNumber() (int, error) {
	blockNumber, err := n.Client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return int(blockNumber), nil
}
