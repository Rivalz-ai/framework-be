package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

func (n *Node) GetSolanaBalance(ctx context.Context, address string) (float64, error) {
	rpcClient := rpc.New(n.URL)
	pubKey, err := solana.PublicKeyFromBase58(address)
	if err != nil {
		return 0, err
	}
	bal, err := rpcClient.GetBalance(ctx, pubKey, rpc.CommitmentFinalized)
	if err != nil {
		return 0, err
	}
	return float64(bal.Value) / float64(solana.LAMPORTS_PER_SOL), nil
}

// GetSolanaTokenBalance gets the balance of a specific SPL token for a wallet address
func (n *Node) GetSolanaTokenBalance(ctx context.Context, walletAddress string, tokenMintAddress string) (uint64, error) {
	rpcClient := rpc.New(n.URL)

	// Parse wallet address
	walletPubKey, err := solana.PublicKeyFromBase58(walletAddress)
	if err != nil {
		return 0, fmt.Errorf("invalid wallet address: %v", err)
	}

	// Parse token mint address
	tokenMint, err := solana.PublicKeyFromBase58(tokenMintAddress)
	if err != nil {
		return 0, fmt.Errorf("invalid token mint address: %v", err)
	}

	// Find the associated token account address
	associatedTokenAccount, _, err := solana.FindAssociatedTokenAddress(walletPubKey, tokenMint)
	if err != nil {
		return 0, fmt.Errorf("failed to find associated token account: %v", err)
	}

	// Get token account balance
	balance, err := rpcClient.GetTokenAccountBalance(ctx, associatedTokenAccount, rpc.CommitmentFinalized)
	if err != nil {
		// If account doesn't exist, return 0 balance
		return 0, nil
	}

	// Convert string amount to uint64
	amount, err := strconv.ParseUint(balance.Value.Amount, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse token balance: %v", err)
	}

	return amount, nil
}

func (n *Node) SendSolanaTransaction(ctx context.Context, signedTxBytes []byte) (string, error) {
	rpcClient := rpc.New(n.URL)

	// Deserialize the signed transaction
	signedTx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(signedTxBytes))
	if err != nil {
		return "", fmt.Errorf("decode signed tx error: %v", err)
	}

	// Send the transaction to Solana
	sig, err := rpcClient.SendTransaction(ctx, signedTx)
	if err != nil {
		return sig.String(), fmt.Errorf("send tx error: %v", err)
	}

	// return sig.String(), nil

	// Wait for transaction confirmation with retry mechanism
	maxRetries := 30
	retryDelay := 300 * time.Millisecond

	for i := 0; i < maxRetries; i++ {
		// Check transaction status
		txStatus, err := rpcClient.GetTransaction(ctx, sig, &rpc.GetTransactionOpts{
			Commitment: rpc.CommitmentFinalized,
		})

		if err != nil {
			// If transaction not found yet, wait and retry
			if i < maxRetries-1 {
				time.Sleep(retryDelay)
				continue
			}
			return sig.String(), fmt.Errorf("get tx status error after %d retries: %v", maxRetries, err)
		}

		// Check if transaction failed
		if txStatus.Meta != nil && txStatus.Meta.Err != nil {
			return sig.String(), fmt.Errorf("transaction failed: %v", txStatus.Meta.Err)
		}

		// Transaction found and successful
		// sleep to make sure the transaction is confirmed
		time.Sleep(time.Second * 1)
		return sig.String(), nil
	}

	return "", fmt.Errorf("transaction confirmation timeout after %d retries", maxRetries)
}

func (n *Node) GetLatestBlockhash(ctx context.Context) (*rpc.GetLatestBlockhashResult, error) {
	rpcClient := rpc.New(n.URL)
	return rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
}

