package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignBuyRagentMarketResponse struct {
	Signature       string `json:"signature"`
	Nonce           int64  `json:"nonce"`
	WalletAddress   string `json:"wallet_address"`
	Quantity        int64  `json:"quantity"`
	TotalPrice      string `json:"total_price"`
	TrackingId      string `json:"tracking_id"`
	TokenAddress    string `json:"token_address"`
	RagentToken     string `json:"ragent_token"`
	ContractAddress string `json:"contract_address"`
}

type SwapResponse struct {
	DstAmount string `json:"dstAmount"`
	Tx        struct {
		Data     string `json:"data"`
		GasPrice string `json:"gasPrice"`
		Gas      int    `json:"gas"`
		Value    string `json:"value"`
		From     string `json:"from"`
		To       string `json:"to"`
	} `json:"tx"`
}

type MkpRagentUserBuy struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Side               string             `json:"-" bson:"side,omitempty"`
	TxHash             string             `json:"tx_hash" bson:"tx_hash,omitempty"`
	UserId             string             `json:"-" bson:"user_id,omitempty"`
	WalletAddress      string             `json:"wallet_address" bson:"wallet_address,omitempty"`
	Amount             string             `json:"-" bson:"amount,omitempty"`
	TotalPrice         string             `json:"-" bson:"total_price,omitempty"`
	MkpRagentId        string             `json:"mkp_ragent_id" bson:"mkp_ragent_id,omitempty"`
	Status             string             `json:"-" bson:"status,omitempty"`
	AccountList        map[string]string  `json:"-" bson:"account_list,omitempty"`
	AccountListBase    map[string]string  `json:"-" bson:"account_list_base,omitempty"`
	RemovedTokens      map[string]string  `json:"-" bson:"removed_tokens,omitempty"`
	RemovedTokensBase  map[string]string  `json:"-" bson:"removed_tokens_base,omitempty"`
	Signature          string             `json:"-" bson:"signature,omitempty"`
	Nonce              int64              `json:"-" bson:"nonce,omitempty"`
	ActiveBid          int                `json:"active_bid,omitempty"`
	RAgentTokenAddress string             `json:"ragent_token" bson:"-"`

	BuySwapTokenTxHash       string            `json:"-" bson:"buy_swap_token_tx_hash,omitempty"`
	SellSwapTokenTxHash      map[string]string `json:"-" bson:"sell_swap_token_tx_hash,omitempty"`
	SellAmountToken          map[string]string `json:"-" bson:"sell_amount_token,omitempty"`
	SellTransferToUserTxHash string            `json:"-" bson:"sell_transfer_to_user_tx_hash,omitempty"`
	TransferTokenTxHash      map[string]string `json:"-" bson:"transfer_token_tx_hash,omitempty"` // save tx hash of transfer token to each account

	SwapProjectTokenHash string `json:"swap_project_token_hash" bson:"swap_project_token_hash,omitempty"`
	SwapRizTokenHash     string `json:"swap_riz_token_hash" bson:"swap_riz_token_hash,omitempty"`

	TransferProjectTokenTxHash string `json:"transfer_project_token_tx_hash" bson:"transfer_project_token_tx_hash,omitempty"`
	TransferRizTokenTxHash     string `json:"transfer_riz_token_tx_hash" bson:"transfer_riz_token_tx_hash,omitempty"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`

	TranferTokenStartTime *time.Time `json:"-" bson:"transfer_token_start_time,omitempty"`

	SwapProjectTokenProcessByWallet string `json:"-" bson:"swap_project_token_process_by_wallet,omitempty"`
	SwapProjectTokenToWallet        string `json:"-" bson:"swap_project_token_to_wallet,omitempty"`
	SwapRizTokenProcessByWallet     string `json:"-" bson:"swap_riz_token_process_by_wallet,omitempty"`
	ProcessByWallet                 string `json:"-" bson:"process_by_wallet,omitempty"`

	TransferUserProjectTokenTxHash    string            `json:"transfer_user_project_token_tx_hash" bson:"transfer_user_project_token_tx_hash,omitempty"`
	TransferUserProjectTokenTxHashMap map[string]string `json:"-" bson:"transfer_user_project_token_tx_hash_map,omitempty"`
	TransferUserRizTokenTxHash        string            `json:"transfer_user_riz_token_tx_hash" bson:"transfer_user_riz_token_tx_hash,omitempty"`

	BatchUnwrapTxHash string     `json:"batch_unwrap_tx_hash" bson:"batch_unwrap_tx_hash,omitempty"`
	BatchUnwrapTime   *time.Time `json:"batch_unwrap_time" bson:"batch_unwrap_time,omitempty"`

	BalanceToken map[string]BalanceTokenRagent `json:"-" bson:"balance_token,omitempty"`

	WalletBase string `json:"wallet_base" bson:"wallet_base,omitempty"`

	TotalFee float64 `json:"total_fee_usdc" bson:"total_fee_usdc,omitempty"`
}

