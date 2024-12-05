package campaign

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *Repository) CreateBatchOrNothing(ctx context.Context, params []*model.Campaign) error {
	return r.CreateBatchOrNothingWithTx(r.db.WithContext(ctx), params)
}

func (r *Repository) CreateBatchOrNothingWithTx(tx *gorm.DB, params []*model.Campaign) error {
	return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&params).Error
}
