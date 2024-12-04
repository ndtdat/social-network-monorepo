package server

import (
	"context"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/authn"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/authz"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/common"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	skitenum "github.com/ndtdat/social-network-monorepo/gokit/pkg/enum"
	"google.golang.org/grpc/credentials/insecure"
	xdscreds "google.golang.org/grpc/credentials/xds"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/xds"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	"runtime/debug"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

func New(
	cfg *config.App, authnInterceptor *authn.Interceptor, authzInterceptor *authz.Interceptor, logger *zap.Logger,
	logOptions []grpc_zap.Option,
) (*grpc.Server, *xds.GRPCServer, *health.Server, error) {
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandlerContext(func(_ context.Context, p any) error {
			logger.Error(fmt.Sprintf("panic triggered: %v %s", p, string(debug.Stack())))

			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}
	enforcement := keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second,
		PermitWithoutStream: true,
	}

	gRPCServerCfg := cfg.GRPC
	maxConnectionAge := gRPCServerCfg.MaxConnectionAge
	maxConnectionAgeDur, err := time.ParseDuration(maxConnectionAge)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("cannot parse maxConnectionAge [%s] due to %v ", maxConnectionAge, err)
	}

	var unaryInterceptors []grpc.UnaryServerInterceptor
	tracingCfg := cfg.Tracing
	if tracingCfg.Enabled {
		unaryInterceptors = append(unaryInterceptors, grpctrace.UnaryServerInterceptor(
			grpctrace.WithUntracedMethods(common.HealthCheckMethod),
			grpctrace.WithServiceName(tracingCfg.ServiceName),
		))
	}

	unaryInterceptors = append(unaryInterceptors,
		grpc_validator.UnaryServerInterceptor(),
		grpc_ctxtags.UnaryServerInterceptor(
			grpc_ctxtags.WithFieldExtractor(
				grpc_ctxtags.CodeGenRequestFieldExtractor,
			),
		),
		grpc_prometheus.UnaryServerInterceptor,
	)

	if gRPCServerCfg.AuthenticationEnabled {
		unaryInterceptors = append(unaryInterceptors, authnInterceptor.Unary())
	}

	if gRPCServerCfg.AuthorizationEnabled {
		unaryInterceptors = append(unaryInterceptors, authzInterceptor.Unary())
	}

	unaryInterceptors = append(
		unaryInterceptors, []grpc.UnaryServerInterceptor{
			grpc_zap.UnaryServerInterceptor(logger, logOptions...),
			GRPCWebErrorWrapperUnary(),
			grpc_recovery.UnaryServerInterceptor(opts...),
		}...,
	)

	streamInterceptors := []grpc.StreamServerInterceptor{
		grpc_validator.StreamServerInterceptor(),
		grpc_ctxtags.StreamServerInterceptor(
			grpc_ctxtags.WithFieldExtractor(
				grpc_ctxtags.CodeGenRequestFieldExtractor,
			),
		),
		grpc_prometheus.StreamServerInterceptor,
	}

	if gRPCServerCfg.AuthenticationEnabled {
		streamInterceptors = append(streamInterceptors, authnInterceptor.Stream())
	}

	if gRPCServerCfg.AuthorizationEnabled {
		streamInterceptors = append(streamInterceptors, authzInterceptor.Stream())
	}

	streamInterceptors = append(streamInterceptors, []grpc.StreamServerInterceptor{
		grpc_zap.StreamServerInterceptor(logger, logOptions...),
		grpc_recovery.StreamServerInterceptor(opts...),
	}...)

	credentials, err := xdscreds.NewServerCredentials(xdscreds.ServerOptions{FallbackCreds: insecure.NewCredentials()})
	if err != nil {
		return nil, nil, nil, err
	}

	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)),
		grpc.KeepaliveEnforcementPolicy(enforcement),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: maxConnectionAgeDur,
		}),
		grpc.Creds(credentials),
		grpc.MaxConcurrentStreams(gRPCServerCfg.MaxConcurrentStreams),
	}

	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthv1.HealthCheckResponse_SERVING)
	reflectionEnable := isReflectionEnabled(gRPCServerCfg, cfg.Environment)

	if gRPCServerCfg.XDS.Enabled {
		xdsServer, err := xds.NewGRPCServer(serverOptions...)
		if err != nil {
			return nil, nil, nil, err
		}

		healthv1.RegisterHealthServer(xdsServer, healthServer)

		if reflectionEnable {
			reflection.Register(xdsServer)
		}

		return nil, xdsServer, healthServer, nil
	}

	server := grpc.NewServer(serverOptions...)
	if reflectionEnable {
		reflection.Register(server)
	}

	healthv1.RegisterHealthServer(server, healthServer)
	grpc_prometheus.Register(server)

	return server, nil, healthServer, nil
}

func isReflectionEnabled(gRPCServerCfg config.GRPC, env skitenum.Environment) bool {
	return gRPCServerCfg.ReflectionEnabled &&
		(env == skitenum.Environment_LOCAL || env == skitenum.Environment_DEVELOPMENT)
}