type MkpRagentUserBuyResponse struct {
	MarketCap     float64 `json:"market_cap" bson:"market_cap"`
	AgentType     string  `json:"agent_type" bson:"agent_type"`
	SellPrice     float64 `json:"sell_price" bson:"sell_price"`
	Sell24hChange float64 `json:"sell_24h_change" bson:"sell_24h_change"`
	Quantity      int64   `json:"quantity" bson:"quantity"`
	X             string  `json:"x" bson:"x"`
	Coingecko     string  `json:"coingecko" bson:"coingecko"`
	CookieDAO     string  `json:"cookie_dao" bson:"cookie_dao"`
	DEXScreener   string  `json:"dexscreener" bson:"dexscreener"`
	USDValue      float64 `json:"usd_value" bson:"usd_value"`

	RagentName  string `json:"ragent_name" bson:"ragent_name"`
	Company     string `json:"company" bson:"company"`
	CompanyLogo string `json:"company_logo" bson:"company_logo"`
	TokenLogo   string `json:"token_logo" bson:"token_logo"`

	ActiveBid     int64  `json:"active_bid" bson:"active_bid"`
	WalletAddress string `json:"wallet_address" bson:"wallet_address"`
	MkpRagentId   string `json:"mkp_ragent_id" bson:"mkp_ragent_id"`
	Side          string `json:"side,omitempty" bson:"-"`
	Status        string `json:"status,omitempty" bson:"-"`
	TrackingId    string `json:"tracking_id,omitempty" bson:"-"`
	TxHash        string `json:"tx_hash,omitempty" bson:"-"`

	ChainId int `json:"chain_id" bson:"chain_id"`
	APY     int `json:"apy" bson:"-"`
}

type MkpRagent struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	Name               string             `json:"name" bson:"name"`
	Company            string             `json:"company" bson:"company"`
	Detail             string             `json:"detail" bson:"detail"`
	Logo               string             `json:"logo" bson:"logo"`
	TokenAddress       string             `json:"token_address" bson:"token_address"`
	TokenType          string             `json:"token_type" bson:"token_type"`
	TokenDecimals      int                `json:"token_decimals" bson:"token_decimals"`
	TokenSymbol        string             `json:"token_symbol" bson:"token_symbol"`
	TokenLogo          string             `json:"token_logo" bson:"token_logo"`
	ChainId            int                `json:"chain_id" bson:"chain_id"`
	Status             string             `json:"status" bson:"status"`
	CreatedAt          time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt          time.Time          `json:"-" bson:"updatedAt"`
	Creator            string             `json:"creator" bson:"creator"`
	UniqueName         string             `json:"-" bson:"unique_name"`
	Holder             int                `json:"-" bson:"holder"`
	ActiveBid          int                `json:"active_bid" bson:"active_bid"`
	ProjectToken       float64            `json:"project_token" bson:"project_token"`
	RizToken           float64            `json:"riz_token" bson:"riz_token"`
	X                  string             `json:"x" bson:"x"`
	Coingecko          string             `json:"coingecko" bson:"coingecko"`
	CookieDAO          string             `json:"cookie_dao" bson:"cookie_dao"`
	DEXScreener        string             `json:"dexscreener" bson:"dexscreener"`
	RagentTokenAddress string             `json:"ragent_token_address" bson:"ragent_token_address"` // nft_contract_address
	IsHighLight        bool               `json:"is_high_light" bson:"is_high_light"`

	BuyPrice      float64 `json:"buy_price" bson:"-"`
	SellPrice     float64 `json:"sell_price" bson:"-"`
	Buy24hChange  float64 `json:"buy_24h_change" bson:"-"`
	Sell24hChange float64 `json:"sell_24h_change" bson:"-"`
	MarketCap     float64 `json:"market_cap" bson:"-"`
	APY           int     `json:"apy" bson:"-"`
	AgentType     string  `json:"agent_type" bson:"-"`
	IsMyAgentOnly bool    `json:"is_my_agent_only" bson:"-"`
}

type APIResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type APIResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TokenPriceResponse struct {
	Price float64 `json:"price"`
}

type BuyStatusResponse struct {
	Status string `json:"status"`
}

type SwarmOwnerResponse struct {
	IsSwarmOwner bool `json:"is_swarm_owner"`
}

type GetQuantityTokenResponse struct {
	RAgentEstValue float64 `json:"ragent_est_value"`
	Slippage       float64 `json:"slippage"`
	TotalPrice     float64 `json:"total_price"`
	TokenIds       []int64 `json:"token_ids"`
}

type GetTotalValueByWalletResponse struct {
	TotalValue float64 `json:"total_value"`
}

type SignBuyRagentMarketRequest struct {
	Quantity int64 `json:"quantity"`
}

type VerifyBuyTransactionRequest struct {
	Quantity int64 `json:"quantity"`
	Nonce    int64 `json:"nonce"`
}

type MkpSetting struct {
	BuyFee       float64 `json:"buy_fee" bson:"buy_fee"`
	SellFee      float64 `json:"sell_fee" bson:"sell_fee"`
	UnwrapFee    float64 `json:"unwrap_fee" bson:"unwrap_fee"`
	BidFee       float64 `json:"bid_fee" bson:"bid_fee"`
	SlippageRate float64 `json:"slippage_rate" bson:"slippage_rate"`
	SwapFee      float64 `json:"swap_fee" bson:"swap_fee"`
}

type RouterResult struct {
	ChainId       string `json:"chainId"`
	ChainIndex    string `json:"chainIndex"`
	DexRouterList []struct {
		Router        string `json:"router"`
		RouterPercent string `json:"routerPercent"`
		SubRouterList []struct {
			DexProtocol []struct {
				DexName string `json:"dexName"`
				Percent string `json:"percent"`
			} `json:"dexProtocol"`
			FromToken struct {
				Decimal              string `json:"decimal"`
				IsHoneyPot           bool   `json:"isHoneyPot"`
				TaxRate              string `json:"taxRate"`
				TokenContractAddress string `json:"tokenContractAddress"`
				TokenSymbol          string `json:"tokenSymbol"`
				TokenUnitPrice       string `json:"tokenUnitPrice"`
			} `json:"fromToken"`
			ToToken struct {
				Decimal              string `json:"decimal"`
				IsHoneyPot           bool   `json:"isHoneyPot"`
				TaxRate              string `json:"taxRate"`
				TokenContractAddress string `json:"tokenContractAddress"`
				TokenSymbol          string `json:"tokenSymbol"`
				TokenUnitPrice       string `json:"tokenUnitPrice"`
			} `json:"toToken"`
		} `json:"subRouterList"`
	} `json:"dexRouterList"`
	EstimateGasFee string `json:"estimateGasFee"`
	FromToken      struct {
		Decimal              string `json:"decimal"`
		IsHoneyPot           bool   `json:"isHoneyPot"`
		TaxRate              string `json:"taxRate"`
		TokenContractAddress string `json:"tokenContractAddress"`
		TokenSymbol          string `json:"tokenSymbol"`
		TokenUnitPrice       string `json:"tokenUnitPrice"`
	} `json:"fromToken"`
	FromTokenAmount       string `json:"fromTokenAmount"`
	PriceImpactPercentage string `json:"priceImpactPercentage"`
	QuoteCompareList      []struct {
		AmountOut string `json:"amountOut"`
		DexLogo   string `json:"dexLogo"`
		DexName   string `json:"dexName"`
		TradeFee  string `json:"tradeFee"`
	} `json:"quoteCompareList"`
	ToToken struct {
		Decimal              string `json:"decimal"`
		IsHoneyPot           bool   `json:"isHoneyPot"`
		TaxRate              string `json:"taxRate"`
		TokenContractAddress string `json:"tokenContractAddress"`
		TokenSymbol          string `json:"tokenSymbol"`
		TokenUnitPrice       string `json:"tokenUnitPrice"`
	} `json:"toToken"`
	ToTokenAmount string `json:"toTokenAmount"`
	TradeFee      string `json:"tradeFee"`
}
type OkxSwapResponse struct {
	RouterResult RouterResult `json:"routerResult"`
	Tx           OkxTx        `json:"tx"`
}

