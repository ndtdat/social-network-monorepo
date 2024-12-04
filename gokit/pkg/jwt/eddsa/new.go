package eddsa

import (
	"context"
	"crypto/ed25519"
	"errors"
	"fmt"
	"github.com/dolthub/swiss"
	kJwt "github.com/kataras/jwt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	base2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/base"
	"go.uber.org/zap"
	"time"
)

type Manager struct {
	alg          kJwt.Alg
	publicKeyMap *swiss.Map[string, ed25519.PublicKey]

	logger *zap.Logger

	privateKeySecretName  string
	currPrivateKeyVersion string

	publicKeySecretName string

	issuer               string
	privateKey           ed25519.PrivateKey
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewManager(
	ctx context.Context, cfg *config.App, logger *zap.Logger,
) (base2.Manager, error) {
	JWTCfg := cfg.JWT

	privateKeyCfg := JWTCfg.PrivateKey
	publicKeyCfg := JWTCfg.PublicKey

	privateKeySecretName := privateKeyCfg.Secret.Name
	privateKeyFile := privateKeyCfg.File

	publicKeySecretName := publicKeyCfg.Secret.Name
	publicKeyFile := publicKeyCfg.File

	if publicKeySecretName == "" && publicKeyFile == "" {
		return nil, fmt.Errorf("public key secret name or public key file must be exist")
	}

	var (
		err          error
		publicKeyMap *swiss.Map[string, ed25519.PublicKey]
		ecPrivateKey ed25519.PrivateKey
	)

	currPrivateKeyVersion := "1"

	if privateKeyFile != "" && privateKeySecretName == "" {
		if ecPrivateKey, err = LoadPrivateKey(privateKeyFile); err != nil {
			return nil, err
		}
	}

	if publicKeyFile != "" && publicKeySecretName == "" {
		publicKeyMap, err = InitPublicKeyMapFromFile(publicKeyFile, currPrivateKeyVersion)
	}

	if err != nil {
		return nil, err
	}

	jm := &Manager{
		alg:                   kJwt.EdDSA,
		publicKeyMap:          publicKeyMap,
		logger:                logger,
		privateKeySecretName:  privateKeySecretName,
		currPrivateKeyVersion: currPrivateKeyVersion,
		publicKeySecretName:   publicKeySecretName,
		issuer:                JWTCfg.Issuer,
		privateKey:            ecPrivateKey,
		accessTokenDuration:   JWTCfg.AccessTokenDuration,
		refreshTokenDuration:  JWTCfg.RefreshTokenDuration,
	}

	return jm, nil
}

func (jm *Manager) GenerateAccessToken(
	id string, sessionID, deviceID, domain string, roles []string, kps ...string,
) (string, error) {
	return jm.Generate(id, sessionID, deviceID, domain, roles, jm.accessTokenDuration, kps...)
}

func (jm *Manager) GenerateRefreshToken(
	id string, sessionID, deviceID, domain string, roles []string, kps ...string,
) (string, error) {
	return jm.Generate(id, sessionID, deviceID, domain, roles, jm.refreshTokenDuration, kps...)
}

func (jm *Manager) Generate(
	id string, sessionID, deviceID, domain string, roles []string, duration time.Duration, kps ...string,
) (string, error) {
	identity := base2.Identity{
		ID: id, Roles: roles, Metadata: map[string]string{}, SessionID: sessionID, Domain: domain,
		DeviceID: deviceID,
	}

	nKeyPair := len(kps) / 2
	for i := 0; i < nKeyPair; i++ {
		keyIdx := i * 2
		identity.Metadata[kps[keyIdx]] = kps[keyIdx+1]
	}

	claims := Claims{
		Claims: kJwt.Claims{
			Issuer: jm.issuer,
		},
		Identity: identity,
	}

	token, err := kJwt.SignWithHeader(
		jm.alg,
		jm.privateKey,
		claims,
		NewHeader(jm.alg.Name(), "JWT", jm.currPrivateKeyVersion),
		kJwt.MaxAge(duration),
	)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func (jm *Manager) ParseUnverifiedClaims(accessToken []byte) (*base2.Identity, error) {
	unVerifiedToken, err := kJwt.Decode(accessToken)
	if err != nil {
		return nil, err
	}

	var claims base2.Identity
	err = unVerifiedToken.Claims(&claims)
	if err != nil {
		return nil, fmt.Errorf("invalid token claims")
	}

	return &claims, nil
}

func (jm *Manager) Verify(accessToken []byte) (*base2.Identity, int64, error) {
	verifiedToken, err := kJwt.VerifyWithHeaderValidator(jm.alg, nil, accessToken, jm.validateHeader)
	if err != nil {
		return nil, 0, err
	}

	var claims Claims
	err = verifiedToken.Claims(&claims)

	standardClaims := claims.Claims

	switch {
	case err != nil:
		return nil, 0, fmt.Errorf("invalid token claims")
	case standardClaims.Issuer != jm.issuer:
		return nil, 0, fmt.Errorf("invalid issuer")
	default:
		return &claims.Identity, standardClaims.Expiry, nil
	}
}

func (jm *Manager) validateHeader(alg string, headerDecoded []byte) (kJwt.Alg, kJwt.PublicKey, kJwt.InjectFunc, error) { //nolint:lll
	var h Header
	err := kJwt.Unmarshal(headerDecoded, &h)
	if err != nil {
		return nil, nil, nil, err
	}

	if h.Alg != alg {
		return nil, nil, nil, kJwt.ErrTokenAlg
	}

	kid := h.Kid
	if h.Kid == "" {
		return nil, nil, nil, fmt.Errorf("kid is empty")
	}

	publicKey, err := jm.loadPublicKey(kid)
	if err != nil {
		return nil, nil, nil, err
	}

	return jm.alg, publicKey, nil, nil
}

func (jm *Manager) loadPublicKey(version string) (ed25519.PublicKey, error) {
	publicKeyMap := jm.publicKeyMap

	publicKey, existed := publicKeyMap.Get(version)
	if existed {
		return publicKey, nil
	}

	return publicKey, fmt.Errorf("public key not found")
}

func (jm *Manager) IsExpiredError(err error) bool {
	return errors.Is(err, kJwt.ErrExpired)
}

func (jm *Manager) Stop() {}
