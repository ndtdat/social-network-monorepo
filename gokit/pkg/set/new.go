package set

import (
	"github.com/dolthub/swiss"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
)

type Set[K comparable] struct {
	imap  *swiss.Map[K, *util.NoUse]
	noUse *util.NoUse
}

func New[K comparable]() *Set[K] {
	return &Set[K]{
		imap:  swiss.NewMap[K, *util.NoUse](42),
		noUse: &util.NoUse{},
	}
}

func (s *Set[K]) Add(key K) {
	s.imap.Put(key, s.noUse)
}

func (s *Set[K]) ItemMap() *swiss.Map[K, *util.NoUse] {
	return s.imap
}

func (s *Set[K]) ItemArray() []K {
	var keys []K
	s.imap.Iter(func(k K, _ *util.NoUse) bool {
		keys = append(keys, k)

		return false
	})

	return keys
}

func (s *Set[K]) Remove(key K) {
	s.imap.Delete(key)
}

func (s *Set[K]) Clear() {
	s.imap = swiss.NewMap[K, *util.NoUse](42)
}

func (s *Set[K]) Count() int {
	return s.imap.Count()
}

func (s *Set[K]) Contains(key K) bool {
	_, exist := s.imap.Get(key)

	return exist
}

func (s *Set[K]) IsEmpty() bool {
	return s.Count() == 0
}

func (s *Set[K]) Get(key K) *util.NoUse {
	value, ok := s.imap.Get(key)
	if !ok {
		return nil
	}

	return value
}
