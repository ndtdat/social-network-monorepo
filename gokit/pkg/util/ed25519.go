package util

import (
	"crypto/ed25519"
	"github.com/kataras/jwt"
	"os"
)

func GenerateEdDSAKeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return jwt.GenerateEdDSA()
}

func StoreEdDSAKeyPair(
	publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey, publicKeyFilePath, privateKeyFilePath string,
) error {
	if err := os.WriteFile(publicKeyFilePath, publicKey, 0600); err != nil {
		return err
	}

	return os.WriteFile(privateKeyFilePath, privateKey, 0600)
}
