package jwt

import (
	"context"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/enum"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/base"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/jwt/eddsa"
	"go.uber.org/zap"
)

type InitFunc func(
	ctx context.Context, cfg *config.App, logger *zap.Logger,
) (base.Manager, error)

var initMap = map[enum.Algorithm]InitFunc{
	enum.Algorithm_EDDSA: eddsa.NewManager,
}

func NewManager(
	ctx context.Context, cfg *config.App, logger *zap.Logger,
) (base.Manager, error) {
	algo := cfg.JWT.Algorithm
	initMethod, exist := initMap[algo]
	if exist {
		return initMethod(ctx, cfg, logger)
	}

	return nil, fmt.Errorf("found unsupported algorithm for JWT: %s", algo)
}
