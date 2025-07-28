package http

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	cdp "github.com/coinbase/cdp-sdk/go"
	"github.com/coinbase/cdp-sdk/go/openapi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	redis "github.com/redis/go-redis/v9"
)

var (
	// apiKeyID     = "33c6474b-9081-4f86-bfb8-6b779e300aa2"
	// apiKeySecret = "HucsfsflhwLurywQlIPvSw5YU+nhy5h7Izm5c1isBQGHOUNN11yUMO4p3wPDKT05xYlnE3wOZlsAGbpbxPGcZg=="
	// walletSecret = "MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgbkqcE68A96F+seUKtG5ZzFM9cXnjEJGisds/ywvY2iShRANCAAQmEwWiLmp/Ay1M27gkjAFqgq3kb+N1MO8CT9f9bnvqboKmt1ZqPQ6wWjfHpcEAH64BKMPndwMhtq8jRov7vG87"

	RPC_URL = "https://ethereum-sepolia-rpc.publicnode.com"
	//wallet was created on coinbase before for test send a TX
	evm_wallet = "0xF8cb7e11BcbbB10f781a9d53E42AC4789034DA48"
	to_wallet  = "0x1a474e25799CF40593ac0C3d8f720a731BDe1d33"
	/*evm: 0x1a474e25799CF40593ac0C3d8f720a731BDe1d33
	sol: BEN4wsX1fPKfJo5BGwwbYsPwDTmgXKYjiYQGYZL52gf9 */

)

type CoinbaseClient struct {
	CdpClient *openapi.ClientWithResponses
}

var coinbaseClient *CoinbaseClient

func NewCoinbaseClient(apiKeyID, apiKeySecret, walletSecret string) (*CoinbaseClient, error) {
	if coinbaseClient != nil {
		return coinbaseClient, nil
	}

	cdpClient, err := createCDPClient(apiKeyID, apiKeySecret, walletSecret)
	if err != nil {
		return nil, err
	}
	coinbaseClient = &CoinbaseClient{CdpClient: cdpClient}
	return coinbaseClient, nil
}

// createCDPClient creates and returns a new CDP client using environment variables for configuration.
func createCDPClient(apiKeyID, apiKeySecret, walletSecret string) (*openapi.ClientWithResponses, error) {

	if apiKeyID == "" {
		log.Fatal("CDP_API_KEY_ID environment variable is required")
	}
	if apiKeySecret == "" {
		log.Fatal("CDP_API_KEY_SECRET environment variable is required")
	}
	if walletSecret == "" {
		log.Fatal("CDP_WALLET_SECRET environment variable is required")
	}

	cdp, err := cdp.NewClient(cdp.ClientOptions{
		APIKeyID:     apiKeyID,
		APIKeySecret: apiKeySecret,
		WalletSecret: walletSecret,
		//BasePath:     apiURL,
	})
	if err != nil {
		return nil, err
	}

	return cdp, nil
}

func exampleUse(apiKeyID, apiKeySecret, walletSecret string) {
	ctx := context.Background()

	coinbaseClient, err := NewCoinbaseClient(apiKeyID, apiKeySecret, walletSecret)
	if err != nil {
		log.Fatalf("Failed to create CDP client: %v", err)
	}

	/*evmAddress, err := coinbaseClient.CreateEVMAccount(ctx)
	if err != nil {
		log.Printf("Failed to create EVM account: %v", err)
	}
	fmt.Println("EVM address:", evmAddress)
	*/
	//create tx
	signedTransaction, _, _, err := coinbaseClient.CreateAndSignEVMTransaction(ctx, nil, RPC_URL, evm_wallet, to_wallet, 11155111, []byte{})
	if err != nil {
		log.Printf("Failed to sign transaction: %v", err)
	}
	//sign tx
	// if err := coinbaseClient.SendSignedEVMTransaction(ctx, signedTransaction); err != nil {
	// 	log.Printf("Failed to send transaction: %v", err)
	// }
	fmt.Println("Transaction sent:", signedTransaction)

}

func exampleUseSolana() {
	// ctx := context.Background()
	// client, err := createCDPClient()
	// if err != nil {
	// 	log.Fatalf("CDP client error: %v", err)
	// }
	/*
		fmt.Println("Creating wallet 1...")
		wallet1, err := createSolanaAccount(ctx, client)
		if err != nil {
			log.Fatalf("Create wallet 1 error: %v", err)
		}
		fmt.Println("Wallet 1:", wallet1)

		fmt.Println("Creating wallet 2...")
		wallet2, err := createSolanaAccount(ctx, client)
		if err != nil {
			log.Fatalf("Create wallet 2 error: %v", err)
		}
		fmt.Println("Wallet 2:", wallet2)
	*/
	// wallet1 := "B7zh8wsADy6LMykD2YfVPrmM2uY3xJi6QYUgMUMtRt4m"
	// wallet2 := "E2L3wGBndw5DyKuuQKjUCs3RTZ1ywSWSsM8mzNRt3qgp"
	// bal1, _ := getSolanaBalance(ctx, wallet1)
	// bal2, _ := getSolanaBalance(ctx, wallet2)
	// fmt.Printf("Wallet 1 balance: %.9f SOL\n", bal1)
	// fmt.Printf("Wallet 2 balance: %.9f SOL\n", bal2)
	// return
	// if bal1 < TRANSFER_AMOUNT {
	// 	fmt.Println("Wallet 1 balance is less than transfer amount")
	// 	return
	// }
	// fmt.Println("Sending transaction from wallet 1 to wallet 2...")
	// sig, err := sendSolanaTransaction(ctx, client, wallet1, wallet2, TRANSFER_AMOUNT)
	// if err != nil {
	// 	log.Fatalf("Send transaction error: %v", err)
	// }
	// fmt.Println("Transaction signature:", sig)

	// time.Sleep(5 * time.Second)
	// bal1After, _ := getSolanaBalance(ctx, wallet1)
	// bal2After, _ := getSolanaBalance(ctx, wallet2)
	// fmt.Printf("Wallet 1 balance after: %.9f SOL\n", bal1After)
	// fmt.Printf("Wallet 2 balance after: %.9f SOL\n", bal2After)
}

