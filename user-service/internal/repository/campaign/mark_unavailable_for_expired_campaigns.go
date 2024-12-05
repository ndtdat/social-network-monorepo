package campaign

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/sqlutil"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	model2 "github.com/ndtdat/social-network-monorepo/user-service/pkg/api/go/user/model"
)

func (r *Repository) MarkUnavailableForExpiredCampaigns(ctx context.Context) error {
	return r.db.WithContext(ctx).Model(&model.Campaign{}).
		Where(sqlutil.EqualClause(model.Campaign_STATUS), model2.CampaignStatus_CS_AVAILABLE).
		Where(sqlutil.LessThanClause(model.Campaign_END_AT), util.CurrentUnix()).
		Updates(map[string]any{
			model.Campaign_STATUS: model2.CampaignStatus_CS_UNAVAILABLE,
		}).Error
}
