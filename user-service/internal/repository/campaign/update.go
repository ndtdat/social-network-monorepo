package campaign

import (
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	pbmodel "github.com/ndtdat/social-network-monorepo/user-service/pkg/api/go/user/model"
	"gorm.io/gorm"
)

func (r *Repository) UpdateWithTx(tx *gorm.DB, campaignID uint64) error {
	return tx.Model(&model.Campaign{}).
		Where(sqlutil.EqualClause(model.Campaign_ID), campaignID).
		Updates(map[string]any{
			model.Campaign_JOINED_QTY: gorm.Expr(
				fmt.Sprintf("%s + 1", model.Campaign_JOINED_QTY),
			),
		}).Update(model.Campaign_STATUS,
		gorm.Expr(
			fmt.Sprintf(`
			IF (%[1]s >= %[2]s , ?, %[3]s)
		`,
				model.Campaign_JOINED_QTY,
				model.Campaign_MAX_QTY,
				model.Campaign_STATUS,
			),
			pbmodel.CampaignStatus_CS_UNAVAILABLE,
		)).Error
}
