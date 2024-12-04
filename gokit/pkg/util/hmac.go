package util

import (
	"crypto/hmac"
	"encoding/hex"
	"hash"
)

func SignHMAC(h func() hash.Hash, msg, key []byte) (string, error) {
	mac := hmac.New(h, key)
	_, err := mac.Write(msg)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(mac.Sum(nil)), nil
}

func VerifyHMAC(h func() hash.Hash, msg, key []byte, hash string) (bool, error) {
	sig, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	mac := hmac.New(h, key)
	_, err = mac.Write(msg)
	if err != nil {
		return false, err
	}

	return hmac.Equal(sig, mac.Sum(nil)), nil
}
