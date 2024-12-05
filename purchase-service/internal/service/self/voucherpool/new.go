package voucherpool

import (
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/uservoucher"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/voucherpool"
)

type Service struct {
	repo            *voucherpool.Repository
	userVoucherRepo *uservoucher.Repository
}

func NewService(repo *voucherpool.Repository, userVoucherRepo *uservoucher.Repository) *Service {
	return &Service{
		repo:            repo,
		userVoucherRepo: userVoucherRepo,
	}
}
