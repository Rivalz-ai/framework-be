package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	e "github.com/Rivalz-ai/framework-be/framework/base/error"
	//"fmt"
)

func RSA_Generate_KEY(size int) (string, string, *e.Error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return "", "", e.NewErr(err, "GENERATE_KEY")
	}
	pri_str := ExportRsaPrivateKeyAsStr(privateKey)
	publicKey := privateKey.PublicKey
	pub_str, err_s := ExportRsaPublicKeyAsStr(&publicKey)
	if err_s != nil {
		return "", "", err_s
	}
	return pri_str, pub_str, nil
}
func RSA_OAEP_Encrypt(secretMessage string, pub_key string) (string, *e.Error) {
	key, err_s := ParseRsaPublicKeyFromStr(pub_key)
	if err_s != nil {
		return "", err_s
	}
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, key, []byte(secretMessage), label)
	if err != nil {
		return "", e.NewErr(err, "ENCRYPT")
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func RSA_OAEP_Decrypt(cipherText string, pri_key string) (string, *e.Error) {
	privKey, err_s := ParseRsaPrivateKeyFromStr(pri_key)
	if err_s != nil {
		return "", err_s
	}
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, privKey, ct, label)
	if err != nil {
		return "", e.NewErr(err, "DECRTYPT")
	}
	return string(plaintext), nil
}

func ExportRsaPrivateKeyAsStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ParseRsaPrivateKeyFromStr(privPEM string) (*rsa.PrivateKey, *e.Error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, e.NewErr(errors.New("failed to parse PEM block containing the key"), "DECODE")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, e.NewErr(err, "PARSE_PCKS_PRIVATE_KEY")
	}

	return priv, nil
}

func ExportRsaPublicKeyAsStr(pubkey *rsa.PublicKey) (string, *e.Error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", e.NewErr(err, "MARSHAL_PKI_PUBLIC_KEY")
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

func ParseRsaPublicKeyFromStr(pubPEM string) (*rsa.PublicKey, *e.Error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, e.NewErr(errors.New("failed to parse PEM block containing the key"), "DECODE_PEM")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, e.NewErr(err, "MARSHAL_PKI_PUBLIC_KEY")
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, e.NewErr(errors.New("Key type is not RSA"), "MARSHAL_PKI_PUBLIC_KEY")
}
