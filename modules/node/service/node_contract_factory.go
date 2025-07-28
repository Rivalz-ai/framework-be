package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	romeAbi "github.com/Rivalz-ai/framework-be/define/abi"
	projectDto "github.com/Rivalz-ai/framework-be/modules/project/dto"
)

func (n *Node) CreateResourceStakingBatch(ctx context.Context, factoryContract, privateKeyString string, data *projectDto.ProjectResponse) (map[string]string, bool, error) {

	contract, err := romeAbi.NewStakingFactory(common.HexToAddress(factoryContract), n.Client)
	if err != nil {
		return nil, false, err
	}

	var (
		tokenAddresses []common.Address
		types          []uint8
		startBlock     *big.Int
		endBlock       *big.Int
	)

	for _, tokenAddress := range data.TokenAddresses {
		tokenAddresses = append(tokenAddresses, common.HexToAddress(tokenAddress))
	}

	for _, t := range data.Types {
		tInt8, err := strconv.ParseUint(t, 10, 8)
		if err != nil {
			return nil, false, err
		}
		types = append(types, uint8(tInt8))
	}

	startBlock = big.NewInt(int64(data.StartBlock))
	endBlock = big.NewInt(int64(data.EndBlock))

	// prepare for sign transaction
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return nil, false, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, false, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := n.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, false, err
	}

	gasPrice, err := n.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, false, err
	}

	// get chain id
	chainID, err := n.Client.ChainID(context.Background())
	if err != nil {
		return nil, false, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, false, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(15000000) // in units
	auth.GasPrice = gasPrice

	tx, err := contract.CreateResourceStakingBatch(auth, tokenAddresses, types, data.ProjectId, startBlock, endBlock)
	if err != nil {
		return nil, true, err
	}

	// Wait for transaction receipt
	receipt, err := bind.WaitMined(ctx, n.Client, tx)
	if err != nil {
		return nil, true, err
	}

	// fmt.Printf("Transaction status: %d\n", receipt.Status)
	if receipt.Status == 1 {
		// fmt.Println("Transaction succeeded")
	} else {
		return nil, false, errors.New("transaction failed")
	}

	// Decode return data
	parsedABI, err := abi.JSON(strings.NewReader(romeAbi.StakingFactoryMetaData.ABI))
	if err != nil {
		return nil, false, err
	}

	var mapStakingAddresses = make(map[string]string)
	if len(receipt.Logs) > 0 {
		for _, log := range receipt.Logs {
			if log.Topics[0] != parsedABI.Events["ResourceStakingCreated"].ID {
				continue
			}

			if len(log.Topics) < 4 {
				continue
			}

			// parse log
			var (
				stakingAddress string
				// creatorAddress string
				resourceToken string
			)

			stakingAddress = log.Topics[1].Hex()[:2] + log.Topics[1].Hex()[26:]
			// creatorAddress = log.Topics[2].Hex()[:2] + log.Topics[2].Hex()[26:]
			resourceToken = log.Topics[3].Hex()[:2] + log.Topics[3].Hex()[26:]

			mapStakingAddresses[resourceToken] = stakingAddress
		}
	}

	return mapStakingAddresses, false, nil
}
