package campaign

import (
	"errors"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	pbmodel "github.com/ndtdat/social-network-monorepo/user-service/pkg/api/go/user/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *Repository) SelectForUpdate(
	tx *gorm.DB, code string, joinAt uint64, status pbmodel.CampaignStatus,
) (*model.Campaign, error) {
	var result *model.Campaign

	if err := tx.Model(&model.Campaign{}).
		Where(sqlutil.EqualClause(model.Campaign_CODE), code).
		Where(sqlutil.EqualClause(model.Campaign_STATUS), status).
		Where(sqlutil.LessThanOrEqualClause(model.Campaign_START_AT), joinAt).
		Where(sqlutil.GreaterThanOrEqualClause(model.Campaign_END_AT), joinAt).
		Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
		First(&result).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}
