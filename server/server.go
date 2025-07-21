package server

import (
	http_srv "github.com/Rivalz-ai/framework-be/framework/service/http"
	//"github.com/Rivalz-ai/framework-be/framework/log"
	"errors"

	"github.com/Rivalz-ai/framework-be/framework/utils"
	//"fmt"
)

type ExtendConfig struct {
	InternalSecret              string
	InternalSecretCreateProject string
	InternalSecretMonitoring    string
	XNumThread                  int
	XPageSize                   int
	XClientID                   string
	XClientSecret               string
	XSleepTime                  int
	JwtSecretKey                string
	JwtExpire                   int
	IsZnodeOwnerCheck           bool
	XURL                        string
	CoinGeckoURL                string
	CoinGeckoToken              string
	RNodeCheckURL               string
	ResourceTokenSignPrivateKey string
	StakingContractPrivateKey   string
	RAgentVersion               string
	//
	Node                     *Node //base
	NodeEth                  *Node
	NodeBsc                  *Node
	NodeSolana               *Node
	RomeStake                *Contract
	Pag                      *Contract
	PagClaim                 *Contract
	ERC6551Registry          *Contract
	AccountImplementAddress  *Contract
	StakingMulticallContract *Contract
	//
	RewardLogoURL string

	//
	Contracts map[string]Contract

	//
	FactoryContract string

	//
	GCSConfig GCSConfig

	//
	USDCAddress string
	RIZAddress  string
	//
	RIZSignPrivateKey string
	//
	WRIZForZnodeSignPrivateKey string
	WRizContractClaimAddress   string
	//
	RagentMarketContract           string
	RagentMarketContractPrivateKey string
	RagentWalletKeeperPrivateKey   string
	//
	BatchTransferContract       string
	BatchTransferContractEth    string
	BatchTransferContractBsc    string
	BatchTransferContractSolana string
	//
	LimitBalanceSwapWallet    int
	ValueTopupSwapWallet      int
	LimitBalanceSwapWalletEth int
	ValueTopupSwapWalletEth   int
	//
	CoinbaseAPIKey       string
	CoinbaseAPIKeySecret string
	CoinbaseWalletSecret string
}
type GCSConfig struct {
	CredentialJson string
	BucketName     string
}
type Node struct {
	NodeUrls    string
	NodeSecrets string
	NodeWeights string //total 100=100% chia cho các node theo tỉ lệ muốn: 10,20,30 sao cho tổng =100
	NodeKeys    string
}
type Contract struct {
	ContractAddress string
	ContractAbi     string
	Method          string
	Decimals        int
}
type Server struct {
	http_srv.HTTP
	ExtendConfig *ExtendConfig
}

