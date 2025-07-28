package service

import (
	"context"
	"errors"
	"math/big"

	"github.com/Rivalz-ai/framework-be/modules/node/dto"
)

func (b *NodeBalancer) GetERC20TokenInfo(ctx context.Context, tokenAddress string) (*dto.Erc20Info, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.GetERC20TokenInfo(ctx, tokenAddress)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) TransferERC20(ctx context.Context, secretKey string, tokenAddress string, toAddress string, amount *big.Int) (string, error) {
	excluded := make(map[string]bool)
	var errCatched error = errors.New("internal error")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := n.TransferERC20(ctx, secretKey, tokenAddress, toAddress, amount)
		if err != nil {
			excluded[n.Key] = true
			errCatched = err
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed: " + errCatched.Error())
}

func (b *NodeBalancer) ApproveERC20(ctx context.Context, secretKey string, spenderAddress string, tokenAddress string, amount *big.Int) (string, error) {
	excluded := make(map[string]bool)
	var errCatched error = errors.New("internal error")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := n.ApproveERC20(ctx, secretKey, spenderAddress, tokenAddress, amount)
		if err != nil {
			excluded[n.Key] = true
			errCatched = err
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed: " + errCatched.Error())
}

func (b *NodeBalancer) GetBalanceERC20(ctx context.Context, walletAddress string, tokenAddress string) (*big.Int, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.GetBalanceERC20(ctx, walletAddress, tokenAddress)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) GetBalanceETH(ctx context.Context, walletAddress string) (*big.Int, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.GetBalanceETH(ctx, walletAddress)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) TransferETH(ctx context.Context, secretKey string, toAddress string, amount *big.Int) (string, error) {
	excluded := make(map[string]bool)
	var errCatched error = errors.New("internal error")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := n.TransferETH(ctx, secretKey, toAddress, amount)
		if err != nil {
			excluded[n.Key] = true
			errCatched = err
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed: " + errCatched.Error())
}

func (b *NodeBalancer) GetAllowanceERC20(ctx context.Context, walletAddress string, tokenAddress string, spenderAddress string) (*big.Int, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.GetAllowanceERC20(ctx, walletAddress, tokenAddress, spenderAddress)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}
