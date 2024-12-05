package repository

import (
	"github.com/ndtdat/social-network-monorepo/user-service/internal/repository/campaign"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/repository/user"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/repository/usercampaign"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(user.NewRepository),
	fx.Provide(campaign.NewRepository),
	fx.Provide(usercampaign.NewRepository),
)
