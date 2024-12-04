package service

import (
	"context"
	grpcclient "github.com/ndtdat/social-network-monorepo/gokit/pkg/client/grpc"
	config2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	jwtbase "github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/base"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/cronjob"
	"github.com/redis/rueidis"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/xds"
	"gorm.io/gorm"
)

type ServerHandler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

type httpRouterHandler func(router *gin.Engine)

type internalServer interface {
	Name() string
	Close() error
	Finalize() error
}

type State uint

const (
	Starting State = iota
	Started
	Closing
	Closed
)

//nolint:interfacebloat,inamedparam
type Service interface {
	AppConfig() *config2.App
	AuthConfig() *config2.Auth
	Logger() *zap.Logger
	JWTManager() jwtbase.Manager
	Rueidis() rueidis.Client
	MySQL() *gorm.DB
	GRPCServer() *grpc.Server
	Options() *Options
	HealthServer() *health.Server
	RegisterInternalServer(internalServer)
	RegisterHTTPRouterHandlers(...httpRouterHandler)
	GRPCClientManager() grpcclient.ClientManager
	CronjobManager() cronjob.Manager

	Serve(context.Context, ...ServerHandler)
	Close()
}

var ignoredTracingPaths = []string{"/metrics"}

type Impl struct {
	GRPCNetListener net.Listener
	opts            *Options

	ctxCancelFunc         context.CancelFunc
	GRPCGatewayClientConn *grpc.ClientConn

	internalServers []internalServer

	httpRouterHandlers []httpRouterHandler
	state              State
}

func (i *Impl) Logger() *zap.Logger {
	return i.opts.Logger
}

func (i *Impl) AppConfig() *config2.App {
	return i.opts.AppConfig
}

func (i *Impl) AuthConfig() *config2.Auth {
	return i.opts.AuthConfig
}

func (i *Impl) JWTManager() jwtbase.Manager {
	return i.opts.JWTManager
}

func (i *Impl) MySQL() *gorm.DB {
	return i.opts.MySQL
}

func (i *Impl) GRPCServer() *grpc.Server {
	return i.opts.GRPCServer
}

func (i *Impl) XDSGRPCServer() *xds.GRPCServer {
	return i.opts.XDSGRPCServer
}

func (i *Impl) HealthServer() *health.Server {
	return i.opts.HealthServer
}

func (i *Impl) GRPCClientManager() grpcclient.ClientManager {
	return i.opts.GRPCClientManager
}

func (i *Impl) CronjobManager() cronjob.Manager {
	return i.opts.CronjobManager
}

func (i *Impl) Rueidis() rueidis.Client {
	return i.opts.Rueidis
}

func (i *Impl) Options() *Options {
	return i.opts
}

func (i *Impl) Init(opts ...Option) {
	for _, opt := range opts {
		opt(i.opts)
	}
}

func NewService(opts ...Option) Service {
	o := NewOptions()

	for _, opt := range opts {
		opt(o)
	}

	s := &Impl{
		opts: o,
	}

	return s
}

func (i *Impl) RegisterHTTPRouterHandlers(hs ...httpRouterHandler) {
	for _, h := range hs {
		if h != nil {
			i.httpRouterHandlers = append(i.httpRouterHandlers, h)
		}
	}
}

func (i *Impl) RegisterInternalServer(is internalServer) {
	if is != nil {
		i.internalServers = append(i.internalServers, is)
	}
}
