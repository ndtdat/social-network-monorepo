package voucherconfiguration

import (
	"context"
	"errors"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	model2 "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
	"gorm.io/gorm"
)

func (r *Repository) GetAvailableByCampaignID(
	ctx context.Context, campaignID uint64,
) (*model.VoucherConfiguration, error) {
	var result *model.VoucherConfiguration

	tx := r.db.WithContext(ctx).Model(&model.VoucherConfiguration{})

	tx = tx.Where(sqlutil.EqualClause(model.VoucherConfiguration_CAMPAIGN_ID), campaignID)
	tx = tx.Where(sqlutil.EqualClause(model.VoucherConfiguration_STATUS), model2.VoucherStatus_VS_AVAILABLE)

	if err := tx.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}
