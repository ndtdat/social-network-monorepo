package singlepod

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/redis/rueidis"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Service struct {
	redisClient             rueidis.Client
	logger                  *zap.Logger
	redSync                 *redsync.Redsync
	Prefix                  string
	ID                      string
	Duration                time.Duration
	clientSideCacheDuration time.Duration
	keepAliveDuration       time.Duration
	sync.Mutex
	isRunning bool
}

func NewService(
	logger *zap.Logger, id string, redisClient rueidis.Client, redSync *redsync.Redsync, prefix string,
	duration time.Duration, clientSideCacheDuration time.Duration, keepAliveDuration time.Duration,
) *Service {
	return &Service{
		logger:                  logger,
		redisClient:             redisClient,
		redSync:                 redSync,
		Prefix:                  prefix,
		Duration:                duration,
		clientSideCacheDuration: clientSideCacheDuration,
		keepAliveDuration:       keepAliveDuration,
		isRunning:               false,
		ID:                      id,
	}
}
