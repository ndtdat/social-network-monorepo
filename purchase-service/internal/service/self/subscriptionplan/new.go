package subscriptionplan

import (
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/subscriptionplan"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/transaction"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/usertier"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/uservoucher"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/voucher"
	"gorm.io/gorm"
)

type Service struct {
	db                   *gorm.DB
	userTierRepo         *usertier.Repository
	subscriptionPlanRepo *subscriptionplan.Repository
	userVoucherRepo      *uservoucher.Repository
	transactionRepo      *transaction.Repository
	voucherService       *voucher.Service
}

func NewService(
	db *gorm.DB, userTierRepo *usertier.Repository, subscriptionPlanRepo *subscriptionplan.Repository,
	userVoucherRepo *uservoucher.Repository, voucherService *voucher.Service, transactionRepo *transaction.Repository,
) *Service {
	return &Service{
		db:                   db,
		userTierRepo:         userTierRepo,
		subscriptionPlanRepo: subscriptionPlanRepo,
		userVoucherRepo:      userVoucherRepo,
		voucherService:       voucherService,
		transactionRepo:      transactionRepo,
	}
}
