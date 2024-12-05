package repository

import (
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/uservoucher"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/repository/voucherpool"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(voucherpool.NewRepository),
	fx.Provide(uservoucher.NewRepository),
)
