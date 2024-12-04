package concurrentset

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/concurrentmap"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
)

type IterCb[K comparable] func(k K) bool

type ConcurrentSet[K comparable] struct {
	cmap  *cmap.ConcurrentMap[K, *util.NoUse]
	noUse *util.NoUse
}

func New[K comparable](nShard int) *ConcurrentSet[K] {
	return &ConcurrentSet[K]{
		cmap:  cmap.New[K, *util.NoUse](nShard),
		noUse: &util.NoUse{},
	}
}

func (s *ConcurrentSet[K]) Add(key K) {
	s.cmap.Set(key, s.noUse)
}

func (s *ConcurrentSet[K]) BufferedItems() []K {
	var keys []K
	for it := range s.cmap.IterBuffered() {
		keys = append(keys, it.Key)
	}

	return keys
}

func (s *ConcurrentSet[K]) IterCb(fn IterCb[K]) {
	s.cmap.IterCb(func(k K, _ *util.NoUse) bool {
		return fn(k)
	})
}

func (s *ConcurrentSet[K]) Remove(key K) {
	s.cmap.Remove(key)
}

func (s *ConcurrentSet[K]) Clear() {
	s.cmap.Clear()
}

func (s *ConcurrentSet[K]) Count() int {
	return s.cmap.Count()
}

func (s *ConcurrentSet[K]) Contains(key K) bool {
	return s.cmap.Contains(key)
}

func (s *ConcurrentSet[K]) IsEmpty() bool {
	return s.cmap.IsEmpty()
}
