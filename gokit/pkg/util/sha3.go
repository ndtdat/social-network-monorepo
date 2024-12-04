package util

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

func LegacyKeccak256(input string) string {
	h := sha3.NewLegacyKeccak256()
	h.Write([]byte(input))

	return hex.EncodeToString(h.Sum(nil))
}

func Keccak256(input string) string {
	h := sha3.New256()
	h.Write([]byte(input))

	return hex.EncodeToString(h.Sum(nil))
}
