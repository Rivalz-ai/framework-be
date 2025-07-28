package utils

import (
	"context"
	"encoding/base64"
	"fmt"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"
)

const (
	projectID  = "rivalz-be"
	locationID = "asia-southeast1"
)

// KMSClient wraps Google KMS logic
type KMSClient struct {
	ctx     context.Context
	client  *kms.KeyManagementClient
	keyName string
}

// NewKMSClient initializes the KMSClient with given service account JSON and KMS key path
func NewKMSClient(ctx context.Context, saJSON []byte, keyRingID, cryptoKeyID string) (*KMSClient, error) {
	keyName := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s",
		projectID, locationID, keyRingID, cryptoKeyID)

	client, err := kms.NewKeyManagementClient(ctx, option.WithCredentialsJSON(saJSON))
	if err != nil {
		return nil, err
	}

	return &KMSClient{
		ctx:     ctx,
		client:  client,
		keyName: keyName,
	}, nil
}

// Encrypt encrypts a plaintext string using Google KMS
func (k *KMSClient) Encrypt(plaintext string) (string, error) {
	req := &kmspb.EncryptRequest{
		Name:      k.keyName,
		Plaintext: []byte(plaintext),
	}
	resp, err := k.client.Encrypt(k.ctx, req)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(resp.Ciphertext), nil
}

// Decrypt decrypts a base64-encoded ciphertext using Google KMS
func (k *KMSClient) Decrypt(b64Ciphertext string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(b64Ciphertext)
	if err != nil {
		return "", err
	}
	req := &kmspb.DecryptRequest{
		Name:       k.keyName,
		Ciphertext: ciphertext,
	}
	resp, err := k.client.Decrypt(k.ctx, req)
	if err != nil {
		return "", err
	}
	return string(resp.Plaintext), nil
}

// Close releases the KMS client
func (k *KMSClient) Close() error {
	return k.client.Close()
}
