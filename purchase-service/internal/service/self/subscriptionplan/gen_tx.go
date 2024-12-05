package subscriptionplan

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/suid"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	model2 "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
	"github.com/shopspring/decimal"
)

func (s *Service) genTx(
	userID uint64, subscriptionPlan *model.SubscriptionPlan, userVoucher *model.UserVoucher,
) *model.Transaction {
	var (
		originalAmount = subscriptionPlan.Price
		discountAmount = decimal.Zero
		userVoucherID  uint64
	)

	if userVoucher != nil {
		userVoucherID = userVoucher.ID
		voucherConfiguration := userVoucher.VoucherConfiguration
		switch userVoucher.VoucherConfiguration.CalculationType {
		case model2.VoucherCalculationType_VCT_AMOUNT:
			discountAmount = voucherConfiguration.Value
		case model2.VoucherCalculationType_VCT_PERCENTAGE:
			discountAmount = originalAmount.Mul(voucherConfiguration.Value)
		}
	}

	return &model.Transaction{
		ID:                   suid.New(),
		UserID:               userID,
		CurrencySymbol:       subscriptionPlan.CurrencySymbol,
		OriginalAmount:       originalAmount,
		DiscountAmount:       discountAmount,
		ActualAmount:         originalAmount.Sub(discountAmount),
		SubscriptionPlanTier: subscriptionPlan.Tier,
		UserVoucherID:        userVoucherID,
	}
}
