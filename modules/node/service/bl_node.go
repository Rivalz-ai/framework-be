package service

import (
	"context"
	"errors"

	//"fmt"
	marketplaceDto "github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	"github.com/Rivalz-ai/framework-be/modules/node/dto"
)

func (b *NodeBalancer) GetChainID() (int, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.GetChainID()
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errors.New("All nodes provider are failed")
}
func (b *NodeBalancer) CallMethod(ctx context.Context, method, walletAddress, token_address, ABI string, args ...interface{}) (interface{}, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.CallMethod(ctx, method, walletAddress, token_address, ABI, args...)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}
func (b *NodeBalancer) GetRewardTokenInfo(ctx context.Context, tokenAddress, ABI string) (*dto.TokenInfo, error) {
	excluded := make(map[string]bool)
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.GetRewardTokenInfo(ctx, tokenAddress, ABI)
		if err != nil {
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errors.New("All nodes provider are failed")
}

func (b *NodeBalancer) BroadcastTransaction(ctx context.Context, fromWallet string, swapResponse *marketplaceDto.OkxSwapResponse, privateKeyHex string) (string, error) {
	excluded := make(map[string]bool)
	var errCatched error = errors.New("internal error")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := n.BroadcastTransaction(ctx, fromWallet, swapResponse, privateKeyHex)
		if err != nil {
			excluded[n.Key] = true
			errCatched = err
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed: " + errCatched.Error())
}

func (b *NodeBalancer) BroadcastSignedTransaction(ctx context.Context, signedTransaction string) (string, error) {
	excluded := make(map[string]bool)
	var errCatched error = errors.New("internal error")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := n.BroadcastSignedTransaction(ctx, signedTransaction)
		if err != nil {
			excluded[n.Key] = true
			errCatched = err
			continue
		}
		return result, nil
	}
	return "", errors.New("All nodes provider are failed: " + errCatched.Error())
}
