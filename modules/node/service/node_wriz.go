package service

import (
	"context"
	"math/big"

	"github.com/Rivalz-ai/framework-be/define/abi"
	"github.com/ethereum/go-ethereum/common"
)

func (n *Node) CheckNonceInUsed(ctx context.Context, wrizContractAddress string, nonce int64) (bool, error) {

	rClient, err := abi.NewWRizClaim(common.HexToAddress(wrizContractAddress), n.Client)
	if err != nil {
		return false, err
	}

	usedNonces, err := rClient.UsedNonces(nil, big.NewInt(nonce))
	if err != nil {
		return false, err
	}

	return usedNonces, nil
}
