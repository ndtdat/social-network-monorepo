package server

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/richererror"
	"google.golang.org/grpc"
)

func GRPCWebErrorWrapperUnary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		_ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		reply, err := handler(ctx, req)

		return reply, richererror.GRPCWebIOSErrorWrapper(ctx, err)
	}
}
