package usertier

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *Repository) Upsert(ctx context.Context, params *model.UserTier) error {
	return r.UpsertWithTx(r.db.WithContext(ctx), params)
}

func (r *Repository) UpsertWithTx(tx *gorm.DB, params *model.UserTier) error {
	return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&params).Error
}
