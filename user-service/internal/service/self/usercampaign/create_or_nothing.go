package usercampaign

import (
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	"gorm.io/gorm"
)

func (s *Service) CreateOrNothingWithTx(tx *gorm.DB, userCampaign *model.UserCampaign) error {
	return s.repo.CreateOrNothingWithTx(tx, userCampaign)
}
