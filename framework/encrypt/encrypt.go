package encrypt

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	e "github.com/Rivalz-ai/framework-be/framework/base/error"
	"golang.org/x/crypto/bcrypt"
)

// hash function
func HashBcrypt(password string) (string, *e.Error) {
	//bcrypt.GenerateFromPassword is auto generate Salt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", e.NewErr(err, "HASH_PASSWORD")
	}
	return string(hash), nil
}

// check plain text string same with hash string
func VerifyHashBcrypt(hashedText, plainText string) *e.Error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))
	if err != nil {
		return e.NewErr(err, "COMPARE_HASH_PASSWORD")
	}
	return nil
}
func HashMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func Base64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func SHA256(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	sha256_hash := hex.EncodeToString(h.Sum(nil))
	return sha256_hash
}
