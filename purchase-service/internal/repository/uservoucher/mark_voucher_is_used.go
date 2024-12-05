package uservoucher

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	model2 "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
	"gorm.io/gorm"
)

func (r *Repository) MarkVoucherIsUsed(tx *gorm.DB, id uint64) error {
	return tx.Model(&model.UserVoucher{}).
		Where(sqlutil.EqualClause(model.UserVoucher_ID), id).
		Updates(map[string]any{
			model.UserVoucher_STATUS: model2.UserVoucherStatus_UVS_USED,
		}).Error
}
