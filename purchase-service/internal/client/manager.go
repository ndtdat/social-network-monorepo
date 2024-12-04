package client

import (
	"context"
	gokitconfig "github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"github.com/ndtdat/social-network-monorepo/purchase-service/config"
	grpcconn "google.golang.org/grpc"
)

type MicroservicesManager struct {
	ctx           context.Context
	appConfig     *gokitconfig.App
	serviceConfig *config.Service
	connections   []*grpcconn.ClientConn
}

func NewMicroservicesManager(
	ctx context.Context,
	appConfig *gokitconfig.App,
	serviceConfig *config.Service,
) *MicroservicesManager {
	return &MicroservicesManager{
		ctx:           ctx,
		appConfig:     appConfig,
		serviceConfig: serviceConfig,
	}
}

func (m *MicroservicesManager) Init() {
	m.preload()
}

func (m *MicroservicesManager) preload() {
	go func() {}()
}

func (m *MicroservicesManager) Close() {
	for _, c := range m.connections {
		c.Close()
	}
}

func (m *MicroservicesManager) GRPCClientManager() {}
