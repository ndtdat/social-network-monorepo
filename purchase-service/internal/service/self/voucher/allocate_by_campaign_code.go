package voucher

import (
	"context"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/common/pkg/sorter"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/suid"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	model2 "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
	"gorm.io/gorm"
	"time"
)

const ExpireDay = 30

func (s *Service) AllocateByCampaignID(ctx context.Context, userID, campaignID uint64) error {
	var (
		configurationRepo = s.configurationRepo
		now               = time.Now()
	)
	// Check voucher is existed by campaign id
	voucherConfiguration, err := configurationRepo.GetAvailableByCampaignID(ctx, campaignID)
	if err != nil {
		return err
	}
	if voucherConfiguration == nil {
		return fmt.Errorf("voucher for campaign code [%v] not found", campaignID)
	}

	voucherConfigurationID := voucherConfiguration.ID

	userVoucher := &model.UserVoucher{
		ID:                     suid.New(),
		UserID:                 userID,
		VoucherConfigurationID: voucherConfigurationID,
		ExpiredAt:              uint64(now.AddDate(0, 0, ExpireDay).Unix()),
		Status:                 model2.UserVoucherStatus_UVS_ALLOCATED,
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx)

		// Get code from pool
		code, err := s.voucherPoolRepo.SelectFirstForUpdate(tx, []*sorter.Sorter{
			{
				Field: model.VoucherPool_CREATED_AT,
				Order: sorter.Order_ASC,
			},
		})
		if err != nil {
			return err
		}

		// Update voucher configuration
		if err = configurationRepo.IncreaseAllocatedQty(tx, voucherConfigurationID); err != nil {
			return err
		}

		// Create user voucher
		userVoucher.VoucherCode = code
		if err = s.userVoucherRepo.CreateWithTx(tx, userVoucher); err != nil {
			return err
		}

		return nil
	})
}
