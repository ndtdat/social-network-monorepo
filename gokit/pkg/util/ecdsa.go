package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
)

func GenerateECDSAKey(c elliptic.Curve) (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func EncodeECDSAPubKey(pubKey *ecdsa.PublicKey) ([]byte, error) {
	x509EncodedPub, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub}), nil
}

func EncodeECDSAPrivKey(privKey *ecdsa.PrivateKey) ([]byte, error) {
	x509Encoded, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded}), nil
}

func GenerateECDSAPEM(c elliptic.Curve) (string, string, error) {
	privateKey, err := GenerateECDSAKey(c)
	if err != nil {
		return "", "", err
	}

	pubPemEncoded, err := EncodeECDSAPubKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}

	privPemEncodedPub, err := EncodeECDSAPrivKey(privateKey)
	if err != nil {
		return "", "", err
	}

	return string(pubPemEncoded), string(privPemEncodedPub), nil
}
