package server

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/rpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *PurchaseServer) IAllocateVoucherByCampaignID(
	ctx context.Context, in *rpc.IAllocateVoucherByCampaignIDRequest,
) (*emptypb.Empty, error) {
	if err := p.voucherService.AllocateByCampaignID(ctx, in.GetUserId(), in.GetCampaignId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
