package campaign

import (
	"github.com/ndtdat/social-network-monorepo/user-service/config"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/repository/campaign"
)

type Service struct {
	serviceCfg *config.Service
	repo       *campaign.Repository
}

func NewService(serviceCfg *config.Service, repo *campaign.Repository) *Service {
	return &Service{serviceCfg: serviceCfg, repo: repo}
}