type OkxTx struct {
	Data                 string   `json:"data"`
	From                 string   `json:"from"`
	Gas                  string   `json:"gas"`
	GasPrice             string   `json:"gasPrice"`
	Value                string   `json:"value"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
	MinReceiveAmount     string   `json:"minReceiveAmount"`
	SignatureData        []string `json:"signatureData"`
	To                   string   `json:"to"`
	Slippage             string   `json:"slippage"`
}

type OkxApproveResponse struct {
	Data               string `json:"data"`
	GasLimit           string `json:"gasLimit"`
	GasPrice           string `json:"gasPrice"`
	DexContractAddress string `json:"dexContractAddress"`
}

type RagentDetailResponse struct {
	Balance     int64   `json:"balance"`
	CompanyName string  `json:"company_name"`
	CompanyLogo string  `json:"company_logo"`
	TokenLogo   string  `json:"token_logo"`
	TokenSymbol string  `json:"token_symbol"`
	AmountToken float64 `json:"amount_token"`
	AmountRiz   float64 `json:"amount_riz"`
}

type SellRagentRequest struct {
	// TokenIds []int64 `json:"token_ids"`
	Hash string `json:"hash"`
}

type GetMyAgentResponse struct {
	TotalValue float64                     `json:"total_value"`
	Data       []*MkpRagentUserBuyResponse `json:"data"`
}

type UnwrapRagentRequest struct {
	// TokenIds []int64 `json:"token_ids"`
	Hash string `json:"hash"`
}

type UnwrapRagentResponse struct {
	TrackingId string `json:"tracking_id"`
}

type SellRagentResponse struct {
	TrackingId string `json:"tracking_id"`
}

//	{
//	    "wallet": "",
//	    "ragent_id": "",
//	    "account_list":"",
//	    "quantity": 0
//	}
type MkpUser struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Wallet          string             `json:"wallet" bson:"wallet"`
	SolanaWallet    string             `json:"solana_wallet" bson:"solana_wallet"`
	RagentId        string             `json:"ragent_id" bson:"ragent_id"`
	AccountList     map[string]string  `json:"account_list" bson:"account_list"`
	AccountListBase map[string]string  `json:"account_list_base" bson:"account_list_base"`
	Quantity        int64              `json:"quantity" bson:"quantity"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type RagentSellUnwrapInfoResponse struct {
	NFTContractAddress string  `json:"nft_contract_address"`
	TokenIds           []int64 `json:"token_ids"`
	WalletKeeper       string  `json:"wallet_keeper"`
	// TrackingId         string  `json:"tracking_id"`
}

type MkpRagentEstPriceCache struct {
	TokenProjectPrice float64 `json:"token_project_price"`
	RizTokenPrice     float64 `json:"riz_token_price"`
	RagentEstPrice    float64 `json:"ragent_est_price"`
	PriceImpact       float64 `json:"price_impact"`
}

type Bid struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	RagentId     string `json:"ragent_id" bson:"ragent_id,omitempty"`
	BidOnchainId string `json:"bid_onchain_id" bson:"bid_onchain_id,omitempty"`

	BidderWallet string  `json:"bidder_wallet" bson:"bidder_wallet,omitempty"`
	BidAmount    int64   `json:"bid_amount" bson:"bid_amount,omitempty"`
	BidPrice     float64 `json:"bid_price" bson:"bid_price,omitempty"`
	IsPremium    bool    `json:"is_premium" bson:"is_premium,omitempty"`
	FDV          float64 `json:"fdv" bson:"fdv,omitempty"`

	PaymentToken         string `json:"payment_token" bson:"payment_token,omitempty"`
	PaymentTokenDecimals int    `json:"payment_token_decimals" bson:"payment_token_decimals,omitempty"`

	VestingType int64 `json:"vesting_type" bson:"vesting_type,omitempty"`

	Status string `json:"status" bson:"status,omitempty"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`

	CreatedBidTxHash string `json:"created_bid_tx_hash" bson:"created_bid_tx_hash,omitempty"`
	RemovedBidTxHash string `json:"removed_bid_tx_hash" bson:"removed_bid_tx_hash,omitempty"`
}

