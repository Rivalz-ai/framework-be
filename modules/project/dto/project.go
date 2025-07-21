package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//"github.com/ethereum/go-ethereum/common"
)

/*
	{
	    "_id" : ObjectId("67939f2c642d85777429a866"),
	    "name" : "Rivalz",
	    "description" : "World Abstraction Layer for AI",
	    "logo" : "https://api.rivalz.ai/fragmentz/rome/project/rivalz_black_logo.png",
	    "website" : "https://rivalz.ai",
	    "xUsername" : "https://x.com/Rivalz_AI",
	    "discordUsername" : "example#1234",
	    "tokenAddress" : "0x67543CF0304C19CA62AC95ba82FD4F4B40788dc1",
	    "tokenType" : "ERC20",
	    "tokenDecimals" : NumberInt(8),
	    "tokenSymbol" : "RIZ",
	    "chainId" : NumberInt(8453),
	    "categoryId" : "",
	    "status" : "active"
	}
*/
type Project struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	Type            string             `json:"type" bson:"type"`
	Logo            string             `json:"logo" bson:"logo"`
	Website         string             `json:"website" bson:"website"`
	XUsername       string             `json:"x_username" bson:"xUsername"`
	DiscordUsername string             `json:"discord_username" bson:"discordUsername"`
	TokenAddress    string             `json:"token_address" bson:"tokenAddress"`
	TokenType       string             `json:"token_type" bson:"tokenType"`
	TokenDecimals   int                `json:"token_decimals" bson:"tokenDecimals"`
	TokenSymbol     string             `json:"token_symbol" bson:"tokenSymbol"`
	TokenLogo       string             `json:"token_logo" bson:"token_logo"`
	ChainId         int                `json:"chain_id" bson:"chainId"`
	CategoryId      string             `json:"category_id" bson:"categoryId"`
	Status          string             `json:"status" bson:"status"`
	IsVerified      bool               `json:"is_verified" bson:"is_verified"`
	CreatedTime     *time.Time         `json:"created_at" bson:"createdAt"`
	UpdatedTime     *time.Time         `json:"updated_at" bson:"updatedAt"`
	Creator         string             `json:"creator" bson:"creator"`
	ProgressStatus  string             `json:"progress_status" bson:"progress_status"`
	UniqueName      string             `json:"unique_name" bson:"unique_name"`
	TxHash          string             `json:"-" bson:"tx_hash"`
	Source          string             `json:"-" bson:"source"`
	TotalStake      int64              `json:"total_stake" bson:"total_stake"`
	SwarmLevel      string             `json:"swarm_level" bson:"swarm_level"`
	IsHighLight     bool               `json:"is_high_light" bson:"is_high_light"`
}

/*
Resource staking to project

	{
	    "_id" : ObjectId("67ada91e0d793436faf95d5b"),
	    "contractAddress" : "0x99d0c977fdf9938178cc7dcd43fe5bb86e4b0dd8",
	    "projectId" : "6772239afe88f0266465cbf7",
	    "resourceId" : "6777d851f0048df8529a32d7",
	    "createdAt" : ISODate("2025-02-13T08:11:10.446+0000"),
	    "updatedAt" : ISODate("2025-02-13T08:11:10.446+0000"),
	    "__v" : NumberInt(0)
	}
*/
type ResourceStaking struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	ContractAddress string             `json:"contract_address" bson:"contractAddress"`
	ProjectId       string             `json:"project_id" bson:"projectId"`
	ResourceId      string             `json:"resource_id" bson:"resourceId"`
	CreatedAt       *time.Time         `json:"created_at" bson:"createdAt"`
	UpdatedAt       *time.Time         `json:"updated_at" bson:"updatedAt"`
}

