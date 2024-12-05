package subscriptionplan

import (
	"context"
	"fmt"
	model2 "github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
	"github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/rpc"
	"gorm.io/gorm"
)

func (s *Service) Buy(
	ctx context.Context, userID uint64, subscriptionPlanTier model.SubscriptionPlanTier,
) (*rpc.BuySubscriptionPlanReply, error) {
	// Get subscription plan
	subscriptionPlan, err := s.subscriptionPlanRepo.FirstByFilters(ctx, subscriptionPlanTier)
	if err != nil {
		return nil, err
	}
	if subscriptionPlan == nil {
		return nil, err
	}

	// Get current user tier
	userTierInfo, err := s.userTierRepo.FirstByFilters(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userTierInfo != nil {
		userTier := userTierInfo.Tier
		if userTier >= subscriptionPlanTier {
			return nil, fmt.Errorf(
				"cannot buy plan tier [%v] because buy tier [%v] < current tier [%v]",
				subscriptionPlanTier, userTier, subscriptionPlanTier,
			)
		}

		userTierInfo.Tier = subscriptionPlanTier
	} else {
		userTierInfo = &model2.UserTier{
			UserID: userID,
			Tier:   subscriptionPlanTier,
		}
	}

	// Get voucher by userID and subscription plan
	userVoucher, err := s.userVoucherRepo.FirstByFilters(ctx, userID, subscriptionPlanTier)
	if err != nil {
		return nil, err
	}

	txInfo := s.genTx(userID, subscriptionPlan, userVoucher)

	if err = s.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx)

		// Mark voucher is used
		if userVoucher != nil && userVoucher.Status == model.UserVoucherStatus_UVS_ALLOCATED {
			if err = s.voucherService.MarkUserVoucherIsUsedWithTx(
				tx, userVoucher.ID, userVoucher.VoucherConfigurationID,
			); err != nil {
				return err
			}
		}

		// Upsert user tier
		if err = s.userTierRepo.UpsertWithTx(tx, userTierInfo); err != nil {
			return err
		}

		// Create transaction
		if err = s.transactionRepo.CreateOrNothingWithTx(tx, txInfo); err != nil {
			return err
		}

		// TODO: Call payment to debit user's balance

		return nil
	}); err != nil {
		return nil, err
	}

	return &rpc.BuySubscriptionPlanReply{}, nil
}
