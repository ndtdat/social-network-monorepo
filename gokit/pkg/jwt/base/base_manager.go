package base

import (
	"time"
)

type Manager interface {
	Generate(
		ID string, sessionID, deviceID, domain string, roles []string, duration time.Duration, kps ...string,
	) (string, error)

	GenerateAccessToken(ID string, sessionID, deviceID, domain string, roles []string, kps ...string) (string, error)

	GenerateRefreshToken(ID string, sessionID, deviceID, domain string, roles []string, kps ...string) (string, error)

	ParseUnverifiedClaims(accessToken []byte) (*Identity, error)

	Verify(accessToken []byte) (*Identity, int64, error) // int64 = expiry

	IsExpiredError(err error) bool

	Stop()
}
