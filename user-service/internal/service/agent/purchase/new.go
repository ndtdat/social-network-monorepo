package purchase

import (
	"github.com/ndtdat/social-network-monorepo/user-service/internal/client"
	"go.uber.org/zap"
)

type Service struct {
	logger        *zap.Logger
	microservices *client.MicroservicesManager
}

func NewService(
	logger *zap.Logger, microservices *client.MicroservicesManager,
) *Service {
	return &Service{
		logger:        logger,
		microservices: microservices,
	}
}
