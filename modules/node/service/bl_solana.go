package service

import (
	"context"
	"errors"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func (b *NodeBalancer) GetTokenBalanceSolana(ctx context.Context, walletAddress string) (float64, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("All nodes provider are failed")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.GetSolanaBalance(ctx, walletAddress)
		if err != nil {
			errCatch = err
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errCatch
}

func (b *NodeBalancer) SendSolanaTransaction(ctx context.Context, signedTxBytes []byte) (string, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("All nodes provider are failed")
	var resultStr string
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return "", err
		}
		result, err := n.SendSolanaTransaction(ctx, signedTxBytes)
		if err != nil {
			errCatch = err
			excluded[n.Key] = true
			resultStr = result
			continue
		}
		return result, nil
	}
	return resultStr, errCatch
}

func (b *NodeBalancer) GetLatestBlockhashSolana(ctx context.Context) (*rpc.GetLatestBlockhashResult, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("All nodes provider are failed")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.GetLatestBlockhash(ctx)
		if err != nil {
			errCatch = err
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errCatch
}

func (b *NodeBalancer) GetSolanaTokenBalance(ctx context.Context, walletAddress string, tokenMintAddress string) (uint64, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("All nodes provider are failed")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, err
		}
		result, err := n.GetSolanaTokenBalance(ctx, walletAddress, tokenMintAddress)
		if err != nil {
			errCatch = err
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return 0, errCatch
}

func (b *NodeBalancer) ApproveSPLToken(ctx context.Context, ownerAddress string, tokenMintAddress string, delegateAddress string, amount uint64) (*solana.Transaction, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("All nodes provider are failed")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.ApproveSPLToken(ctx, ownerAddress, tokenMintAddress, delegateAddress, amount)
		if err != nil {
			errCatch = err
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errCatch
}

func (b *NodeBalancer) RevokeSPLTokenApproval(ctx context.Context, ownerAddress string, tokenMintAddress string) (*solana.Transaction, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("All nodes provider are failed")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return nil, err
		}
		result, err := n.RevokeSPLTokenApproval(ctx, ownerAddress, tokenMintAddress)
		if err != nil {
			errCatch = err
			excluded[n.Key] = true
			continue
		}
		return result, nil
	}
	return nil, errCatch
}

func (b *NodeBalancer) GetSPLTokenApproval(ctx context.Context, ownerAddress string, tokenMintAddress string) (uint64, *solana.PublicKey, error) {
	excluded := make(map[string]bool)
	errCatch := errors.New("All nodes provider are failed")
	for i := 0; i < len(b.nodes); i++ { //number of retry time
		n, err := b.GetNode(excluded)
		if err != nil {
			return 0, nil, err
		}
		result, _, err := n.GetSPLTokenApproval(ctx, ownerAddress, tokenMintAddress)
		if err != nil {
			errCatch = err
			excluded[n.Key] = true
			continue
		}
		return result, nil, nil
	}
	return 0, nil, errCatch
}