func (sv *Server) LoadExtendConfig() {
	jwt_secret_key := sv.Config.ReadVAR("http/config/JWT_SECRET")
	jwt_expire := sv.Config.ReadVAR("http/config/JWT_EXPIRE")
	IsZnodeOwnerCheck := sv.Config.ReadVAR("http/rivalz-rome/config/IS_ZNODE_OWNER_CHECK")
	internal_secret := sv.Config.ReadVAR("http/rivalz-rome/config/INTERNAL_SECRET")
	//
	x_num_thread := sv.Config.ReadVAR("http/rivalz-rome/config/X_NUM_THREAD")
	x_page_size := sv.Config.ReadVAR("http/rivalz-rome/config/X_PAGE_SIZE")
	x_client_id := sv.Config.ReadVAR("http/rivalz-rome/config/X_CLIENT_ID")
	x_client_secret := sv.Config.ReadVAR("http/rivalz-rome/config/X_CLIENT_SECRET")
	x_sleep_time := sv.Config.ReadVAR("http/rivalz-rome/config/X_SLEEP_TIME")
	x_url := sv.Config.ReadVAR("http/rivalz-rome/config/X_URL")
	coin_gecko_url := sv.Config.ReadVAR("http/rivalz-rome/config/COINGECKO_URL")
	coin_gecko_token := sv.Config.ReadVAR("http/rivalz-rome/config/COINGECKO_TOKEN")
	//
	node_urls := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint/NODE_URLS")
	node_secrets := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint/NODE_SECRETS")
	node_weights := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint/NODE_WEIGHTS")
	node_keys := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint/NODE_KEYS")
	//
	rome_contract := sv.Config.ReadVAR("http/rivalz-rome/node/rome-project-stake-contract/CONTRACT_ADDRESS")
	rome_contract_abi := sv.Config.ReadVAR("http/rivalz-rome/node/rome-project-stake-contract/CONTRACT_ABI")
	rome_contract_method := sv.Config.ReadVAR("http/rivalz-rome/node/rome-project-stake-contract/METHOD")
	//
	pag_contract := sv.Config.ReadVAR("http/rivalz-rome/node/pag-contract/CONTRACT_ADDRESS")
	pag_decimal := sv.Config.ReadVAR("http/rivalz-rome/node/pag-contract/DECIMALS")
	//
	pag_claim_contract := sv.Config.ReadVAR("http/rivalz-rome/node/pag-claim-contract/CONTRACT_ADDRESS")
	pag_claim_decimals := sv.Config.ReadVAR("http/rivalz-rome/node/pag-claim-contract/DECIMALS")
	//
	erc6551_registry_contract := sv.Config.ReadVAR("http/rivalz-rome/node/ERC6551-Registry-Address/CONTRACT_ADDRESS")
	//
	account_implement_address := sv.Config.ReadVAR("http/rivalz-rome/node/account-implement-address/CONTRACT_ADDRESS")
	//
	reward_logo_url := sv.Config.ReadVAR("http/rivalz-rome/config/REWARD_LOGO_URL")
	//
	rnode_check_url := sv.Config.ReadVAR("http/rivalz-rome/config/RNODE_CHECK_URL")
	//
	resource_token_sign_private_key := sv.Config.ReadVAR("http/rivalz-rome/config/RESOURCE_TOKEN_SIGN_PRIVATE_KEY")
	//
	staking_multicall_contract := sv.Config.ReadVAR("http/rivalz-rome/node/staking-multi-call/CONTRACT_ADDRESS")
	//
	contracts := sv.Config.ReadVAR("http/rivalz-rome/contract/base/address")
	//
	factory_contract := sv.Config.ReadVAR("http/rivalz-rome/node/factory-contract/CONTRACT_ADDRESS")
	//
	gcs_credential_json := sv.Config.ReadVAR("http/rivalz-rome/gcs/CREDENTIAL")
	gcs_bucket_name := sv.Config.ReadVAR("http/rivalz-rome/gcs/BUCKET_NAME")
	//
	staking_contract_private_key := sv.Config.ReadVAR("http/rivalz-rome/config/STAKING_RESOURCE_PRIVATE_KEY")
	//
	usdc_address := sv.Config.ReadVAR("http/rivalz-rome/node/token-contract/USDC")
	riz_address := sv.Config.ReadVAR("http/rivalz-rome/node/token-contract/RIZ")
	internal_secret_create_project := sv.Config.ReadVAR("http/rivalz-rome/config/INTERNAL_SECRET_CREATE_PROJECT")
	//
	riz_sign_private_key := sv.Config.ReadVAR("http/rivalz-rome/config/RIZ_SIGN_PRIVATE_KEY")
	ragent_version := sv.Config.ReadVAR("http/rivalz-rome/config/RAGENT_VERSION")
	//
	wriz_for_znode_sign_private_key := sv.Config.ReadVAR("http/rivalz-rome/config/WRIZ_CLAIMED_SIGN_PRIVATE_KEY")
	wriz_contract_claim_address := sv.Config.ReadVAR("http/rivalz-rome/node/wriz-claim-contract/CONTRACT_ADDRESS")
	//
	ragent_market_contract := sv.Config.ReadVAR("http/rivalz-rome/node/ragent-market-contract/CONTRACT_ADDRESS")
	ragent_market_contract_private_key := sv.Config.ReadVAR("http/rivalz-rome/config/RAGENT_MARKET_CONTRACT_PRIVATE_KEY")
	ragent_wallet_keeper_private_key := sv.Config.ReadVAR("http/rivalz-rome/config/RAGENT_WALLET_KEEPER_PRIVATE_KEY")
	//
	internal_secret_monitoring := sv.Config.ReadVAR("http/rivalz-rome/config/INTERNAL_SECRET_MONITORING")
	//
	batch_transfer_contract := sv.Config.ReadVAR("http/rivalz-rome/node/batch-transfer-contract/CONTRACT_ADDRESS")
	batch_transfer_contract_eth := sv.Config.ReadVAR("http/rivalz-rome/node/batch-transfer-contract-eth/CONTRACT_ADDRESS")
	batch_transfer_contract_bsc := sv.Config.ReadVAR("http/rivalz-rome/node/batch-transfer-contract-bsc/CONTRACT_ADDRESS")
	batch_transfer_contract_solana := sv.Config.ReadVAR("http/rivalz-rome/node/batch-transfer-contract-solana/CONTRACT_ADDRESS")
	//
	limit_balance_swap_wallet := sv.Config.ReadVAR("http/rivalz-rome/config/LIMIT_BALANCE_SWAP_WALLET")
	value_topup_swap_wallet := sv.Config.ReadVAR("http/rivalz-rome/config/VALUE_TOPUP_SWAP_WALLET")
	limit_balance_swap_wallet_eth := sv.Config.ReadVAR("http/rivalz-rome/config/LIMIT_BALANCE_SWAP_WALLET_ETH")
	value_topup_swap_wallet_eth := sv.Config.ReadVAR("http/rivalz-rome/config/VALUE_TOPUP_SWAP_WALLET_ETH")
	//
	node_urls_eth := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-eth/NODE_URLS")
	node_secrets_eth := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-eth/NODE_SECRETS")
	node_weights_eth := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-eth/NODE_WEIGHTS")
	node_keys_eth := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-eth/NODE_KEYS")
	//
	node_urls_bsc := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-bsc/NODE_URLS")
	node_secrets_bsc := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-bsc/NODE_SECRETS")
	node_weights_bsc := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-bsc/NODE_WEIGHTS")
	node_keys_bsc := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-bsc/NODE_KEYS")
	//
	node_urls_solana := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-solana/NODE_URLS")
	node_secrets_solana := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-solana/NODE_SECRETS")
	node_weights_solana := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-solana/NODE_WEIGHTS")
	node_keys_solana := sv.Config.ReadVAR("http/rivalz-rome/node/endpoint-solana/NODE_KEYS")
	//
	coinbase_api_key := sv.Config.ReadVAR("http/rivalz-rome/config/COINBASE_API_KEY")
	coinbase_api_key_secret := sv.Config.ReadVAR("http/rivalz-rome/config/COINBASE_API_KEY_SECRET")
	coinbase_wallet_secret := sv.Config.ReadVAR("http/rivalz-rome/config/COINBASE_WALLET_SECRET")
	//
	if ragent_version == "" {
		ragent_version = "4.0.1" // Default version if not specified
	}
	if internal_secret == "" {
		panic(errors.New("INTERNAL_SECRET not found"))
	}
	if jwt_expire == "" {
		panic(errors.New("JWT_EXPIRE not found"))
	}
	if jwt_secret_key == "" {
		panic(errors.New("JWT_SECRET not found"))
	}
	if IsZnodeOwnerCheck == "" {
		panic(errors.New("IS_ZNODE_OWNER_CHECK not found"))
	}
	//
	if x_page_size == "" {
		panic(errors.New("X_PAGE_SIZE not found"))
	}
	if x_num_thread == "" {
		panic(errors.New("X_NUM_THREAD not found"))
	}
	if x_client_id == "" {
		panic(errors.New("X_CLIENT_ID not found"))
	}
	if x_client_secret == "" {
		panic(errors.New("X_CLIENT_SECRET not found"))
	}
	if x_sleep_time == "" {
		panic(errors.New("X_SLEEP_TIME not found"))
	}
	if x_url == "" {
		panic(errors.New("X_URL not found"))
	}
	if x_page_size == "" {
		panic(errors.New("X_PAGE_SIZE not found"))
	}
	//
	if coin_gecko_url == "" {
		panic(errors.New("COIN_GECKO_URL not found"))
	}
	if coin_gecko_token == "" {
		panic(errors.New("COIN_GECKO_TOKEN not found"))
	}
	//
	if node_urls == "" {
		panic(errors.New("NODE_URL not found"))
	}
	if node_secrets == "" {
		panic(errors.New("NODE_SECRET not found"))
	}
	if node_weights == "" {
		panic(errors.New("NODE_WEIGHT not found"))
	}
	if node_keys == "" {
		panic(errors.New("NODE_KEYS not found"))
	}
	//
	if rome_contract == "" {
		panic(errors.New("CONTRACT_ADDRESS not found"))
	}
	if rome_contract_abi == "" {
		panic(errors.New("CONTRACT_ABI not found"))
	}
	if rome_contract_method == "" {
		panic(errors.New("METHOD not found"))
	}
	if pag_contract == "" {
		panic(errors.New("PAG_CONTRACT not found"))
	}
	if pag_decimal == "" {
		panic(errors.New("DECIMALS not found"))
	}
	if pag_claim_contract == "" {
		panic(errors.New("PAG_CLAIM_CONTRACT not found"))
	}
	if pag_claim_decimals == "" {
		panic(errors.New("DECIMALS not found"))
	}
	if erc6551_registry_contract == "" {
		panic(errors.New("ERC6551-Registry-Address not found"))
	}
	if account_implement_address == "" {
		panic(errors.New("ACCOUNT_IMPLEMENT_ADDRESS not found"))
	}
	//
	if reward_logo_url == "" {
		panic(errors.New("REWARD_LOGO_URL not found"))
	}
	//
	if rnode_check_url == "" {
		panic(errors.New("RNODE_CHECK_URL not found"))
	}
	//
	if resource_token_sign_private_key == "" {
		panic(errors.New("RESOURCE_TOKEN_SIGN_PRIVATE_KEY not found"))
	}
	//
	if staking_multicall_contract == "" {
		panic(errors.New("STAKING_MULTI_CALL not found"))
	}
	//
	if contracts == "" {
		panic(errors.New("CONTRACTS not found"))
	}
	//
	if factory_contract == "" {
		panic(errors.New("FACTORY_CONTRACT not found"))
	}
	//
	if gcs_credential_json == "" {
		panic(errors.New("GCS_CREDENTIAL_JSON not found"))
	}
	if gcs_bucket_name == "" {
		panic(errors.New("GCS_BUCKET_NAME not found"))
	}
	//
	if staking_contract_private_key == "" {
		panic(errors.New("STAKING_CONTRACT_PRIVATE_KEY not found"))
	}
	//
	if usdc_address == "" {
		panic(errors.New("USDC_ADDRESS not found"))
	}
	if riz_address == "" {
		panic(errors.New("RIZ_ADDRESS not found"))
	}
	if internal_secret_create_project == "" {
		panic(errors.New("INTERNAL_SECRET_CREATE_PROJECT not found"))
	}
	if riz_sign_private_key == "" {
		panic(errors.New("RIZ_SIGN_PRIVATE_KEY not found"))
	}
	if wriz_for_znode_sign_private_key == "" {
		panic(errors.New("WRIZ_FOR_ZNODE_SIGN_PRIVATE_KEY not found"))
	}
	if wriz_contract_claim_address == "" {
		panic(errors.New("WRIZ_CONTRACT_CLAIM_ADDRESS not found"))
	}
	if internal_secret_monitoring == "" {
		panic(errors.New("INTERNAL_SECRET_MONITORING not found"))
	}
	if ragent_market_contract == "" {
		panic(errors.New("RAGENT_MARKET_CONTRACT not found"))
	}
	if ragent_market_contract_private_key == "" {
		panic(errors.New("RAGENT_MARKET_CONTRACT_PRIVATE_KEY not found"))
	}
	if ragent_wallet_keeper_private_key == "" {
		panic(errors.New("RAGENT_WALLET_KEEPER_PRIVATE_KEY not found"))
	}
	if batch_transfer_contract == "" {
		panic(errors.New("BATCH_TRANSFER_CONTRACT not found"))
	}
	if batch_transfer_contract_eth == "" {
		panic(errors.New("BATCH_TRANSFER_CONTRACT_ETH not found"))
	}
	if batch_transfer_contract_bsc == "" {
		panic(errors.New("BATCH_TRANSFER_CONTRACT_BSC not found"))
	}
	if batch_transfer_contract_solana == "" {
		panic(errors.New("BATCH_TRANSFER_CONTRACT_SOLANA not found"))
	}
	if limit_balance_swap_wallet == "" {
		panic(errors.New("LIMIT_BALANCE_SWAP_WALLET not found"))
	}
	if value_topup_swap_wallet == "" {
		panic(errors.New("VALUE_TOPUP_SWAP_WALLET not found"))
	}
	if limit_balance_swap_wallet_eth == "" {
		panic(errors.New("LIMIT_BALANCE_SWAP_WALLET_ETH not found"))
	}
	if value_topup_swap_wallet_eth == "" {
		panic(errors.New("VALUE_TOPUP_SWAP_WALLET_ETH not found"))
	}
	//
	if node_urls_eth == "" {
		panic(errors.New("NODE_URLS_ETH not found"))
	}
	if node_secrets_eth == "" {
		panic(errors.New("NODE_SECRETS_ETH not found"))
	}
	if node_weights_eth == "" {
		panic(errors.New("NODE_WEIGHTS_ETH not found"))
	}
	if node_keys_eth == "" {
		panic(errors.New("NODE_KEYS_ETH not found"))
	}
	//
	if node_urls_bsc == "" {
		panic(errors.New("NODE_URLS_BSC not found"))
	}
	if node_secrets_bsc == "" {
		panic(errors.New("NODE_SECRETS_BSC not found"))
	}
	if node_weights_bsc == "" {
		panic(errors.New("NODE_WEIGHTS_BSC not found"))
	}
	if node_keys_bsc == "" {
		panic(errors.New("NODE_KEYS_BSC not found"))
	}
	//
	if node_urls_solana == "" {
		panic(errors.New("NODE_URLS_SOLANA not found"))
	}
	if node_secrets_solana == "" {
		panic(errors.New("NODE_SECRETS_SOLANA not found"))
	}
	if node_weights_solana == "" {
		panic(errors.New("NODE_WEIGHTS_SOLANA not found"))
	}
	if node_keys_solana == "" {
		panic(errors.New("NODE_KEYS_SOLANA not found"))
	}
	//
	if coinbase_api_key == "" {
		panic(errors.New("COINBASE_API_KEY not found"))
	}
	if coinbase_api_key_secret == "" {
		panic(errors.New("COINBASE_API_KEY_SECRET not found"))
	}
	if coinbase_wallet_secret == "" {
		panic(errors.New("COINBASE_WALLET_SECRET not found"))
	}
	// contractsMap := make(map[string]Contract)
	// for _, contract := range contracts {
	// 	contractMap, ok := contract.(map[string]interface{})
	// 	if !ok {
	// 		panic(errors.New("CONTRACTS is invalid"))
	// 	}
	// 	for key, value := range contractMap {
	// 		valueMap, ok := value.(map[string]interface{})
	// 		if !ok {
	// 			panic(errors.New("CONTRACTS is invalid"))
	// 		}
	// 		contractsMap[key] = Contract{
	// 			ContractAddress: valueMap["address"].(string),
	// 			Decimals:        valueMap["decimals"].(int),
	// 		}
	// 	}
	// }
	//
	sv.ExtendConfig = &ExtendConfig{
		JwtSecretKey:      jwt_secret_key,
		JwtExpire:         utils.ItoInt(jwt_expire),
		IsZnodeOwnerCheck: utils.StringToBool(IsZnodeOwnerCheck),
		//
		InternalSecret: internal_secret,
		//X social
		XNumThread: utils.ItoInt(x_num_thread),
		XPageSize:  utils.ItoInt(x_page_size),
		//filter data with current_time - created_at < XExecBeforeRemainingDay
		XClientID:     x_client_id,
		XClientSecret: x_client_secret,
		XSleepTime:    utils.ItoInt(x_sleep_time),
		XURL:          x_url,
		//
		CoinGeckoURL:   coin_gecko_url,
		CoinGeckoToken: coin_gecko_token,
		//
		Node: &Node{
			NodeUrls:    node_urls,
			NodeSecrets: node_secrets,
			NodeWeights: node_weights,
			NodeKeys:    node_keys,
		},
		NodeEth: &Node{
			NodeUrls:    node_urls_eth,
			NodeSecrets: node_secrets_eth,
			NodeWeights: node_weights_eth,
			NodeKeys:    node_keys_eth,
		},
		NodeBsc: &Node{
			NodeUrls:    node_urls_bsc,
			NodeSecrets: node_secrets_bsc,
			NodeWeights: node_weights_bsc,
			NodeKeys:    node_keys_bsc,
		},
		NodeSolana: &Node{
			NodeUrls:    node_urls_solana,
			NodeSecrets: node_secrets_solana,
			NodeWeights: node_weights_solana,
			NodeKeys:    node_keys_solana,
		},
		RomeStake: &Contract{
			ContractAddress: rome_contract,
			ContractAbi:     rome_contract_abi,
			Method:          rome_contract_method,
		},
		Pag: &Contract{
			ContractAddress: pag_contract,
			ContractAbi:     "",
			Method:          "",
			Decimals:        utils.ItoInt(pag_decimal),
		},
		PagClaim: &Contract{
			ContractAddress: pag_claim_contract,
			ContractAbi:     "",
			Method:          "",
			Decimals:        utils.ItoInt(pag_claim_decimals),
		},
		ERC6551Registry: &Contract{
			ContractAddress: erc6551_registry_contract,
			ContractAbi:     "",
			Method:          "",
			Decimals:        0,
		},
		AccountImplementAddress: &Contract{
			ContractAddress: account_implement_address,
			ContractAbi:     "",
			Method:          "",
			Decimals:        0,
		},
		RewardLogoURL:               reward_logo_url,
		RNodeCheckURL:               rnode_check_url,
		ResourceTokenSignPrivateKey: resource_token_sign_private_key,
		StakingMulticallContract: &Contract{
			ContractAddress: staking_multicall_contract,
			ContractAbi:     "",
			Method:          "",
			Decimals:        0,
		},
		Contracts: map[string]Contract{
			"base": {
				ContractAddress: contracts,
				ContractAbi:     "",
				Method:          "",
				Decimals:        8,
			},
		},
		FactoryContract: factory_contract,
		GCSConfig: GCSConfig{
			CredentialJson: gcs_credential_json,
			BucketName:     gcs_bucket_name,
		},
		StakingContractPrivateKey:      staking_contract_private_key,
		USDCAddress:                    usdc_address,
		RIZAddress:                     riz_address,
		InternalSecretCreateProject:    internal_secret_create_project,
		RIZSignPrivateKey:              riz_sign_private_key,
		RAgentVersion:                  ragent_version,
		WRIZForZnodeSignPrivateKey:     wriz_for_znode_sign_private_key,
		WRizContractClaimAddress:       wriz_contract_claim_address,
		InternalSecretMonitoring:       internal_secret_monitoring,
		RagentMarketContract:           ragent_market_contract,
		RagentMarketContractPrivateKey: ragent_market_contract_private_key,
		RagentWalletKeeperPrivateKey:   ragent_wallet_keeper_private_key,
		BatchTransferContract:          batch_transfer_contract,
		BatchTransferContractEth:       batch_transfer_contract_eth,
		BatchTransferContractBsc:       batch_transfer_contract_bsc,
		BatchTransferContractSolana:    batch_transfer_contract_solana,
		LimitBalanceSwapWallet:         utils.ItoInt(limit_balance_swap_wallet),
		ValueTopupSwapWallet:           utils.ItoInt(value_topup_swap_wallet),
		LimitBalanceSwapWalletEth:      utils.ItoInt(limit_balance_swap_wallet_eth),
		ValueTopupSwapWalletEth:        utils.ItoInt(value_topup_swap_wallet_eth),
		CoinbaseAPIKey:                 coinbase_api_key,
		CoinbaseAPIKeySecret:           coinbase_api_key_secret,
		CoinbaseWalletSecret:           coinbase_wallet_secret,
	}
	if sv.ExtendConfig.XNumThread <= 0 {
		panic(errors.New("X_NUM_THREAD is invalid"))
	}
	if sv.ExtendConfig.XPageSize <= 0 {
		panic(errors.New("X_PAGE_SIZE is invalid"))
	}
	if sv.ExtendConfig.JwtExpire <= 0 {
		panic(errors.New("JWT_EXPIRE is invalid"))
	}
	if sv.ExtendConfig.Pag.Decimals <= 0 {
		panic(errors.New("PAG DECIMALS is invalid"))
	}
	if sv.ExtendConfig.PagClaim.Decimals <= 0 {
		panic(errors.New("PAG CLAIM DECIMALS is invalid"))
	}
}
