package campaign

import (
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	pbmodel "github.com/ndtdat/social-network-monorepo/user-service/pkg/api/go/user/model"
	"gorm.io/gorm"
)

func (s *Service) SelectValidCampaignForUpdate(tx *gorm.DB, code string, joinAt uint64) (*model.Campaign, error) {
	return s.repo.SelectForUpdate(tx, code, joinAt, pbmodel.CampaignStatus_CS_AVAILABLE)
}