// createEVMAccount creates a new EVM account using the CDP client.
func (c *CoinbaseClient) CreateEVMAccount(ctx context.Context) (string, error) {
	log.Println("Creating EVM account...")

	response, err := c.CdpClient.CreateEvmAccountWithResponse(
		ctx,
		nil,
		openapi.CreateEvmAccountJSONRequestBody{},
	)
	if err != nil {
		return "", err
	}

	if response.StatusCode() != 201 {
		return "", fmt.Errorf("failed to create EVM account: %v", response.Status())
	}

	evmAddress := response.JSON201.Address
	log.Printf("EVM account created: %v", evmAddress)
	return evmAddress, nil
}

// createAndSignEVMTransaction creates and signs an EVM transaction using the CDP client.
func (c *CoinbaseClient) CreateAndSignEVMTransaction(ctx context.Context, cacheClient *redis.Client, nodeUrl string, evmAddress, toContractAddress string, chainId int, data []byte) (string, uint64, *big.Int, error) {
	toAddress := common.HexToAddress(toContractAddress)

	etherClient, err := ethclient.Dial(nodeUrl)
	if err != nil {
		return "", 0, nil, fmt.Errorf("failed to create ether client: %v", err)
	}

	// Create a call message to estimate gas
	msg := ethereum.CallMsg{
		From: common.HexToAddress(evmAddress),
		To:   &toAddress,
		Data: data,
	}

	gasLimit, err := etherClient.EstimateGas(ctx, msg)
	if err != nil {
		return "", 0, nil, fmt.Errorf("failed to estimate gas: %v", err)
	}

	nonece, err := etherClient.PendingNonceAt(ctx, common.HexToAddress(evmAddress))
	if err != nil {
		return "", 0, nil, fmt.Errorf("failed to get nonce: %v", err)
	}

	gasPricetmp, err := etherClient.SuggestGasPrice(ctx)
	if err != nil {
		return "", 0, nil, fmt.Errorf("failed to get gas price: %v", err)
	}
	gasPricetmp.Mul(gasPricetmp, big.NewInt(115))
	gasPricetmp.Div(gasPricetmp, big.NewInt(100))

	etherscanClient := NewEtherscanClient(cacheClient)
	gasPriceResp, err := etherscanClient.GetGasPrice(ctx, chainId)
	if err != nil {
		return "", 0, nil, fmt.Errorf("failed to get gas price: %v", err)
	}
	gasPriceFl, ok := new(big.Float).SetString(gasPriceResp.FastGasPrice)
	if !ok {
		return "", 0, nil, fmt.Errorf("failed to set gas price: %v", err)
	}
	gasPriceFl.Mul(gasPriceFl, big.NewFloat(1e9))
	gasPrice, _ := gasPriceFl.Int(nil)

	if gasPricetmp.Cmp(gasPrice) > 0 {
		gasPrice = gasPricetmp
	}

	gasTipCap, err := etherClient.SuggestGasTipCap(ctx)
	if err != nil {
		return "", 0, nil, fmt.Errorf("failed to get gas tip cap: %v", err)
	}

	// Add some buffer to the gas limit (20% more)
	gasLimit = gasLimit * 120 / 100

	transaction := types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainId)),
		Nonce:     nonece,
		To:        &toAddress,
		Value:     big.NewInt(0),
		Data:      data,
		Gas:       gasLimit,
		GasFeeCap: gasPrice,
		GasTipCap: gasTipCap,
	}

	// Serialize transaction to RLP
	rlpTx := types.NewTx(&transaction)
	rlpData, err := rlpTx.MarshalBinary()
	if err != nil {
		return "", 0, nil, err
	}

	rlpHex := hex.EncodeToString(rlpData)
	rlpHex = "0x" + rlpHex

	response, err := c.CdpClient.SignEvmTransactionWithResponse(
		ctx,
		evmAddress,
		nil,
		openapi.SignEvmTransactionJSONRequestBody{
			Transaction: rlpHex,
		},
	)
	if err != nil {
		return "", 0, nil, err
	}

	if response.StatusCode() != 200 {
		return "", 0, nil, fmt.Errorf("failed to sign transaction: %v", response.Status())
	}

	log.Printf("Signed transaction: %v", response.JSON200.SignedTransaction)
	return response.JSON200.SignedTransaction, uint64(gasLimit), gasPrice, nil
}

func (c *CoinbaseClient) CreateSolanaAccount(ctx context.Context) (string, error) {
	resp, err := c.CdpClient.CreateSolanaAccountWithResponse(ctx, nil, openapi.CreateSolanaAccountJSONRequestBody{})
	if err != nil {
		return "", err
	}
	if resp.JSON201 == nil {
		return "", fmt.Errorf("No address returned")
	}
	return resp.JSON201.Address, nil
}
