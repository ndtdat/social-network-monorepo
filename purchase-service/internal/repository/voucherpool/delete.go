package voucherpool

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"gorm.io/gorm"
)

func (r *Repository) DeleteWithTx(tx *gorm.DB, code string) error {
	return tx.Where(sqlutil.EqualClause(model.VoucherPool_CODE), code).Delete(&model.VoucherPool{}).Error
}

func (r *Repository) DeleteBatchWithTx(tx *gorm.DB, codes []string) error {
	return tx.Where(sqlutil.InClause(model.VoucherPool_CODE), codes).Delete(&model.VoucherPool{}).Error
}
