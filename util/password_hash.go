package util

import (
	"crypto"
	"encoding/hex"
)

func GetPasswordHash(password string) (hashString string) {
	hash := crypto.SHA512.New()

	hash.Write(
		[]byte(password),
	)

	hashString = hex.EncodeToString(
		hash.Sum(nil),
	)

	return
}
