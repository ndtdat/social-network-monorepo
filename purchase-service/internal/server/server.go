package server

import (
	"context"
	"fmt"
	serviceconfig "github.com/ndtdat/social-network-monorepo/purchase-service/config"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/client"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/subscriptionplan"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/voucher"

	pb "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase"
	"go.uber.org/zap"
)

type PurchaseServer struct {
	name          string
	logger        *zap.Logger
	serviceCfg    *serviceconfig.Service
	microservices *client.MicroservicesManager

	voucherService          *voucher.Service
	subscriptionPlanService *subscriptionplan.Service

	pb.UnimplementedPurchaseServer
}

func NewServer(
	logger *zap.Logger, serviceCfg *serviceconfig.Service, microservices *client.MicroservicesManager,
	voucherService *voucher.Service, subscriptionPlanService *subscriptionplan.Service,
) *PurchaseServer {
	return &PurchaseServer{
		name:                    "Purchase",
		logger:                  logger,
		serviceCfg:              serviceCfg,
		microservices:           microservices,
		voucherService:          voucherService,
		subscriptionPlanService: subscriptionPlanService,
	}
}

func (p *PurchaseServer) init(ctx context.Context) error {
	if err := p.voucherService.Init(ctx); err != nil {
		panic(fmt.Sprintf("init voucher failed, due to %v", err))
	}

	if err := p.subscriptionPlanService.Init(ctx); err != nil {
		panic(fmt.Sprintf("init subscription plan failed, due to %v", err))
	}

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
