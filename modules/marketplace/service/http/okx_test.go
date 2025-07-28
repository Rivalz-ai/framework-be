package http_test

import (
	"fmt"
	"testing"

	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/service/http"
)

func TestGetOkxTokenPrice(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		tokenAddress string
		decimals     int
		want         float64
		want2        float64
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:         "test",
			tokenAddress: "0x91da780bc7f4b7cf19abe90411a2a296ec5ff787",
			decimals:     18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2, gotErr := http.GetOkxTokenPrice(tt.tokenAddress, tt.decimals, int(define.BASE))
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetOkxTokenPrice() failed: %v", gotErr)
				}
				return
			}

			fmt.Println(got, got2)
		})
	}
}

func TestGetOkxSwapWithAutoSlippage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fromTokenAddress string
		toTokenAddress   string
		amount           string
		userWallet       string
		want             *dto.OkxSwapResponse
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name:             "test swap USDC to WETH",
			fromTokenAddress: "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913",
			toTokenAddress:   "0x67543cf0304c19ca62ac95ba82fd4f4b40788dc1",
			amount:           "1000000",
			userWallet:       "0x742d35Cc6634C0532925a3b8D4C9db96C4b4d4d4",
		},
		{
			name:             "test swap WETH to USDC",
			fromTokenAddress: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			toTokenAddress:   "0xa0b86a33e6ba3e0b4aaa2b4b8b8b8b8b8b8b8b8b",
			amount:           "1000000000000000000",
			userWallet:       "0x742d35Cc6634C0532925a3b8D4C9db96C4b4d4d4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := http.GetOkxSwapWithAutoSlippage(tt.fromTokenAddress, tt.toTokenAddress, tt.amount, tt.userWallet)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetOkxSwapWithAutoSlippage() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetOkxSwapWithAutoSlippage() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetOkxSwapWithAutoSlippage() = %v, want %v", got, tt.want)
			}
		})
	}
}
