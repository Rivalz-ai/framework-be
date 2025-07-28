package dto
import (	
	//"time"
	"math/big"
)
/**/
type StakeInfo struct {
	TokenId    *big.Int
	StartBlock *big.Int
	IsStaked   bool
}
type PendingReward struct {
	TokenAddress string `json:"token_address" bson:"token_address"`
	Amount       int64  `json:"amount" bson:"amount"`
}
type TokenInfo struct {
	RewardAddress string
	Decimals      uint8
	Name          string
}

type Erc20Info struct {
	Address  string
	Decimals uint8
	Name     string
	Symbol   string
}
