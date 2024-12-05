package usertier

import (
	"context"
	"errors"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"gorm.io/gorm"
)

func (r *Repository) FirstByFilters(ctx context.Context, userID uint64) (*model.UserTier, error) {
	var result *model.UserTier

	tx := r.db.WithContext(ctx).Model(&model.UserTier{})

	tx = tx.Where(sqlutil.EqualClause(model.UserTier_USER_ID), userID)

	if err := tx.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}
