package voucherpool

import (
	"github.com/ndtdat/social-network-monorepo/purchase-service/config"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/uservoucher"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/voucherpool"
)

type Service struct {
	serviceCfg      *config.Service
	repo            *voucherpool.Repository
	userVoucherRepo *uservoucher.Repository
}

func NewService(
	serviceCfg *config.Service, repo *voucherpool.Repository, userVoucherRepo *uservoucher.Repository,
) *Service {
	return &Service{
		serviceCfg:      serviceCfg,
		repo:            repo,
		userVoucherRepo: userVoucherRepo,
	}
}
