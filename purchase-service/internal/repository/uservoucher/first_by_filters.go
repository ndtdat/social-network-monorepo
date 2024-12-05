package uservoucher

import (
	"context"
	"errors"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	pbmodel "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
	"gorm.io/gorm"
)

func (r *Repository) FirstByFilters(
	ctx context.Context, userID uint64, subscriptionPlanTier pbmodel.SubscriptionPlanTier,
) (*model.UserVoucher, error) {
	var result *model.UserVoucher

	tx := r.db.WithContext(ctx).Model(&model.UserVoucher{})

	tx = tx.Where(sqlutil.EqualClause(model.UserVoucher_USER_ID), userID)
	tx = tx.Preload(
		model.UserVoucher_PRELOAD_VOUCHER_CONFIGURATION,
		func(db *gorm.DB) *gorm.DB {
			return db.Where(sqlutil.EqualClause(model.VoucherConfiguration_APPLIED_TIER), subscriptionPlanTier)
		})

	if err := tx.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}
