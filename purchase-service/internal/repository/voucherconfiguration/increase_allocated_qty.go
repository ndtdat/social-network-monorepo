package voucherconfiguration

import (
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	pbmodel "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
	"gorm.io/gorm"
)

func (r *Repository) IncreaseAllocatedQty(tx *gorm.DB, id uint64) error {
	return tx.Model(&model.VoucherConfiguration{}).
		Where(sqlutil.EqualClause(model.VoucherConfiguration_ID), id).
		Updates(map[string]any{
			model.VoucherConfiguration_ALLOCATED_QTY: gorm.Expr(
				fmt.Sprintf("%s + 1", model.VoucherConfiguration_ALLOCATED_QTY),
			),
		}).Update(model.VoucherConfiguration_STATUS,
		gorm.Expr(
			fmt.Sprintf(`
			IF (%[1]s >= %[2]s , ?, %[3]s)
		`,
				model.VoucherConfiguration_ALLOCATED_QTY,
				model.VoucherConfiguration_MAX_QTY,
				model.VoucherConfiguration_STATUS,
			),
			pbmodel.VoucherStatus_VS_UNAVAILABLE,
		)).Error
}