type UserBid struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	WalletAddress string `json:"wallet_address" bson:"wallet_address,omitempty"`
	BidId         string `json:"bid_id" bson:"bid_id,omitempty"`

	BidAmount int64   `json:"bid_amount" bson:"bid_amount,omitempty"`
	BidPrice  float64 `json:"bid_price" bson:"bid_price,omitempty"`

	VestingType int64     `json:"vesting_type" bson:"vesting_type,omitempty"`
	VestingTime time.Time `json:"vesting_time" bson:"vesting_time,omitempty"`

	Status string `json:"status" bson:"status,omitempty"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`

	IsPremium bool `json:"is_premium" bson:"-"`

	BidOnchainId string `json:"bid_onchain_id" bson:"bid_onchain_id,omitempty"`

	ReleaseBidTxHash string `json:"release_bid_tx_hash" bson:"release_bid_tx_hash,omitempty"`
	AgreeBidTxHash   string `json:"agree_bid_tx_hash" bson:"agree_bid_tx_hash,omitempty"`
	AgreeBidWallet   string `json:"accept_bid_wallet_address" bson:"accept_bid_wallet_address,omitempty"`
}

type UnwrapTokenQuantityResponse struct {
	AmountProjectToken float64 `json:"amount_project_token"`
	AmountRiz          float64 `json:"amount_riz"`
	TokenIds           []int64 `json:"token_ids"`
}

type BidResponse struct {
	ID string `json:"id"`
	BidderInfo
	RagentsToSell  int64   `json:"ragents_to_sell"`
	EstValue       float64 `json:"est_value"`       // Numeric value for the "$0000" part in "Est Value / Premium"
	PremiumPercent float64 `json:"premium_percent"` // Numeric value for the "+000%" part in "Est Value / Premium"
	FDV            float64 `json:"fdv"`             // Numeric value for the "$00000" FDV part
	Vesting        int64   `json:"vesting"`         // Vesting period in days (1, 7, 30, 0 for no vesting)
	BidId          string  `json:"bid_id"`
	IsOwner        bool    `json:"is_owner"`
	RagentId       string  `json:"ragent_id"`
	MC             float64 `json:"mc"`
	TokenAmount    int64   `json:"token_amount"`
}

type VestingTokenResponse struct {
	BidderInfo
	BidId          string    `json:"bid_id"`
	RagentsToSell  int64     `json:"ragents_to_sell"` // Number of rAgents to sell (0000)
	EstValue       float64   `json:"est_value"`       // Estimated value in dollars
	PremiumPercent float64   `json:"premium_percent"` // Premium percentage (+000%)
	FDV            float64   `json:"fdv"`             // Fully Diluted Value
	VestingStatus  string    `json:"vesting_status"`  // "Completed" or remaining time countdown
	VestingEndsAt  time.Time `json:"vesting_ends_at"` // Unix timestamp when vesting ends (for countdown)
	IsClaimable    bool      `json:"is_claimable"`    // Whether token is claimable
	RagentId       string    `json:"ragent_id"`
}

type MyActiveBidResponse struct {
	BidderInfo
	RagentsToSell   int64     `json:"ragents_to_sell"`    // Number of rAgents to sell
	EstValue        float64   `json:"est_value"`          // Estimated value in dollars
	FullBidEstValue float64   `json:"full_bid_est_value"` // Full bid est value
	PremiumPercent  float64   `json:"premium_percent"`    // Premium percentage (+000%)
	FDV             float64   `json:"fdv"`                // Fully Diluted Value
	VestingPeriod   int64     `json:"vesting"`            // Vesting period in days (1, 7, 30, 0 for no vesting)
	BidId           string    `json:"bid_id"`
	Status          string    `json:"status"`     // Status of the bid
	CreatedAt       time.Time `json:"created_at"` // Unix timestamp of when bid was created
	RagentId        string    `json:"ragent_id"`
}

