package eddsa

import (
	"crypto/ed25519"
	"github.com/dolthub/swiss"
	"github.com/kataras/jwt"
	"os"
)

func LoadPublicKey(keyFile string) (ed25519.PublicKey, error) {
	bs, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	return ParsePublicKey(bs)
}

func LoadPrivateKey(keyFile string) (ed25519.PrivateKey, error) {
	bs, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	return ParsePrivateKey(bs)
}

func ParsePrivateKey(payload []byte) (ed25519.PrivateKey, error) {
	key, err := jwt.ParsePrivateKeyEdDSA(payload)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func ParsePublicKey(payload []byte) (ed25519.PublicKey, error) {
	key, err := jwt.ParsePublicKeyEdDSA(payload)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func InitPublicKeyMapFromFile(file string, currentVersion string) (*swiss.Map[string, ed25519.PublicKey], error) {
	publicKeyMap := swiss.NewMap[string, ed25519.PublicKey](42)

	if publicKey, _ := LoadPublicKey(file); publicKey != nil {
		publicKeyMap.Put(currentVersion, publicKey)
	}

	return publicKeyMap, nil
}
