package service

import (
	"context"
	"errors"
	"math/big"
)

func (n *NodeBalancer) LendingGetBorrowFeePerMonth(ctx context.Context, lendingAddress string) (*big.Int, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(n.nodes); i++ {
		node, err := n.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := node.LendingGetBorrowFeePerMonth(ctx, lendingAddress)
		if err != nil {
			excluded[node.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}

func (n *NodeBalancer) LendingTotalLendingAmount(ctx context.Context, lendingAddress string) (*big.Int, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(n.nodes); i++ {
		node, err := n.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := node.LendingTotalLendingAmount(ctx, lendingAddress)
		if err != nil {
			excluded[node.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}

func (n *NodeBalancer) LendingGetTotalBorrowAmount(ctx context.Context, lendingAddress string) (*big.Int, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(n.nodes); i++ {
		node, err := n.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := node.LendingGetTotalBorrowAmount(ctx, lendingAddress)
		if err != nil {
			excluded[node.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}

func (n *NodeBalancer) LendingGetBorrowInfoOfBorrower(ctx context.Context, lendingAddress string, borrower string) ([]map[string]interface{}, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(n.nodes); i++ {
		node, err := n.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := node.LendingGetBorrowInfoOfBorrower(ctx, lendingAddress, borrower)
		if err != nil {
			excluded[node.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}

func (n *NodeBalancer) LendingPendingReward(ctx context.Context, lendingAddress string, lender string) (map[string]interface{}, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(n.nodes); i++ {
		node, err := n.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := node.LendingPendingReward(ctx, lendingAddress, lender)
		if err != nil {
			excluded[node.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}
