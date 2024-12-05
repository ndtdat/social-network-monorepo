package redis

import (
	"fmt" //nolint:goimports,gofumpt
	"github.com/go-redis/redis/v8"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	//nolint:nolintlint,gci
	//nolint:gofumpt
)

func NewClient(cfg *config.App) (redis.UniversalClient, error) { //nolint:lll
	redisCfg := cfg.Redis

	dsn := fmt.Sprintf("%s:%s", redisCfg.Host, redisCfg.Port)

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{dsn},
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	return client, nil
}
