package client

import (
	"context"
	"fmt"
	gokitconfig "github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase"
	"github.com/ndtdat/social-network-monorepo/user-service/config"
	grpcconn "google.golang.org/grpc"
)

type MicroservicesManager struct {
	ctx        context.Context
	appCfg     *gokitconfig.App
	serviceCfg *config.Service

	purchaseService purchase.PurchaseClient

	connections []*grpcconn.ClientConn
}

func NewMicroservicesManager(
	ctx context.Context,
	appConfig *gokitconfig.App,
	serviceConfig *config.Service,
) *MicroservicesManager {
	return &MicroservicesManager{
		ctx:        ctx,
		appCfg:     appConfig,
		serviceCfg: serviceConfig,
	}
}

func (m *MicroservicesManager) Init() {
	m.preload()
}

func (m *MicroservicesManager) preload() {
	go func() { _, _ = m.PurchaseClient() }()
}

func (m *MicroservicesManager) Close() {
	for _, c := range m.connections {
		c.Close()
	}
}

func (m *MicroservicesManager) GRPCClientManager() {}

func (m *MicroservicesManager) PurchaseClient() (purchase.PurchaseClient, error) {
	cfg := m.serviceCfg.Microservices.Purchase
	if m.purchaseService == nil {
		purchaseService, conn, err := m.newPurchaseClient(cfg.Host, cfg.Port, cfg.TLS)
		if err != nil {
			return nil, fmt.Errorf("cannot connect to payment service due to %v", err)
		}

		m.connections = append(m.connections, conn)
		m.purchaseService = purchaseService
	}

	return m.purchaseService, nil
}
