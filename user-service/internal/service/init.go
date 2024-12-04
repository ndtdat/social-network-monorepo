package service

import (
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/auth"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(auth.NewService),
)
