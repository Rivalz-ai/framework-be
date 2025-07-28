package service

import (
	"context"
	"errors"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/log"
	commonType "github.com/Rivalz-ai/framework-be/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	//"fmt"
)

func (n *Node) GetTransactionReceiptByHash(ctx context.Context, txHash string) (*commonType.TxWrap, *types.Receipt, error) {
	// Chuyển đổi txHash thành kiểu common.Hash
	hash := common.HexToHash(txHash)
	//Lấy thông tin transaction
	tx, isPending, err := n.Client.TransactionByHash(context.Background(), hash)
	_ = isPending
	if err != nil {
		log.Error("Can not get transaction: "+err.Error(), "GetTransactionReceiptByHash", txHash)
		return nil, nil, err
	}
	// Lấy receipt của transaction
	receipt, err := n.Client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		log.Error("Can not get transaction receipt: "+err.Error(), "GetTransactionReceiptByHash", txHash)
		return nil, nil, err
	}
	// Lấy chainID
	chainID, err := n.Client.NetworkID(context.Background())
	if err != nil {
		log.Error("Can not get chainID: "+err.Error(), "GetTransactionReceiptByHash", txHash)
		return nil, nil, err
	}
	//get signer
	signer := types.LatestSignerForChainID(chainID)
	//get from address
	from, err := types.Sender(signer, tx)
	if err != nil {
		log.Error("Can not get From Address: "+err.Error(), "GetTransactionReceiptByHash", txHash)
		return nil, nil, err
	}
	txData := &commonType.TxWrap{
		From:  strings.ToLower(from.Hex()),
		To:    strings.ToLower(tx.To().Hex()),
		EthTx: tx,
	}
	//fmt.Printf("Receipt: %v\n", receipt)
	//In thông tin của receipt log
	/*
		for i, logEntry := range receipt.Logs {
			fmt.Printf("\nLog %d:\n", i)
			fmt.Printf("  Address: %s\n", logEntry.Address.Hex())
			fmt.Printf("  Topics: %v\n", logEntry.Topics)
			fmt.Printf("  Data: %x\n", logEntry.Data)
			fmt.Printf("  Block Number: %d\n", logEntry.BlockNumber)
			fmt.Printf("  Tx Index: %d\n", logEntry.TxIndex)
			fmt.Printf("  Log Index: %d\n", logEntry.Index)
		}
	*/
	return txData, receipt, nil
}
func (n *Node) ParseLogs(ctx context.Context, method, ABI string, logs []*types.Log) ([]*types.Log, error) {
	parsedABI, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		return nil, errors.New("Failed to load the contract ABI")
	}
	var logs_resp []*types.Log
	for _, vLog := range logs {
		// Kiểm tra log có phải là event Transfer không
		eventID := parsedABI.Events[method].ID
		if vLog.Topics[0] != eventID {
			continue
		}
		// Decode value từ Data
		var value interface{}
		err := parsedABI.UnpackIntoInterface(&value, method, vLog.Data)
		if err != nil {
			continue
		}
		logs_resp = append(logs_resp, vLog)
	}
	return logs_resp, nil
}
