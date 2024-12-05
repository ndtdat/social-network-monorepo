package auth

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/suid"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
)

type Info struct {
	ExternalUserID string
	DeviceID       string
	Roles          []string
	UserID         uint64
	OperatorID     uint64
	TenantID       uint64
	SessionID      uint64
}

func NullableIdentityIDFromCtx(ctx context.Context) uint64 {
	claims := util.IdentityClaimsFromCtx(ctx)
	if claims == nil {
		return 0
	}

	return util.SafeParseUint64WithDefault(claims.ID, 0)
}

func IdentityIDFromCtx(ctx context.Context) uint64 {
	claims := util.IdentityClaimsFromCtx(ctx)
	if claims == nil {
		panic("cannot get identity claims")
	}

	return util.MustParseUint64(claims.ID)
}

func UserAgentFromCtx(ctx context.Context) string {
	return util.UserAgentFromCtx(ctx)
}

func ClaimsFromCtx(ctx context.Context) (uint64, []string) {
	claims := util.IdentityClaimsFromCtx(ctx)
	if claims == nil {
		panic("cannot get identity claims")
	}

	var roles []string
	for _, r := range claims.Roles {
		roles = append(roles, r)
	}

	return util.MustParseUint64(claims.ID), roles
}

func IPFromCtx(ctx context.Context) string {
	return util.ClientIPFromCtx(ctx)
}

func SessionID(ctx context.Context) string {
	return util.SessionIDFromCtx(ctx)
}

func DeviceIDFromCtx(ctx context.Context) string {
	return util.DeviceIDFromCtx(ctx)
}

func IdentityIDInfoFromCtx(ctx context.Context) *Info {
	claims := util.IdentityClaimsFromCtx(ctx)
	if claims == nil {
		panic("cannot get identity claims")
	}

	deviceID := claims.DeviceID
	if deviceID == "" {
		deviceID = util.Uint64ToString(suid.New())
	}

	return &Info{
		UserID: util.MustParseUint64(claims.ID),
		Roles:  claims.Roles,
	}
}
