package local

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/concurrentmap"
	"go.uber.org/zap"
)

type Item any

type Service struct {
	logger    *zap.Logger
	cache     *cmap.ConcurrentMap[string, Item]
	maxNumKey int
}

func NewService(logger *zap.Logger, maxNumKey uint32, nShard int) *Service {
	return &Service{
		logger:    logger,
		cache:     cmap.New[string, Item](nShard),
		maxNumKey: int(maxNumKey),
	}
}

func (s *Service) Set(key string, item Item) {
	if s.cache.Count() >= s.maxNumKey {
		s.cache.Clear()
	}

	s.cache.Set(key, item)
}

func (s *Service) Get(key string) (Item, bool) {
	item, exist := s.cache.Get(key)

	return item, exist
}

func (s *Service) Remove(key string) {
	s.cache.Remove(key)
}
