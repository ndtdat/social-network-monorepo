package server

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/authn"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/authz"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/cache/redissync"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/cache/rueidis"
	grpcclient "github.com/ndtdat/social-network-monorepo/gokit/pkg/client/grpc"
	gokitcfg "github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/db/mysql"
	grpcserver "github.com/ndtdat/social-network-monorepo/gokit/pkg/grpc/server"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt"
	jwtbase "github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/base"
	gokitlog "github.com/ndtdat/social-network-monorepo/gokit/pkg/log"
	gokit "github.com/ndtdat/social-network-monorepo/gokit/pkg/service"
	"github.com/ndtdat/social-network-monorepo/purchase-service/config"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/client"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service"
	pb "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/xds"
	"gorm.io/gorm"
)

func Run() {
	app := fx.New(
		config.Module,
		fx.WithLogger(gokitlog.NewFxEventLogger),
		fx.Provide(context.Background),
		gokitlog.Module,
		authn.Module,
		authz.Module,
		jwt.Module,
		client.Module,
		repository.Module,
		service.Module,
		grpcserver.Module,
		mysql.Module,
		rueidis.Module,
		redissync.Module,
		fx.Provide(grpcclient.NewClient),

		fx.Provide(NewService),
		fx.Provide(NewServer),
		fx.Invoke(registerAppHooks),
	)

	app.Run()
}

func registerAppHooks(
	appContext context.Context,
	cfg *gokitcfg.App,
	lifecycle fx.Lifecycle,
	logger *zap.Logger,
	grpcServer *grpc.Server,
	xdsGRPCServer *xds.GRPCServer,
	service gokit.Service,
	server *PurchaseServer,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				if server == nil {
					logger.Fatal("Cannot create server")
				}

				service.RegisterInternalServer(server)
				if err := server.init(); err != nil {
					return err
				}

				if cfg.GRPC.XDS.Enabled {
					pb.RegisterPurchaseServer(xdsGRPCServer, server)
				} else {
					pb.RegisterPurchaseServer(grpcServer, server)
				}

				go service.Serve(appContext, nil)

				return nil
			},
			OnStop: func(context.Context) error {
				service.Close()

				return nil
			},
		},
	)
}

func NewService(
	appConfig *gokitcfg.App, authConfig *gokitcfg.Auth, logger *zap.Logger,
	jwtManager jwtbase.Manager, grpcServer *grpc.Server, healthServer *health.Server,
	microservices *client.MicroservicesManager, db *gorm.DB,
) gokit.Service {
	return gokit.NewService(
		gokit.Logger(logger),
		gokit.AppConfig(appConfig),
		gokit.AuthConfig(authConfig),
		gokit.JWTManager(jwtManager),
		gokit.GRPCServer(grpcServer),
		gokit.HealthServer(healthServer),
		gokit.GRPCClientManager(microservices),
		gokit.MySQL(db),
	)
}
