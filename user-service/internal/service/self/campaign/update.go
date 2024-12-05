package campaign

import (
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	"gorm.io/gorm"
)

func (s *Service) UpdateWithTx(tx *gorm.DB, updatedCampaign *model.Campaign) error {
	return s.repo.UpdateWithTx(tx, updatedCampaign)
}
