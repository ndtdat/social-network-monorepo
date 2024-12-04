package authz

import (
	"context"
	"github.com/dolthub/swiss"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/common"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	base2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/base"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/richererror"
	util2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	jwtManager      base2.Manager
	accessibleRoles *swiss.Map[string, []string]
	apiKeyHeader    *swiss.Map[string, string]
}

func NewInterceptor(jwtManager base2.Manager, authConfig *config.Auth) (*Interceptor, error) {
	return &Interceptor{
		jwtManager,
		authConfig.AuthRouteMap(),
		authConfig.APIKeyHeaderMap(),
	}, nil
}

func (ai *Interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		if _, err := ai.authorize(ctx, info.FullMethod); err != nil {
			return nil, richererror.GRPCWebIOSErrorWrapper(ctx, err)
		}

		return handler(ctx, req)
	}
}

func (ai *Interceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv any,
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := stream.Context()

		if _, err := ai.authorize(ctx, info.FullMethod); err != nil {
			return richererror.GRPCWebIOSErrorWrapper(ctx, err)
		}

		return handler(srv, stream)
	}
}

func (ai *Interceptor) authorize(ctx context.Context, method string) (*base2.Identity, error) {
	accessibleRoles, exist := ai.accessibleRoles.Get(method)
	md, hasMD := metadata.FromIncomingContext(ctx)

	if isPublicMethod := !exist || len(accessibleRoles) == 0; isPublicMethod {
		return nil, nil //nolint:nilnil
	}

	if !hasMD {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	claims := util2.IdentityClaimsFromMD(md)
	internalCall := util2.FieldFromMD(md, common.InternalCallHeader)
	if internalCall == "true" {
		return ai.checkPermission(claims, accessibleRoles, true)
	}

	if claims == nil || claims.ID == "" {
		return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")
	}

	return ai.checkPermission(claims, accessibleRoles, false)
}

func (ai *Interceptor) checkPermission(
	identityClaims *base2.Identity, accessibleRoles []string, isSystemCall bool,
) (*base2.Identity, error) {
	if isSystemCall {
		return identityClaims, nil
	}

	for _, role := range accessibleRoles {
		for _, claimedRole := range identityClaims.Roles {
			if role == claimedRole {
				return identityClaims, nil
			}
		}
	}

	return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
