package service

import (
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/voucher"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/voucherpool"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(voucherpool.NewService),
	fx.Provide(voucher.NewService),
)
