package service

import (
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/auth"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/campaign"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(auth.NewService),
	fx.Provide(campaign.NewService),
)
