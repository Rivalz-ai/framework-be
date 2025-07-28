package service

import (
	"errors"
)

func (b *NodeBalancer) GetBlockNumber() (int, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.GetBlockNumber()
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errors.New("All nodes provider are failed")
}
