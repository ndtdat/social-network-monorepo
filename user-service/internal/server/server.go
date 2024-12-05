package server

import (
	serviceconfig "github.com/ndtdat/social-network-monorepo/user-service/config"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/client"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/auth"

	pb "github.com/ndtdat/social-network-monorepo/user-service/pkg/api/go/user"
	"go.uber.org/zap"
)

type UserServer struct {
	name          string
	logger        *zap.Logger
	serviceCfg    *serviceconfig.Service
	microservices *client.MicroservicesManager

	authService *auth.Service

	pb.UnimplementedUserServer
}

func NewServer(
	logger *zap.Logger, serviceCfg *serviceconfig.Service, microservices *client.MicroservicesManager,
	authService *auth.Service,
) *UserServer {
	return &UserServer{
		name:          "User",
		logger:        logger,
		serviceCfg:    serviceCfg,
		microservices: microservices,
		authService:   authService,
	}
}

func (u *UserServer) init() error {
	return nil
}

func (u *UserServer) Finalize() error {
	return nil
}

func (u *UserServer) Close() error {
	return nil
}

func (u *UserServer) Name() string {
	return u.name
}
