package authn

import (
	"context"
	"github.com/dolthub/swiss"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/common"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	base2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/base"
	richererror2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/richererror"
	util2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
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
		claims, err := ai.authenticate(ctx, info.FullMethod)
		if err != nil {
			return nil, richererror2.GRPCWebIOSErrorWrapper(ctx, err)
		}
		if claims != nil {
			ctx = util2.SetIdentityClaims(ctx, claims)
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

		if claims, err := ai.authenticate(ctx, info.FullMethod); err != nil {
			return richererror2.GRPCWebIOSErrorWrapper(ctx, err)
		} else if claims != nil {
			ctx = util2.SetIdentityClaims(ctx, claims)
		}

		return handler(srv, &grpc_middleware.WrappedServerStream{
			ServerStream:   stream,
			WrappedContext: ctx,
		})
	}
}

func (ai *Interceptor) authenticate(ctx context.Context, method string) (*base2.Identity, error) {
	accessibleRoles, exist := ai.accessibleRoles.Get(method)
	if isPublicMethod := !exist; isPublicMethod {
		return nil, nil //nolint:nilnil
	}

	md, hasMD := metadata.FromIncomingContext(ctx)
	if !hasMD {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	if len(accessibleRoles) == 0 {
		apiKeyHeader, exist := ai.apiKeyHeader.Get(method)
		if exist {
			requestClaims := &base2.Identity{}
			apiKeyValue := util2.FieldFromMD(md, apiKeyHeader)
			requestClaims.APIKey = map[string]string{}
			requestClaims.APIKey[apiKeyHeader] = apiKeyValue

			return requestClaims, nil
		}

		return nil, nil
	}

	internalCall := util2.FieldFromMD(md, common.InternalCallHeader)
	if internalCall == "true" {
		return util2.IdentityClaimsFromMD(md), nil
	}

	var accessToken []byte
	value := util2.FieldFromMD(md, common.AuthorizationHeader)
	if value != "" {
		accessToken = []byte(strings.ReplaceAll(value, "Bearer ", ""))
	} else {
		return nil, richererror2.Unauthenticated("authorization token is not provided")
	}

	jm := ai.jwtManager
	requestClaims, _, err := jm.Verify(accessToken)
	if err != nil {
		if jm.IsExpiredError(err) {
			return nil, richererror2.NewRicherCode(
				codes.Unauthenticated, err.Error(),
				richererror2.NewErrorDetail(richererror2.ErrorAccessTokenExpired.ToInt(), err.Error()),
			)
		}

		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return requestClaims, nil
}
