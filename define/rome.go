package define

type TokenType string

const (
	ERC20   TokenType = "ERC20"
	ERC721  TokenType = "ERC721"
	ERC1155 TokenType = "ERC1155"
)

type ChainId int

const (
	BASE         ChainId = 8453
	BASE_SEPOLIA ChainId = 84532
	ETHEREUM     ChainId = 1
	POLYGON      ChainId = 137
	ARBITRUM     ChainId = 42161
	AVAX         ChainId = 43114
	BSC          ChainId = 56
	SOLANA       ChainId = 999999999
)

type ProjectStatus string

const (
	ACTIVE   ProjectStatus = "active"
	INACTIVE ProjectStatus = "inactive"
	PENDING  ProjectStatus = "pending"
	REJECTED ProjectStatus = "rejected"
)

const (
	RESOURCE_ACTIVE   = "active"
	RESOURCE_INACTIVE = "inactive"
	RESOURCE_PENDING  = "pending"
	RESOURCE_REJECTED = "rejected"
)

const (
	TX_TYPE_MINT                   = "mint"
	TX_TYPE_TRANSFER               = "transfer"
	TX_TYPE_BURN                   = "burn"
	TX_TYPE_APPROVE                = "approve"
	TX_TYPE_LENDING                = "lending"
	TX_TYPE_BORROW                 = "borrow"
	TX_TYPE_STAKE                  = "stake"
	TX_TYPE_BUY_RAGENT_MARKETPLACE = "buyRagentMarketplace"
	TX_CREATE_BID                  = "createBidMarketplace"
	TX_RELEASE_BID                 = "releaseBidMarketplace"
	TX_AGREE_BID                   = "agreeBidMarketplace"
	TX_REMOVE_BID                  = "removeBidMarketplace"
)

const (
	TX_PENDING = "pending"
	TX_SUCCESS = "success"
	TX_FAILED  = "failed"
)

const (
	RESOURCE_TYPE_RC  = "RC"
	RESOURCE_TYPE_RX  = "RX"
	RESOURCE_TYPE_RD  = "RD"
	RESOURCE_TYPE_RE  = "RE"
	RESOURCE_TYPE_BRC = "BRC"
	RESOURCE_TYPE_BRX = "BRX"
	RESOURCE_TYPE_BRD = "BRD"
	RESOURCE_TYPE_BRE = "BRE"
)
const (
	CATEGORY_RESOURCE = "resource"
	CATEGORY_BORROW   = "borrow"
)

type MetadataType string

const (
	PAG MetadataType = "pag"
	RCM MetadataType = "rc"
	RXM MetadataType = "rx"
)

const (
	PROJECT_AGENT = "Agents"
	PROJECT_DEPIN = "DePIN"
	PROJECT_SWARM = "Swarm"
)

type SwarmLevel string

const (
	LEGION SwarmLevel = "LEGION"
)

const (
	LOG_MINT_RESOURCE     = "mintResource"
	LOG_BORROW_RESOURCE   = "borrowResource"
	LOG_REPAY_RESOURCE    = "repayResource"
	LOG_REPAY_BORROW      = "repayBorrow"
	LOG_TRANSFER_RESOURCE = "transferResource"
	LOG_TRANSFER_BORROW   = "transferBorrow"
	LOG_BURN_RESOURCE     = "burnResource"
	LOG_BURN_BORROW       = "burnBorrow"
	LOG_STAKE_RESOURCE    = "stakeResource"
	LOG_UNSTAKE_RESOURCE  = "unstakeResource"
	LOG_LEND_RESOURCE     = "lendResource"
	LOG_UNLEND_RESOURCE   = "unlendResource"
	LOG_X_CONNECT         = "xConnect"
	LOG_MINT_PAG          = "mintPAG"
)

const (
	REWARD_TYPE_MINT_RC             = "mintRC"
	REWARD_TYPE_MINT_RX             = "mintRX"
	REWARD_TYPE_MINT_RD             = "mintRD"
	REWARD_TYPE_MINT_RE             = "mintRE"
	REWARD_TYPE_CREATE_PAG          = "createPAG"
	REWARD_TYPE_BORROW_ONE_MONTH_RC = "borrowOneMonthRC"
	REWARD_TYPE_BORROW_ONE_MONTH_RD = "borrowOneMonthRD"
	REWARD_TYPE_BORROW_ONE_MONTH_RE = "borrowOneMonthRE"
	REWARD_TYPE_BORROW_RX_SESSION2  = "borrowRX"
	REWARD_TYPE_BORROW_RC_SESSION2  = "borrowRC"
	REWARD_TYPE_BORROW_RD_SESSION2  = "borrowRD"
	REWARD_TYPE_BORROW_RE_SESSION2  = "borrowRE"
)

