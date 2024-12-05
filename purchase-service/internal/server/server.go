package server

import (
	serviceconfig "github.com/ndtdat/social-network-monorepo/purchase-service/config"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/client"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/voucher"

	pb "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase"
	"go.uber.org/zap"
)

type PurchaseServer struct {
	name          string
	logger        *zap.Logger
	serviceCfg    *serviceconfig.Service
	microservices *client.MicroservicesManager

	voucherService *voucher.Service

	pb.UnimplementedPurchaseServer
}

func NewServer(
	logger *zap.Logger, serviceCfg *serviceconfig.Service, microservices *client.MicroservicesManager,
	voucherService *voucher.Service,
) *PurchaseServer {
	return &PurchaseServer{
		name:           "Purchase",
		logger:         logger,
		serviceCfg:     serviceCfg,
		microservices:  microservices,
		voucherService: voucherService,
	}
}

func (p *PurchaseServer) init() error {
	return nil
}

func (p *PurchaseServer) Finalize() error {
	return nil
}

func (p *PurchaseServer) Close() error {
	return nil
}

func (p *PurchaseServer) Name() string {
	return p.name
}