type BidHistoryResponse struct {
	BidderInfo
	RagentsToSell int64 `json:"ragents_to_sell"` // Number of rAgents to sell
	// EstValue       float64   `json:"est_value"`       // Estimated value in dollars
	TotalValue     float64   `json:"total_value"`     // Total value in dollars
	PremiumPercent float64   `json:"premium_percent"` // Premium percentage (+000%)
	FDV            float64   `json:"fdv"`             // Fully Diluted Value
	VestingPeriod  int64     `json:"vesting"`         // Vesting period in days (1, 7, 30, 0 for no vesting)
	Status         string    `json:"status"`          // "Accepted", "Rejected", "Expired", etc.
	ActionDate     time.Time `json:"action_date"`     // Unix timestamp of when the bid was accepted/rejected/expired
}

type BidderInfo struct {
	BidderName   string `json:"bidder_name"`   // Swarm name
	BidderLogo   string `json:"bidder_logo"`   // Swarm logo URL
	BidderToken  string `json:"bidder_token"`  // $SWARM token symbol
	BidderSymbol string `json:"bidder_symbol"` // token symbol
	BidderXLink  string `json:"bidder_x_link"` // X social link
	BidderType   string `json:"bidder_type"`   // "Swarm" or "User"
}

type AcceptBidQuantityResponse struct {
	AmountAutoAgreed int64   `json:"amount_auto_agreed"`
	TokenReceive     float64 `json:"token_receive"`
	TotalValue       float64 `json:"total_value"`

	NFTContractAddress string `json:"nft_contract_address"`

	// VestingTime time.Time `json:"vesting_time,omitempty"` // appear when this is vesting token
}

type MkpSwapWallet struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Wallet string             `json:"wallet" bson:"wallet"`
	Secret string             `json:"secret" bson:"secret"`
	Status string             `json:"status" bson:"status"`
}

// LiFi API structures
type LiFiContractCall struct {
	FromAmount         string `json:"fromAmount"`
	FromTokenAddress   string `json:"fromTokenAddress"`
	ToContractAddress  string `json:"toContractAddress"`
	ToContractCallData string `json:"toContractCallData"`
	ToContractGasLimit string `json:"toContractGasLimit"`
}

type LiFiQuoteRequest struct {
	ToAmount      string             `json:"toAmount"`
	FromAmount    string             `json:"fromAmount"`
	ContractCalls []LiFiContractCall `json:"contractCalls"`
	FromChain     int                `json:"fromChain"`
	FromToken     string             `json:"fromToken"`
	FromAddress   string             `json:"fromAddress"`
	ToChain       int                `json:"toChain"`
	ToToken       string             `json:"toToken"`
	ToAddress     string             `json:"toAddress"`
}

type LiFiQuoteResponse struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Tool        string `json:"tool"`
	ToolDetails struct {
		Key     string `json:"key"`
		Name    string `json:"name"`
		LogoURI string `json:"logoURI"`
	} `json:"toolDetails"`
	Action struct {
		FromChainId int `json:"fromChainId"`
		ToChainId   int `json:"toChainId"`
		FromToken   struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			LogoURI  string `json:"logoURI"`
		} `json:"fromToken"`
		ToToken struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			LogoURI  string `json:"logoURI"`
		} `json:"toToken"`
		FromAmount string `json:"fromAmount"`
		ToAmount   string `json:"toAmount"`
		Slippage   string `json:"slippage"`
	} `json:"action"`
	Estimate struct {
		Tool              string `json:"tool"`
		FromAmount        string `json:"fromAmount"`
		ToAmount          string `json:"toAmount"`
		ToAmountMin       string `json:"toAmountMin"`
		ApprovalAddress   string `json:"approvalAddress"`
		ExecutionDuration int    `json:"executionDuration"`
		FeeCosts          []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Token       struct {
				Address  string `json:"address"`
				ChainId  int    `json:"chainId"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				Name     string `json:"name"`
				LogoURI  string `json:"logoURI"`
			} `json:"token"`
			Amount     string `json:"amount"`
			AmountUSD  string `json:"amountUSD"`
			Percentage string `json:"percentage"`
			Included   bool   `json:"included"`
		} `json:"feeCosts"`
		GasCosts []struct {
			Type      string `json:"type"`
			Estimate  string `json:"estimate"`
			Limit     string `json:"limit"`
			Amount    string `json:"amount"`
			AmountUSD string `json:"amountUSD"`
			Token     struct {
				Address  string `json:"address"`
				ChainId  int    `json:"chainId"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				Name     string `json:"name"`
				LogoURI  string `json:"logoURI"`
			} `json:"token"`
		} `json:"gasCosts"`
	} `json:"estimate"`
	TransactionRequest struct {
		Data     string `json:"data"`
		To       string `json:"to"`
		Value    string `json:"value"`
		From     string `json:"from"`
		ChainId  int    `json:"chainId"`
		GasLimit string `json:"gasLimit"`
		GasPrice string `json:"gasPrice"`
	} `json:"transactionRequest"`
	IncludedSteps []struct {
		Action struct {
		} `json:"action"`
		Estimate struct {
			FromAmount      string `json:"fromAmount"`
			ApprovalAddress string `json:"approvalAddress"`
		} `json:"estimate"`
	} `json:"includedSteps"`
}