/*
	{
	    "_id" : ObjectId("67940d461056081debab7936"),
	    "userId" : "6793fe7d1056081debab6d32",
	    "userWallet" : "0xfa120ca4010f1e1379c6112ce0255b6f734cf01a",
	    "projectId" : "67721c9afe88f0266465cbf6",
	    "resourceId" : "6765770a620ea414f0bfc243",
	    "stakedAmount" : NumberInt(1),
	    "txHash" : "0x4eae954de9542a07fdc6a7f91efb265d913c9dad0b3d55a31b8ff1792aee0341",
	    "createdAt" : ISODate("2025-01-24T21:59:34.868+0000"),
	    "__v" : NumberInt(0)
	}
*/
type UserStaking struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId       string             `json:"user_id" bson:"userId"`
	UserWallet   string             `json:"user_wallet" bson:"userWallet"`
	ProjectId    string             `json:"project_id" bson:"projectId"`
	ResourceId   string             `json:"resource_id" bson:"resourceId"`
	StakedAmount int                `json:"staked_amount" bson:"stakedAmount"`
	TxHash       string             `json:"tx_hash" bson:"txHash"`
	CreatedAt    time.Time          `json:"created_at" bson:"createdAt"`
}
type TokenStakingReward struct {
	ProjectId string   `json:"project_id" bson:"project_id"`
	Rewards   []string `json:"rewards" bson:"rewards"`
	MyStaking int64    `json:"my_staking" bson:"my_staking"` //sô lần stake ko phải balance
}
type SwarmMember struct {
	ProjectId   string `json:"project_id" bson:"project_id"`
	SwarmMember int    `json:"swarm_member" bson:"swarm_member"`
	SwarmRAgent int64  `json:"swarm_r_agent" bson:"swarm_r_agent"`
}
type ProjectInfo struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	Logo            string             `json:"logo" bson:"logo"`
	Website         string             `json:"website" bson:"website"`
	XUsername       string             `json:"x_username" bson:"xUsername"`
	DiscordUsername string             `json:"discord_username" bson:"discordUsername"`
	TokenAddress    string             `json:"token_address" bson:"tokenAddress"`
	TokenType       string             `json:"token_type" bson:"tokenType"`
	TokenDecimals   int                `json:"token_decimals" bson:"tokenDecimals"`
	TokenSymbol     string             `json:"token_symbol" bson:"tokenSymbol"`
	ChainId         int                `json:"chain_id" bson:"chainId"`
	CategoryId      string             `json:"category_id" bson:"categoryId"`
	Status          string             `json:"status" bson:"status"`
	Type            string             `json:"type" bson:"type"`

	CreatedTime *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedTime *time.Time `json:"updated_at" bson:"updatedAt"`
	//
	MarketCap      float64  `json:"market_cap" bson:"market_cap"`
	SwarmRAgent    int64    `json:"swarm_ragent" bson:"swarmRAgent"`
	TokenRewards   []string `json:"token_rewards" bson:"token_rewards"`
	IsStaked       bool     `json:"is_staked" bson:"is_staked"`
	MyStakeBalance int64    `json:"my_stake_balance" bson:"my_stake_balance"`

	IsVerified bool `json:"is_verified" bson:"is_verified"`

	UniqueName string `json:"unique_name" bson:"unique_name"`

	IsHighLight bool `json:"is_high_light" bson:"is_high_light"`

	Treasury float64 `json:"treasury" bson:"treasury"`
}

/*
	{
	            "tokenAddress": "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
	            "amount": "0",
	            "stakingAddress": "0xc4bf4c24333124f4e3494da6198de7b0d85cdd7b",
	            "resourceId": "67657593620ea414f0bfc242",
	            "resourceTokenSymbol": "$RX",
	            "rewardTokenName": "USDC"
	        },
*/
type RewardP struct {
	TokenAddress        string `json:"token_address" bson:"token_address"`
	Amount              int64  `json:"amount" bson:"amount"`
	StakingAddress      string `json:"staking_address" bson:"staking_address"`
	ResourceId          string `json:"resource_id" bson:"resource_id"`
	ResourceTokenSymbol string `json:"resource_token_symbol" bson:"resource_token_symbol"`
	RewardTokenName     string `json:"resource_token_name" bson:"resource_token_name"`
}

