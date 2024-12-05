package voucherconfiguration

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	pbmodel "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
)

func (r *Repository) MarkUnavailableForExpiredBatch(ctx context.Context) error {
	return r.db.WithContext(ctx).Model(&model.VoucherConfiguration{}).
		Where(sqlutil.EqualClause(model.VoucherConfiguration_STATUS), pbmodel.VoucherStatus_VS_AVAILABLE).
		Where(sqlutil.LessThanClause(model.VoucherConfiguration_END_AT), util.CurrentUnix()).
		Updates(map[string]any{
			model.VoucherConfiguration_STATUS: pbmodel.VoucherStatus_VS_UNAVAILABLE,
		}).Error
}
