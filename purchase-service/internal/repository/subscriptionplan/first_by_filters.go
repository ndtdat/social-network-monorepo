package subscriptionplan

import (
	"context"
	"errors"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	pbmodel "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
	"gorm.io/gorm"
)

func (r *Repository) FirstByFilters(
	ctx context.Context, subscriptionPlanTier pbmodel.SubscriptionPlanTier,
) (*model.SubscriptionPlan, error) {
	var result *model.SubscriptionPlan

	tx := r.db.WithContext(ctx).Model(&model.SubscriptionPlan{})

	tx = tx.Where(sqlutil.EqualClause(model.SubscriptionPlan_TIER), subscriptionPlanTier)

	if err := tx.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}
