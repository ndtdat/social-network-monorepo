package user

import (
	"context"
	"errors"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	"gorm.io/gorm"
)

func (r *Repository) FirstByFilters(ctx context.Context, email string) (*model.User, error) {
	var result *model.User

	tx := r.db.WithContext(ctx).Model(&model.User{})

	tx = tx.Where(sqlutil.EqualClause(model.User_EMAIL), email)

	if err := tx.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}
