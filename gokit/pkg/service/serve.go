package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/common"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/enum"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/xds"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net"
	"net/http"
	"strings"
	//noline:revive
	_ "google.golang.org/grpc/resolver"
	_ "google.golang.org/grpc/xds"
)

func (i *Impl) serveGRPC() {
	logger := i.Logger()
	appConfig := i.AppConfig()

	cfg := appConfig.GRPC
	grpcDomain := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	lis, err := net.Listen("tcp", grpcDomain)
	if lis == nil {
		logger.Fatal(fmt.Sprintf("Failed to listen to %s (it seems like the port is occupied)", grpcDomain))
	}
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to listen to %s", lis.Addr()), zap.Error(err))
	}

	i.GRPCNetListener = lis
	xDSEnabled := cfg.XDS.Enabled
	logger.Info(fmt.Sprintf("gRPC server listening on %v with xDS=%v", lis.Addr(), xDSEnabled))

	go func() {
		if xDSEnabled {
			err = i.XDSGRPCServer().Serve(lis)
		} else {
			err = i.GRPCServer().Serve(lis)
		}

		if err != nil && i.state == Starting {
			logger.Warn("Error when starting gRPC", zap.Error(err))
		}
	}()
}

func serverHandlersContainAllNil(handlers []ServerHandler) bool {
	for _, h := range handlers {
		if h != nil {
			return false
		}
	}

	return true
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)

			return
		}

		c.Next()
	}
}

func (i *Impl) serveGRPCGateway(ctx context.Context, serverHandlers ...ServerHandler) {
	logger := i.Logger()
	appConfig := i.AppConfig()

	gRPCServerCfg := appConfig.GRPC
	gRPCPort := gRPCServerCfg.Port
	grpcDomain := net.JoinHostPort(gRPCServerCfg.Host, gRPCServerCfg.Port)

	httpServerCfg := appConfig.HTTP
	httpPort := httpServerCfg.Port
	httpDomain := net.JoinHostPort(httpServerCfg.Host, httpPort)

	disableGRPCGateway := httpServerCfg.DisableGRPCGateway
	tracingCfg := appConfig.Tracing
	tracingEnabled := tracingCfg.Enabled

	env := appConfig.Environment
	switch env {
	case enum.Environment_PRODUCTION:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	if tracingEnabled {
		router.Use(
			gintrace.Middleware(tracingCfg.ServiceName, gintrace.WithIgnoreRequest(func(c *gin.Context) bool {
				return util.StringInSlice(c.Request.URL.Path, ignoredTracingPaths)
			})),
		)
	}

	if len(serverHandlers) == 0 || serverHandlersContainAllNil(serverHandlers) {
		disableGRPCGateway = true
	}

	//nolint:nestif
	if !disableGRPCGateway {
		creds, _ := xds.NewClientCredentials(xds.ClientOptions{
			FallbackCreds: insecure.NewCredentials(),
		})

		gateWayDomain := grpcDomain
		xdsCfg := gRPCServerCfg.XDS
		if xdsCfg.Enabled {
			gateWayDomain = fmt.Sprintf("%s:%s", xdsCfg.Host, gRPCPort)
		}

		conn, err := grpc.DialContext(
			ctx,
			gateWayDomain,
			grpc.WithBlock(),
			grpc.WithTransportCredentials(creds),
		)
		if err != nil {
			logger.Fatal(fmt.Sprintf("Failed to dial gRPC server %s", grpcDomain), zap.Error(err))
		}
		i.GRPCGatewayClientConn = conn

		grpcGatewayMux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(CustomMatcher))

		for _, sh := range serverHandlers {
			if sh == nil {
				continue
			}

			if err = sh(ctx, grpcGatewayMux, conn); err != nil {
				logger.Fatal("Failed to register gateway:", zap.Error(err))
			}
		}

		mux := http.NewServeMux()
		mux.Handle("/", grpcGatewayMux)

		path := httpServerCfg.Path
		if path == "" {
			logger.Fatal("Path for HTTP server is empty")
		}
		relativePath := fmt.Sprintf("%s/*any", httpServerCfg.Path)
		router.Any(relativePath, gin.WrapF(mux.ServeHTTP))
	}

	if i.httpRouterHandlers != nil {
		for _, h := range i.httpRouterHandlers {
			h(router)
		}
	}

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.POST("/ping", gin.WrapF(i.ping))

	logger.Info(fmt.Sprintf("HTTP server listening on %v with GRPCGateway=%v", httpDomain, !disableGRPCGateway))
	go func() {
		if err := router.Run(httpDomain); err != nil && i.state == Starting {
			logger.Fatal("Error when starting HTTP server", zap.Error(err))
		}
	}()
}

func (i *Impl) startTracing() {
	logger := i.Logger()
	appConfig := i.AppConfig()

	if tracingCfg := appConfig.Tracing; tracingCfg.Enabled {
		logger.Info("Tracer starts")
		tracer.Start(
			tracer.WithUDS(tracingCfg.SocketPath),
			tracer.WithServiceName(tracingCfg.ServiceName),
		)
	}
}

func (i *Impl) autoGOMAXPROCS() {
	logger := i.Logger()

	if _, err := maxprocs.Set(maxprocs.Logger(func(msg string, args ...any) {
		logger.Info(fmt.Sprintf(msg, args))
	})); err != nil {
		logger.Error(fmt.Sprintf("Error when auto setting GOMAXPROCS: %v", err))
	}
}

func (i *Impl) Serve(ctx context.Context, serverHandlers ...ServerHandler) {
	ctx, cancel := context.WithCancel(ctx)
	i.ctxCancelFunc = cancel

	i.state = Starting

	i.autoGOMAXPROCS()
	i.initGRPCClientManager()
	i.startTracing()
	i.serveGRPC()
	i.serveGRPCGateway(ctx, serverHandlers...)
	i.startCronjobManager()

	i.state = Started
}

func (i *Impl) initGRPCClientManager() {
	if gcm := i.GRPCClientManager(); gcm != nil {
		i.Logger().Info("Init gRPC client manager")
		gcm.Init()
	}
}

func (i *Impl) startCronjobManager() {
	cm := i.CronjobManager()
	logger := i.Logger()

	if cm != nil {
		logger.Info("Start cronjob manager")
		if err := cm.Start(); err != nil {
			logger.Error(fmt.Sprintf("Cannot start cron manager due to %v", err))
		}
	}
}

func (i *Impl) ping(_ http.ResponseWriter, _ *http.Request) {
	// No need implementation here
}

//nolint:gocritic
//func (svc *ServiceImpl) serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
//	logger := svc.Logger()
//	appConfig := svc.AppConfig()
//
//	logger.Info("Start serveSwaggerFile")
//
//	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
//		logger.Debug(fmt.Sprintf("Not Found: %s", r.URL.Path))
//		http.NotFound(w, r)
//
//		return
//	}
//
//	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
//	p = path.Join(appConfig.Swagger.Path, p)
//
//	logger.Info(fmt.Sprintf("Serving swagger-file: %s", p))
//
//	http.ServeFile(w, r, p)
//}

func CustomMatcher(key string) (string, bool) {
	switch strings.ToLower(key) {
	case common.IdentityIDHeader:
		return key, true
	case common.IdentityRolesHeader:
		return key, true
	case common.IdentityMetadataHeader:
		return key, true
	case common.AcceptLanguage:
		return key, true
	case common.SessionIDHeader:
		return key, true
	case common.ClientIPHeader:
		return key, true
	case common.XForwardedForHeader:
		return key, true
	case common.DomainHeader:
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