type BalanceTokenRagent struct {
	ProjectToken string
	RizToken     string
}

type LiFiStatusResponse struct {
	TransactionId string `json:"transactionId"`
	Sending       struct {
		TxHash string `json:"txHash"`
		TxLink string `json:"txLink"`
		Amount string `json:"amount"`
		Token  struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			LogoURI  string `json:"logoURI"`
		} `json:"token"`
		ChainId      int    `json:"chainId"`
		GasPrice     string `json:"gasPrice"`
		GasUsed      string `json:"gasUsed"`
		GasAmount    string `json:"gasAmount"`
		GasAmountUSD string `json:"gasAmountUSD"`
		AmountUSD    string `json:"amountUSD"`
		Value        string `json:"value"`
		Timestamp    int64  `json:"timestamp"`
	} `json:"sending"`
	Receiving struct {
		TxHash string `json:"txHash"`
		TxLink string `json:"txLink"`
		Amount string `json:"amount"`
		Token  struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			LogoURI  string `json:"logoURI"`
		} `json:"token"`
		ChainId      int    `json:"chainId"`
		GasPrice     string `json:"gasPrice"`
		GasUsed      string `json:"gasUsed"`
		GasAmount    string `json:"gasAmount"`
		GasAmountUSD string `json:"gasAmountUSD"`
		AmountUSD    string `json:"amountUSD"`
		Value        string `json:"value"`
		Timestamp    int64  `json:"timestamp"`
	} `json:"receiving"`
	LifiExplorerLink string `json:"lifiExplorerLink"`
	FromAddress      string `json:"fromAddress"`
	ToAddress        string `json:"toAddress"`
	Tool             string `json:"tool"`
	Status           string `json:"status"` // "DONE", "PENDING", "FAILED", etc.
	SubStatus        string `json:"substatus"`
	SubStatusMessage string `json:"substatusMessage"`
}

type SessionUserReward struct {
	WalletAddress string `bson:"wallet_address" json:"wallet_address"`
	Reward        int64  `bson:"reward" json:"reward"`
	Point         int64  `bson:"point" json:"point"`
}

type SessionLevel struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	Name             string             `bson:"name" json:"name"`
	PointLimit       int64              `bson:"point_limit" json:"point_limit"`
	Apy              int64              `bson:"apy" json:"apy"`
	Rebate           int64              `bson:"rebate" json:"rebate"`
	AdditionalReward string             `bson:"additional_reward" json:"additional_reward"`
	Reward           int64              `bson:"reward" json:"reward"`
	RewardDaily      int64              `bson:"reward_daily" json:"reward_daily"`
	NextLevel        string             `bson:"next_level" json:"next_level"`
	NextLevelPoint   int64              `bson:"next_level_point" json:"next_level_point"`
}
type UserLevel struct {
	Level        string        `bson:"level" json:"level"`
	SessionLevel *SessionLevel `bson:"session_level" json:"session_level"`
}
