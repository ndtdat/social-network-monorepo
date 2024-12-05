package service

import (
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/cron/manager"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/subscriptionplan"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/voucher"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/voucherpool"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(voucherpool.NewService),
	fx.Provide(voucher.NewService),
	fx.Provide(subscriptionplan.NewService),
	fx.Provide(manager.NewManager),
)
