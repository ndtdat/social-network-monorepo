package util

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/common"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

var defaultAcceptLanguage = "en-US,en;q=0.9"

func SetIdentityClaimsIntoHeader(ctx context.Context, stream grpc.ServerStream, claims *base.Identity) error {
	inMd, ok := metadata.FromIncomingContext(ctx)
	var newInMd metadata.MD

	if !ok {
		newInMd = metadata.New(nil)
	} else {
		newInMd = inMd.Copy()
	}

	newInMd.Set(common.IdentityIDHeader, claims.ID)
	newInMd.Set(common.IdentityRolesHeader, EncodeIdentityRolesHeader(claims.Roles))
	newInMd.Set(common.IdentityMetadataHeader, EncodeIdentityMetadataHeader(claims.Metadata))
	newInMd.Set(common.IdentityAPIKeyHeader, EncodeIdentityAPIKeyHeader(claims.APIKey))

	return stream.SetHeader(newInMd)
}

func SetIdentityClaims(ctx context.Context, claims *base.Identity) context.Context {
	inMd, ok := metadata.FromIncomingContext(ctx)
	var newInMd metadata.MD

	if !ok {
		newInMd = metadata.New(nil)
	} else {
		newInMd = inMd.Copy()
	}

	newInMd.Set(common.IdentityIDHeader, claims.ID)
	newInMd.Set(common.IdentityRolesHeader, EncodeIdentityRolesHeader(claims.Roles))
	newInMd.Set(common.IdentityMetadataHeader, EncodeIdentityMetadataHeader(claims.Metadata))
	newInMd.Set(common.IdentityAPIKeyHeader, EncodeIdentityAPIKeyHeader(claims.APIKey))

	return metadata.NewIncomingContext(ctx, newInMd)
}

func IdentityClaimsFromCtx(ctx context.Context) *base.Identity {
	inMd, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	return IdentityClaimsFromMD(inMd)
}

func AcceptLanguageFromCtx(ctx context.Context) string {
	language := FieldFromIncomingCtx(ctx, common.AcceptLanguage)
	if language == "" {
		return defaultAcceptLanguage
	}

	return language
}

func XForwardedForFromCtx(ctx context.Context) string {
	return FieldFromIncomingCtx(ctx, common.XForwardedForHeader)
}

func ClientIPFromCtx(ctx context.Context) string {
	return FieldFromIncomingCtx(ctx, common.ClientIPHeader)
}

func DomainFromCtx(ctx context.Context) string {
	return FieldFromIncomingCtx(ctx, common.DomainHeader)
}

func DeviceIDFromCtx(ctx context.Context) string {
	return FieldFromIncomingCtx(ctx, common.DeviceIDHeader)
}

func SessionIDFromCtx(ctx context.Context) string {
	return FieldFromIncomingCtx(ctx, common.SessionIDHeader)
}

func UserAgentFromCtx(ctx context.Context) string {
	return FieldFromIncomingCtx(ctx, common.UserAgentHeader)
}

func IdentityClaimsFromMD(inMd metadata.MD) *base.Identity {
	if inMd == nil {
		return nil
	}

	claims := base.Identity{}
	claimID := FieldFromMD(inMd, common.IdentityIDHeader)
	//if claimID == "" {
	//	return nil
	//}

	// TODO required this later
	sessionID := FieldFromMD(inMd, common.SessionIDHeader)
	//if sessionID == "" {
	//	return nil
	//}

	deviceID := FieldFromMD(inMd, common.DeviceIDHeader)
	//if deviceID == "" {
	//	return nil
	//}

	claims.ID = claimID
	claims.SessionID = sessionID
	claims.Roles = DecodeIdentityRolesHeader(FieldFromMD(inMd, common.IdentityRolesHeader))
	claims.Metadata = DecodeIdentityMetadataHeader(FieldFromMD(inMd, common.IdentityMetadataHeader))
	claims.APIKey = DecodeIdentityAPIKeyHeader(FieldFromMD(inMd, common.IdentityAPIKeyHeader))
	claims.DeviceID = deviceID

	return &claims
}

func IsInternalCall(ctx context.Context) bool {
	return FieldFromIncomingCtx(ctx, common.InternalCallHeader) == "true"
}

func EncodeIdentityRolesHeader(roles []string) string {
	return strings.Join(roles, common.HeaderSeparator)
}

func DecodeIdentityRolesHeader(header string) []string {
	return strings.Split(header, common.HeaderSeparator)
}

func EncodeIdentityMetadataHeader(metadata map[string]string) string {
	return strings.Join(MapToPairs(metadata), common.HeaderSeparator)
}

func DecodeIdentityMetadataHeader(header string) map[string]string {
	return PairsToMap(strings.Split(header, common.HeaderSeparator))
}

func EncodeIdentityAPIKeyHeader(apiKey map[string]string) string {
	return strings.Join(MapToPairs(apiKey), common.HeaderSeparator)
}

func DecodeIdentityAPIKeyHeader(apiKey string) map[string]string {
	return PairsToMap(strings.Split(apiKey, common.HeaderSeparator))
}
