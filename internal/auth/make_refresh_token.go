package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() (string, error) {
	tokenSrc := make([]byte, 32)
	rand.Read(tokenSrc)
	tokenStr := hex.EncodeToString(tokenSrc)
	return tokenStr, nil
}
