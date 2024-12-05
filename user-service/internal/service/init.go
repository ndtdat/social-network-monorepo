package service

import (
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/agent/purchase"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/auth"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/campaign"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/cron/manager"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/usercampaign"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(auth.NewService),
	fx.Provide(campaign.NewService),
	fx.Provide(usercampaign.NewService),
	fx.Provide(manager.NewManager),
	fx.Provide(purchase.NewService),
)
