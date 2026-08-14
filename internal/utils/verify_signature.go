package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func VerifySignature(signature []byte, payload []byte, secret []byte) bool {
	hash := hmac.New(sha256.New, secret)
	hash.Write(payload)
	code := hex.EncodeToString(hash.Sum(nil))
	return hmac.Equal(signature, []byte("sha256="+code))
}
