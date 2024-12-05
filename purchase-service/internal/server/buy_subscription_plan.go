package server

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/common/pkg/auth"
	"github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/rpc"
)

func (p *PurchaseServer) BuySubscriptionPlan(
	ctx context.Context, in *rpc.BuySubscriptionPlanRequest,
) (*rpc.BuySubscriptionPlanReply, error) {
	return p.subscriptionPlanService.Buy(ctx, auth.IdentityIDFromCtx(ctx), in.GetSubscriptionPlanTier())
}
