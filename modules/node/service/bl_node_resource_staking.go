package service

import (
	"context"
	"errors"

	"github.com/Rivalz-ai/framework-be/modules/node/dto"
	//"fmt"
)

func (b *NodeBalancer) GetResourceStakeBalance(ctx context.Context, wallet_address, staking_contract_address string) ([]dto.StakeInfo, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.GetResourceStakeBalance(ctx, wallet_address, staking_contract_address)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) GetTokenRewards(ctx context.Context, staking_contract_address string) ([]string, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}

		result, err := n.GetTokenRewards(ctx, staking_contract_address)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) GetTotalStakedTokens(ctx context.Context, staking_contract_address string) (int64, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.GetTotalStakedTokens(ctx, staking_contract_address)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errors.New("All nodes provider are failed")
}
func (b *NodeBalancer) GetPendingRewards(ctx context.Context, wallet_address, staking_contract_address string) ([]dto.PendingReward, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.GetPendingRewards(ctx, wallet_address, staking_contract_address)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}
func (b *NodeBalancer) GetRewardPerMonth(ctx context.Context, staking_contract_address, reward_token_address string) (int64, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.GetRewardPerMonth(ctx, staking_contract_address, reward_token_address)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) GetRewardPerWeek(ctx context.Context, staking_contract_address, reward_token_address string) (int64, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.GetRewardPerWeek(ctx, staking_contract_address, reward_token_address)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) SBlocksPerDay(ctx context.Context, staking_contract_address string) (int64, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.SBlocksPerDay(ctx, staking_contract_address)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) GetResourceStakeBalanceBatch(ctx context.Context, wallet_address string, staking_contract_address []string, multicall_contract_address string) (int64, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.GetResourceStakeBalanceBatch(ctx, wallet_address, staking_contract_address, multicall_contract_address)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errors.New("All nodes provider are failed")
}