/*
	{
	    "_id" : ObjectId("67ada8840d793436faf95d40"),
	    "contractAddress" : "0xe6f84d9645e9b4b47ba014c3c4e9c05f28a2c790",
	    "projectId" : "67721c9afe88f0266465cbf6",
	    "resourceId" : "67657593620ea414f0bfc242",
	    "createdAt" : ISODate("2025-02-13T08:08:36.229+0000"),
	    "updatedAt" : ISODate("2025-02-13T08:08:36.229+0000"),
	    "__v" : NumberInt(0),
	    "debug_resourceId" : "67657593620ea414f0bfc242",
	    "resource" : [
	        {
	            "_id" : ObjectId("67657593620ea414f0bfc242"),
	            "name" : "$RX",
	            "description" : "Social rAgent, based on X account",
	            "tokenAddress" : "0x831e64d4518a3d546c6436d1b8e9a9c19adbc618",
	            "chainId" : NumberInt(84532),
	            "tokenType" : "ERC721",
	            "tokenSymbol" : "$RX",
	            "status" : "active",
	            "createdAt" : ISODate("2024-03-20T10:00:00.000+0000"),
	            "updatedAt" : ISODate("2024-03-20T10:00:00.000+0000"),
	            "type" : "RX",
	            "category" : "resource",
	            "claimerContractAddress" : "0x83a0c404f18aae1dcdab15a9bf03e2cc960574b6"
	        }
	    ]
	}
*/
type Resource struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	ContractAddress string             `json:"contract_address" bson:"contractAddress"`
	ProjectId       string             `json:"project_id" bson:"projectId"`
	ResourceId      string             `json:"resource_id" bson:"resourceId"`
	CreatedAt       *time.Time         `json:"created_at" bson:"createdAt"`
	UpdatedAt       *time.Time         `json:"updated_at" bson:"updatedAt"`
	DebugResourceId string             `json:"debug_resource_id" bson:"debug_resource_id"`
	Resource        []ResourceP        `json:"resource" bson:"resource"`
}
type ResourceP struct {
	ID                     primitive.ObjectID `json:"_id" bson:"_id"`
	Name                   string             `json:"name" bson:"name"`
	Description            string             `json:"description" bson:"description"`
	TokenAddress           string             `json:"token_address" bson:"tokenAddress"`
	ChainId                int                `json:"chain_id" bson:"chainId"`
	TokenType              string             `json:"token_type" bson:"tokenType"`
	TokenSymbol            string             `json:"token_symbol" bson:"tokenSymbol"`
	Status                 string             `json:"status" bson:"status"`
	CreatedAt              *time.Time         `json:"created_at" bson:"createdAt"`
	UpdatedAt              *time.Time         `json:"updated_at" bson:"updatedAt"`
	Type                   string             `json:"type" bson:"type"`
	Category               string             `json:"category" bson:"category"`
	ClaimerContractAddress string             `json:"claimer_contract_address" bson:"claimerContractAddress"`
}

type ProjectDetail struct {
	ID              string     `json:"_id" bson:"_id"`
	Name            string     `json:"name" bson:"name"`
	Type            string     `json:"type" bson:"type"`
	Description     string     `json:"description" bson:"description"`
	Logo            string     `json:"logo" bson:"logo"`
	Website         string     `json:"website" bson:"website"`
	XUsername       string     `json:"x_username" bson:"xUsername"`
	DiscordUsername string     `json:"discord_username" bson:"discordUsername"`
	TokenAddress    string     `json:"token_address" bson:"tokenAddress"`
	TokenType       string     `json:"token_type" bson:"tokenType"`
	TokenDecimals   int        `json:"token_decimals" bson:"tokenDecimals"`
	TokenSymbol     string     `json:"token_symbol" bson:"tokenSymbol"`
	TokenLogo       string     `json:"token_logo" bson:"token_logo"`
	ChainId         int        `json:"chain_id" bson:"chainId"`
	CategoryId      string     `json:"category_id" bson:"categoryId"`
	Status          string     `json:"status" bson:"status"`
	IsVerified      bool       `json:"is_verified" bson:"is_verified"`
	CreatedTime     *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedTime     *time.Time `json:"updated_at" bson:"updatedAt"`
	//
	MarketCap        float64           `json:"market_cap" bson:"market_cap"`
	SwarmRAgent      int64             `json:"swarm_ragent" bson:"swarmRAgent"`
	SwarmMember      int               `json:"swarm_member" bson:"swarm_member"`
	SwarmReach       int               `json:"swarm_reach" bson:"swarm_reach"`
	SwarmLevel       string            `json:"swarm_level" bson:"swarm_level"`
	ResourceStakings []ResourceMapping `json:"resource_stakings" bson:"resource_stakings"`
	Rewards          []RewardP         `json:"rewards" bson:"rewards"`
}
type TokenInfo struct {
	Symbol   string `json:"symbol" bson:"symbol"`
	Name     string `json:"name" bson:"name"`
	Image    string `json:"image" bson:"image"`
	Decimals int    `json:"decimals" bson:"decimals"`
}
type ResourceMapping struct {
	TokenName             string                  `json:"token_name" bson:"token_name"`
	ResourceId            string                  `json:"resource_id" bson:"resource_id"`
	ResourceTokenAddress  string                  `json:"resource_token_address" bson:"resource_token_address"`
	ResourceTokenSymbol   string                  `json:"resource_token_symbol" bson:"resource_token_symbol"`
	ResourceTokenDecimals int                     `json:"resource_token_decimals" bson:"resource_token_decimals"`
	StakingAddress        string                  `json:"staking_address" bson:"staking_address"`
	MyStaking             int64                   `json:"my_staking" bson:"my_staking"`
	PendingRewards        []nodeDto.PendingReward `json:"pending_rewards" bson:"pending_rewards"`
	TotalStaked           int64                   `json:"total_staked" bson:"total_staked"`
	TokenRewards          []string                `json:"token_rewards" bson:"token_rewards"`
	Index                 int                     `json:"index" bson:"index"` //sort
}
type PendingReward struct {
	TokenAddress        string `json:"token_address" bson:"token_address"`
	Amount              int64  `json:"amount" bson:"amount"`
	StakingAddress      string `json:"staking_address" bson:"staking_address"`
	ResourceId          string `json:"resource_id" bson:"resource_id"`
	ResourceTokenSymbol string `json:"resource_token_symbol" bson:"resource_token_symbol"`
}
type ClaimReward struct {
	RewardTokenAddress string `json:"reward_token_address"`
	MyReward           string `json:"my_reward"`
	ProjectName        string `json:"project_name"`
	ProjectLogo        string `json:"project_logo"`
	StakingAddress     string `json:"staking_address"`
	MonthlyReward      string `json:"monthly_reward"`
	APY                string `json:"apy"`
	MyStakingBalance   int64  `json:"my_staking_balance"`
}
type RewardTokenInfo struct {
	TokenAddress   string   `json:"token_address"`
	TokenName      string   `json:"token_name"`
	TokenDecimals  int      `json:"token_decimals"`
	TokenLogo      string   `json:"token_logo"`
	WeeklyReward   string   `json:"weekly_reward"`
	TotalReward    string   `json:"total_reward"`
	Description    string   `json:"description"`
	ResourceTokens []string `json:"resource_tokens"`
}

