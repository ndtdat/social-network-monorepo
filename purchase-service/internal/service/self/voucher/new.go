package voucher

import (
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/uservoucher"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/voucherconfiguration"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/voucherpool"
	"gorm.io/gorm"
)

type Service struct {
	db                *gorm.DB
	configurationRepo *voucherconfiguration.Repository
	voucherPoolRepo   *voucherpool.Repository
	userVoucherRepo   *uservoucher.Repository
}

func NewService(
	db *gorm.DB, configurationRepo *voucherconfiguration.Repository, voucherPoolRepo *voucherpool.Repository,
	userVoucherRepo *uservoucher.Repository,
) *Service {
	return &Service{
		db:                db,
		configurationRepo: configurationRepo,
		voucherPoolRepo:   voucherPoolRepo,
		userVoucherRepo:   userVoucherRepo,
	}
}
