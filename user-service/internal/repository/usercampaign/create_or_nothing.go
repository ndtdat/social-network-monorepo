package usercampaign

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *Repository) CreateOrNothing(ctx context.Context, params *model.UserCampaign) error {
	return r.CreateOrNothingWithTx(r.db.WithContext(ctx), params)
}

func (r *Repository) CreateOrNothingWithTx(tx *gorm.DB, params *model.UserCampaign) error {
	return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&params).Error
}
