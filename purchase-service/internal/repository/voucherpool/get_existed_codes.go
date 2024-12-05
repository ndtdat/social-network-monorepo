package voucherpool

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"gorm.io/gorm"
)

func (r *Repository) buildRawQueryForGetExistedCodes(codes []string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if len(codes) != 0 {
			tx = tx.Where(sqlutil.InClause(model.VoucherPool_CODE), codes)
		}

		return tx
	}
}

func (r *Repository) GetExistedCodes(
	ctx context.Context, codes []string, countRow bool,
) ([]string, int64, error) {
	return r.GetExistedGiftCodesWithTx(r.db.WithContext(ctx), codes, countRow)
}

func (r *Repository) GetExistedGiftCodesWithTx(
	tx *gorm.DB, codes []string, countRow bool,
) ([]string, int64, error) {
	var (
		results []string
		total   int64
	)

	tx = tx.Model(&model.VoucherPool{}).Scopes(
		r.buildRawQueryForGetExistedCodes(codes),
	)

	if countRow {
		tx = tx.Count(&total)
	}

	return results, total, tx.Pluck(model.VoucherPool_CODE, &results).Error
}