const (
	REPU_TYPE_IDENTITY_WORLD_ID       = "identityWorldId"
	REPU_TYPE_IDENTITY_DISCORD        = "identityDiscord"
	REPU_TYPE_IDENTITY_ACTIVE_WALLET  = "identityActiveWallet"
	REPU_TYPE_IDENTITY_SELECT_COUNTRY = "identitySelectCountry"
	REPU_TYPE_ACTION_MINT_RAGENT      = "actionMintRAgent"
	REPU_TYPE_ACTION_5D_ACTIVE_RX     = "action5dActiveRX"
	REPU_TYPE_ACTION_5D_ACTIVE_RC     = "action5dActiveRC"
	REPU_TYPE_ACTION_5D_ACTIVE_RD     = "action5dActiveRD"
	REPU_TYPE_ACTION_5D_ACTIVE_RE     = "action5dActiveRE"
	REPU_TYPE_LEND_ANY_RAGENT         = "lendAnyRAgent"
	REPU_TYPE_STAKED_ANY_RAGENT       = "stakedAnyRAgent"
	REPU_TYPE_STAKED_ANY_TWO_RAGENT   = "stakedAnyTwoRAgent"
	REPU_TYPE_BORROW_RC               = "borrowRC"
	REPU_TYPE_BORROW_RD               = "borrowRD"
	REPU_TYPE_BORROW_RE               = "borrowRE"
	REPU_TYPE_PAG_MINT                = "pagMint"
	REPU_TYPE_ZNODE_OWNER             = "znodeOwner"
)

type ReputationCategory string

const (
	IDENTITY ReputationCategory = "identity"
	ACTION   ReputationCategory = "action"
)
const (
	MINT_STATUS_MINTABLE     = "mintable"
	MINT_STATUS_MINTED       = "minted"
	MINT_STATUS_NOT_ELIGIBLE = "not_eligible"
)

const (
	PROJECT_PROGRESS_STATUS_PENDING  = "pending"
	PROJECT_PROGRESS_STATUS_VERIFIED = "verified"
)
const (
	SESSION2_STAKE_PROJECT_LOWEST_REWARD  = "stakeProjectLowestReward"
	SESSION2_STAKE_PROJECT_INCREASE_LEVEL = "stakeProjectIncreaseLevel"
	SESSION2_STAKE_FRIST_USER_PROJECT     = "firstUserStakeProject"
	SESSION2_STAKE_TEN_PROJECT            = "stakeTenProject"
	SESSION2_STAKE_HIGHLIGHT_PROJECT      = "stakeHighlightProject"
	SESSION2_STAKE_FIRST_24h_REWARD       = "stakeFirst24hReward"
	SESSION2_STAKE_LAST_24h_REWARD        = "stakeLast24hReward"
	//
	SESSION2_MKP_BUY_ANY_AGENT_FRIST_MONTH        = "mkpBuyAnyAgentFristMonth"       //not use
	SESSION2_MKP_BID_ACCEPT_ANY_AGENT_FIRST_MONTH = "mkpBidAcceptAnyAgentFirstMonth" //not use
	//many time
	SESSION2_MKP_BIDDER                 = "mkpBidder" //not use
	SESSION2_MKP_BIDDER_WAS_ACCEPT      = "mkpBidderWasAccept"
	SESSION2_MKP_ACCEPT_BID             = "mkpAcceptBid"
	SESSION2_MKP_BUY_AGENT              = "mkpBuyAgent"
	SESSION2_MKP_SELL_AGENT             = "mkpSellAgent"
	SESSION2_MKP_BUY_AGENT_HIGHLIGHT    = "mkpBuyAgentHighlight" //not use
	SESSION2_MKP_HOLD_AGENT_EVERY_2WEEK = "mkpHoldAgentEvery2Week"
)
