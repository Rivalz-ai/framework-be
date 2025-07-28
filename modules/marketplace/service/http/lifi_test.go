package http_test

import (
	"testing"

	"github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/service/http"
)

func TestGetLiFiQuote(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		request *dto.LiFiQuoteRequest
		want    *dto.LiFiQuoteResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid cross-chain swap request",
			request: &dto.LiFiQuoteRequest{
				ToAmount: "17293822569102704640", // 100 token with 18 decimals
				ContractCalls: []dto.LiFiContractCall{
					{
						// FromAmount:         "1000000000000000000000",
						FromTokenAddress:   "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
						ToContractAddress:  "0xaea46a60368a7bd060eec7df8cba43b7ef41ad85",
						ToContractCallData: "0x",
						ToContractGasLimit: "200000",
					},
				},
				FromChain:   8453, // Base mainnet
				FromToken:   "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
				FromAddress: "0xc5815b405F6af3B80e2E2c2A2766E41fC584459b",
				ToChain:     1, // Ethereum mainnet
				ToToken:     "0xaea46a60368a7bd060eec7df8cba43b7ef41ad85",
			},
			want:    nil, // We'll check the response structure in the test
			wantErr: false,
		},
		// {
		// 	name: "invalid request with empty fields",
		// 	request: &dto.LiFiQuoteRequest{
		// 		ToAmount:      "",
		// 		ContractCalls: []dto.LiFiContractCall{},
		// 		FromChain:     0,
		// 		FromToken:     "",
		// 		FromAddress:   "",
		// 		ToChain:       0,
		// 		ToToken:       "",
		// 	},
		// 	want:    nil,
		// 	wantErr: true,
		// },
		// {
		// 	name:    "nil request",
		// 	request: nil,
		// 	want:    nil,
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := http.BroadcastLiFiTransaction(tt.request)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetLiFiQuote() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetLiFiQuote() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetLiFiQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}
