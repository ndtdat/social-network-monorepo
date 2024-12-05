package voucherpool

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
)

func (r *Repository) GetAvailableTotal(ctx context.Context) (int64, error) {
	var total int64

	return total, r.db.WithContext(ctx).Model(&model.VoucherPool{}).
		Count(&total).Error
}
