package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
)

// GenerateSignature generates a signature from key and writes it to q.
func GenerateSignature(key string, q url.Values) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(q.Encode()))

	expectedMAC := mac.Sum(nil)

	return hex.EncodeToString(expectedMAC)
}
