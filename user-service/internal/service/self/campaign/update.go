package campaign

import (
	"gorm.io/gorm"
)

func (s *Service) IncreaseJoinedQty(tx *gorm.DB, campaignID uint64) error {
	return s.repo.UpdateWithTx(tx, campaignID)
}
