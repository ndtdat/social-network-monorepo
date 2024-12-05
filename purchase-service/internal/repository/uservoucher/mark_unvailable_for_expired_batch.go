package uservoucher

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	pbmodel "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
)

func (r *Repository) MarkUnavailableForExpiredBatch(ctx context.Context) error {
	return r.db.WithContext(ctx).Model(&model.UserVoucher{}).
		Where(sqlutil.EqualClause(model.UserVoucher_STATUS), pbmodel.UserVoucherStatus_UVS_ALLOCATED).
		Where(sqlutil.LessThanClause(model.UserVoucher_EXPIRED_AT), util.CurrentUnix()).
		Updates(map[string]any{
			model.UserVoucher_STATUS: pbmodel.UserVoucherStatus_UVS_EXPIRED,
		}).Error
}
