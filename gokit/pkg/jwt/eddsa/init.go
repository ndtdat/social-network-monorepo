package eddsa

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewManager),
)
