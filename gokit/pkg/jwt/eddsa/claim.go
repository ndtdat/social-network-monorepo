package eddsa

import (
	kJwt "github.com/kataras/jwt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/base"
)

//nolint:govet
type Claims struct {
	base.Identity
	kJwt.Claims
}
