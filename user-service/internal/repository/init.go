package repository

import (
	"github.com/ndtdat/social-network-monorepo/user-service/internal/repository/campaign"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/repository/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(user.NewRepository),
	fx.Provide(campaign.NewRepository),
)
