package campaign

import (
	"context"
	"errors"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	"gorm.io/gorm"
)

func (r *Repository) FirstByFilters(ctx context.Context, code string, joinAt uint64) (*model.Campaign, error) {
	var result *model.Campaign

	tx := r.db.WithContext(ctx).Model(&model.Campaign{})

	tx = tx.Where(sqlutil.EqualClause(model.Campaign_CODE), code)
	tx = tx.Where(sqlutil.LessThanOrEqualClause(model.Campaign_START_AT), joinAt)
	tx = tx.Where(sqlutil.GreaterThanOrEqualClause(model.Campaign_END_AT), joinAt)

	if err := tx.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}