// ApproveSPLToken approves another account to spend tokens on behalf of the owner
func (n *Node) ApproveSPLToken(ctx context.Context, ownerAddress string, tokenMintAddress string, delegateAddress string, amount uint64) (*solana.Transaction, error) {
	// Parse addresses
	ownerPubKey, err := solana.PublicKeyFromBase58(ownerAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid owner address: %v", err)
	}

	tokenMint, err := solana.PublicKeyFromBase58(tokenMintAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid token mint address: %v", err)
	}

	delegatePubKey, err := solana.PublicKeyFromBase58(delegateAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid delegate address: %v", err)
	}

	// Find the associated token account for the owner
	ownerTokenAccount, _, err := solana.FindAssociatedTokenAddress(ownerPubKey, tokenMint)
	if err != nil {
		return nil, fmt.Errorf("failed to find owner's token account: %v", err)
	}

	// Get latest blockhash
	rpcClient := rpc.New(n.URL)
	latestBlockhash, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest blockhash: %v", err)
	}

	// Create approve instruction
	approveInstruction := token.NewApproveInstruction(
		amount,
		ownerTokenAccount,
		delegatePubKey,
		ownerPubKey,
		[]solana.PublicKey{}, // No additional signers
	).Build()

	// Create transaction
	tx, err := solana.NewTransaction(
		[]solana.Instruction{approveInstruction},
		latestBlockhash.Value.Blockhash,
		solana.TransactionPayer(ownerPubKey),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}

	return tx, nil
}

// RevokeSPLTokenApproval revokes approval for a delegate to spend tokens
func (n *Node) RevokeSPLTokenApproval(ctx context.Context, ownerAddress string, tokenMintAddress string) (*solana.Transaction, error) {
	// Parse addresses
	ownerPubKey, err := solana.PublicKeyFromBase58(ownerAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid owner address: %v", err)
	}

	tokenMint, err := solana.PublicKeyFromBase58(tokenMintAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid token mint address: %v", err)
	}

	// Find the associated token account for the owner
	ownerTokenAccount, _, err := solana.FindAssociatedTokenAddress(ownerPubKey, tokenMint)
	if err != nil {
		return nil, fmt.Errorf("failed to find owner's token account: %v", err)
	}

	// Get latest blockhash
	rpcClient := rpc.New(n.URL)
	latestBlockhash, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest blockhash: %v", err)
	}

	// Create revoke instruction
	revokeInstruction := token.NewRevokeInstruction(
		ownerTokenAccount,
		ownerPubKey,
		[]solana.PublicKey{}, // No additional signers
	).Build()

	// Create transaction
	tx, err := solana.NewTransaction(
		[]solana.Instruction{revokeInstruction},
		latestBlockhash.Value.Blockhash,
		solana.TransactionPayer(ownerPubKey),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}

	return tx, nil
}

// GetSPLTokenApproval gets the current approval amount for a delegate
func (n *Node) GetSPLTokenApproval(ctx context.Context, ownerAddress string, tokenMintAddress string) (uint64, *solana.PublicKey, error) {
	rpcClient := rpc.New(n.URL)

	// Parse addresses
	ownerPubKey, err := solana.PublicKeyFromBase58(ownerAddress)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid owner address: %v", err)
	}

	tokenMint, err := solana.PublicKeyFromBase58(tokenMintAddress)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid token mint address: %v", err)
	}

	// Find the associated token account for the owner
	ownerTokenAccount, _, err := solana.FindAssociatedTokenAddress(ownerPubKey, tokenMint)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to find owner's token account: %v", err)
	}

	// Get token account info
	accountInfo, err := rpcClient.GetAccountInfo(ctx, ownerTokenAccount)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get token account info: %v", err)
	}

	if accountInfo.Value == nil {
		return 0, nil, fmt.Errorf("token account not found")
	}

	// Decode token account data
	var tokenAccount token.Account
	err = bin.NewBinDecoder(accountInfo.Value.Data.GetBinary()).Decode(&tokenAccount)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to decode token account: %v", err)
	}

	// Check if there's an active delegation
	if tokenAccount.Delegate == nil {
		return 0, nil, nil // No active delegation
	}

	return tokenAccount.DelegatedAmount, tokenAccount.Delegate, nil
}
