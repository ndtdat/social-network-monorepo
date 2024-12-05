package voucherpool

import (
	"errors"
	"github.com/ndtdat/social-network-monorepo/common/pkg/sorter"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *Repository) SelectFirstForUpdate(tx *gorm.DB, sorters []*sorter.Sorter) (string, error) {
	var (
		result *model.VoucherPool
	)

	tx = tx.Model(&model.VoucherPool{}).
		Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate, Options: clause.LockingOptionsSkipLocked})

	for _, sorter := range sorters {
		tx = tx.Order(sorter.GetExp())
	}

	if err := tx.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}

		return "", err
	}

	return result.Code, nil
}
