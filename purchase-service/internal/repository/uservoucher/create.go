package uservoucher

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"gorm.io/gorm"
)

func (r *Repository) Create(ctx context.Context, params *model.UserVoucher) error {
	return r.CreateWithTx(r.db.WithContext(ctx), params)
}

func (r *Repository) CreateWithTx(tx *gorm.DB, params *model.UserVoucher) error {
	return tx.Model(&model.UserVoucher{}).Create(&params).Error
}
