package purchase

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/rpc"
)

func (s *Service) IAllocateVoucherByCampaignID(ctx context.Context, userID, campaignID uint64) error {
	client, err := s.microservices.PurchaseClient()
	if err != nil {
		return err
	}

	if _, err = client.IAllocateVoucherByCampaignID(ctx, &rpc.IAllocateVoucherByCampaignIDRequest{
		UserId:     userID,
		CampaignId: campaignID,
	}); err != nil {
		return err
	}

	return nil
}
