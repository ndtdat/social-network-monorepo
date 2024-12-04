package rueidis

import (
	"fmt" //nolint:goimports,gofumpt
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"

	//nolint:nolintlint,gci
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidishook"
	//nolint:gofumpt
)

type MarshalFunc func(v any) ([]byte, error)

func NewClient(cfg *config.App) (rueidis.Client, error) { //nolint:lll
	redisCfg := cfg.Redis
	tracingCfg := cfg.Tracing
	tracingEnabled := tracingCfg.Enabled

	dsn := fmt.Sprintf("%s:%s", redisCfg.Host, redisCfg.Port)
	db := redisCfg.DB

	clientSideCachingCfg := redisCfg.ClientSideCaching
	opt := rueidis.ClientOption{
		Password:     redisCfg.Password,
		SelectDB:     redisCfg.DB,
		InitAddress:  []string{dsn},
		DisableRetry: redisCfg.DisableRetry,
	}

	switch clientSideCachingCfg.Enabled {
	case true:
		opt.CacheSizeEachConn = clientSideCachingCfg.CacheSizeMegaBytes * 1024 * 1024
		if clientSideCachingCfg.BroadcastMode {
			opt.ClientTrackingOptions = genBroadcastOptions(clientSideCachingCfg.Prefixes)
		}

	default:
		opt.DisableCache = true
	}

	client, err := rueidis.NewClient(opt)
	if err != nil {
		return nil, err
	}

	if tracingEnabled {
		client = rueidishook.WithHook(
			client,
			newDatadogHook(fmt.Sprintf("rueidis.%s", tracingCfg.ServiceName), dsn, db, true),
		)
	}

	return client, nil
}

func genBroadcastOptions(prefixes []string) []string {
	var opts []string
	for _, p := range prefixes {
		opts = append(opts, "PREFIX", p)
	}

	opts = append(opts, "BCAST")

	return opts
}