type RewardToken struct {
	RewardAddress   string `json:"reward_address" bson:"reward_address"`
	MinRewardAmount int64  `json:"min_reward_amount" bson:"min_reward_amount"`
	Decimals        int    `json:"decimals" bson:"decimals"`
	Name            string `json:"name" bson:"name"`
}

type ProjectData struct {
	Name         string      `json:"name" bson:"name"`
	Logo         string      `json:"logo" bson:"logo"`
	Type         ProjectType `json:"type" bson:"type"`
	Description  string      `json:"description" bson:"description"`
	TokenAddress string      `json:"token_address,otmitempty" bson:"token_address"`
	Website      string      `json:"website" bson:"website"`
	XAccount     string      `json:"x_account" bson:"x_account"`
	Creator      string      `json:"creator" bson:"creator"`
}

type ProjectResponse struct {
	ProjectId      string   `json:"project_id" bson:"project_id"`
	TokenAddresses []string `json:"token_addresses" bson:"token_addresses"`
	Types          []string `json:"types" bson:"types"`
	StartBlock     int      `json:"start_block" bson:"start_block"`
	EndBlock       int      `json:"end_block" bson:"end_block"`
	UniqueName     string   `json:"unique_name" bson:"unique_name"`
}

type ProjectType string

const (
	AGENT ProjectType = "Agents"
	DEPIN ProjectType = "DePIN"
	SWARM ProjectType = "Swarm"
)

type APIResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AvailableProject struct {
	IsAvailable bool `json:"is_available"`
}

type ProjectStatus struct {
	ProcessStatus string `json:"process_status"`
	Status        string `json:"status"`
}

type SignCreateProjectResponse struct {
	Signature    string `json:"signature"`
	TotalFee     int64  `json:"total_fee"`
	FeeTokenType int    `json:"fee_token_type"`
	Nonce        int64  `json:"nonce"`
	EncodedName  string `json:"encoded_name"`
	Creator      string `json:"creator"`
	Factory      string `json:"factory"`
	TokenAddress string `json:"token_address"`
}

type SignCreateProjectRequest struct {
	ProjectName  string `json:"project_name"`
	FeeTokenType int    `json:"fee_token_type"`
}

type GetRizPriceResponse struct {
	RizPrice float64 `json:"riz_price"`
}

type CreateProjectResponse struct {
	UniqueName string `json:"unique_name"`
}
type XUserProject struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	XUserId      string             `json:"xuser_id" bson:"xuser_id"`
	XId          string             `json:"x_id" bson:"x_id"`
	ProjectId    string             `json:"project_id" bson:"project_id"`
	ClientId     string             `json:"client_id" bson:"client_id"`
	ClientSecret string             `json:"client_secret" bson:"client_secret"`
}

type ProjectHighlight struct {
	ProjectId string `json:"project_id" bson:"project_id"`
	UserId    string `json:"user_id" bson:"user_id"`
	Week      string `json:"week" bson:"week"`
}

type Session2 struct {
	ID    string    `json:"_id" bson:"_id"`
	Start time.Time `json:"start" bson:"start"`
	End   time.Time `json:"end" bson:"end"`
}
