package http

import (
	"context"
	"testing"
)

func TestGetLiFiStatus(t *testing.T) {
	// Test with a sample transaction hash
	txHash := "0xf638b65d8340d26ad08edba7aae77d51eeceaf96b3394fff3eb2590e63ef8d40"

	ctx := context.Background()

	result, err := GetLiFiStatus(ctx, txHash, 1, 56)
	if err != nil {
		t.Fatalf("GetLiFiStatus failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected result to not be nil")
	}

	// Print the result for verification
	t.Logf("Transaction ID: %s", result.TransactionId)
	t.Logf("Status: %s", result.Status)
	t.Logf("SubStatus: %s", result.SubStatus)
	t.Logf("SubStatus Message: %s", result.SubStatusMessage)
	t.Logf("Tool: %s", result.Tool)
	t.Logf("From Address: %s", result.FromAddress)
	t.Logf("To Address: %s", result.ToAddress)

	// Verify essential fields are present
	if result.TransactionId == "" {
		t.Error("Expected TransactionId to not be empty")
	}

	if result.Status == "" {
		t.Error("Expected Status to not be empty")
	}
}
