package voucherconfiguration

import (
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"gorm.io/gorm"
)

func (r *Repository) IncreaseRedeemedQty(tx *gorm.DB, id uint64) error {
	return tx.Model(&model.VoucherConfiguration{}).
		Where(sqlutil.EqualClause(model.VoucherConfiguration_ID), id).
		Updates(map[string]any{
			model.VoucherConfiguration_REDEEMED_QTY: gorm.Expr(
				fmt.Sprintf("%s + 1", model.VoucherConfiguration_REDEEMED_QTY),
			),
		}).Error
}
