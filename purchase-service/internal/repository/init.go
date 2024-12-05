package repository

import (
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/subscriptionplan"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/transaction"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/usertier"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/uservoucher"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/voucherconfiguration"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/voucherpool"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(voucherpool.NewRepository),
	fx.Provide(uservoucher.NewRepository),
	fx.Provide(voucherconfiguration.NewRepository),
	fx.Provide(usertier.NewRepository),
	fx.Provide(subscriptionplan.NewRepository),
	fx.Provide(transaction.NewRepository),
)
