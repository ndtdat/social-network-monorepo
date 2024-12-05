package uservoucher

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"gorm.io/gorm"
)

func (r *Repository) buildRawQueryForGetAllocatedVoucherCodes(codes []string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if len(codes) != 0 {
			tx = tx.Where(sqlutil.InClause(model.UserVoucher_VOUCHER_CODE), codes)
		}

		return tx
	}
}

func (r *Repository) GetAllocatedVoucherCodes(
	ctx context.Context, codes []string, countRow bool,
) ([]string, int64, error) {
	return r.GetAllocatedVoucherCodesWithTx(r.db.WithContext(ctx), codes, countRow)
}

func (r *Repository) GetAllocatedVoucherCodesWithTx(
	tx *gorm.DB, codes []string, countRow bool,
) ([]string, int64, error) {
	var (
		results []string
		total   int64
	)

	tx = tx.Model(&model.UserVoucher{}).Scopes(r.buildRawQueryForGetAllocatedVoucherCodes(codes))

	if countRow {
		tx = tx.Count(&total)
	}

	return results, total, tx.Pluck(model.UserVoucher_VOUCHER_CODE, &results).Error
}
