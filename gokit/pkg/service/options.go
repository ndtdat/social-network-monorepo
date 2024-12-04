package service

import (
	grpcclient "github.com/ndtdat/social-network-monorepo/gokit/pkg/client/grpc"
	config2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/cronjob"
	jwtbase "github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/base"
	"github.com/redis/rueidis"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/xds"
	"gorm.io/gorm"
)

type Option func(o *Options)

type Options struct {
	CronjobManager    cronjob.Manager
	GRPCClientManager grpcclient.ClientManager
	Rueidis           rueidis.Client
	JWTManager        jwtbase.Manager
	GRPCServer        *grpc.Server
	Logger            *zap.Logger
	MySQL             *gorm.DB
	XDSGRPCServer     *xds.GRPCServer
	HealthServer      *health.Server
	AuthConfig        *config2.Auth
	AppConfig         *config2.App
	Name              string
}

func HealthServer(h *health.Server) Option {
	return func(o *Options) {
		o.HealthServer = h
	}
}

func AppConfig(cf *config2.App) Option {
	return func(o *Options) {
		o.AppConfig = cf
	}
}

func AuthConfig(cf *config2.Auth) Option {
	return func(o *Options) {
		o.AuthConfig = cf
	}
}

func Logger(logger *zap.Logger) Option {
	return func(o *Options) {
		o.Logger = logger
	}
}

func JWTManager(jm jwtbase.Manager) Option {
	return func(o *Options) {
		o.JWTManager = jm
	}
}

func MySQL(db *gorm.DB) Option {
	return func(o *Options) {
		o.MySQL = db
	}
}

func GRPCServer(gs *grpc.Server) Option {
	return func(o *Options) {
		o.GRPCServer = gs
	}
}

func XDSGRPCServer(gs *xds.GRPCServer) Option {
	return func(o *Options) {
		o.XDSGRPCServer = gs
	}
}

func Rueidis(rc rueidis.Client) Option {
	return func(o *Options) {
		o.Rueidis = rc
	}
}

func GRPCClientManager(gcm grpcclient.ClientManager) Option {
	return func(o *Options) {
		o.GRPCClientManager = gcm
	}
}
func CronjobManager(cm cronjob.Manager) Option {
	return func(o *Options) {
		o.CronjobManager = cm
	}
}

func NewOptions() *Options {
	return &Options{}
}
