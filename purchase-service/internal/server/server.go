package server

import (
	serviceconfig "github.com/ndtdat/social-network-monorepo/purchase-service/config"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/client"

	pb "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase"
	"go.uber.org/zap"
)

type PurchaseServer struct {
	name          string
	logger        *zap.Logger
	serviceCfg    *serviceconfig.Service
	microservices *client.MicroservicesManager

	pb.UnimplementedPurchaseServer
}

func NewServer(
	logger *zap.Logger, serviceCfg *serviceconfig.Service, microservices *client.MicroservicesManager,
) *PurchaseServer {
	return &PurchaseServer{
		name:          "Purchase",
		logger:        logger,
		serviceCfg:    serviceCfg,
		microservices: microservices,
	}
}

func (u *PurchaseServer) init() error {
	return nil
}

func (u *PurchaseServer) Finalize() error {
	return nil
}

func (u *PurchaseServer) Close() error {
	return nil
}

func (u *PurchaseServer) Name() string {
	return u.name
}
