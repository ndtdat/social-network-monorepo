package cmap

import (
	"encoding/json"
	"fmt"
	"github.com/dolthub/swiss"
	"sync"
)

type ConcurrentMap[K comparable, V any] struct {
	shards []*Shard[K, V]
	nShard int
}

type Shard[K comparable, V any] struct {
	items *swiss.Map[K, V]
	sync.RWMutex
}

func New[K comparable, V any](nShard int) *ConcurrentMap[K, V] {
	m := ConcurrentMap[K, V]{
		shards: []*Shard[K, V]{},
		nShard: nShard,
	}
	for i := 0; i < nShard; i++ {
		m.shards = append(m.shards, &Shard[K, V]{items: swiss.NewMap[K, V](42)})
	}

	return &m
}

func (m *ConcurrentMap[K, V]) GetShard(key K) *Shard[K, V] {
	return m.shards[uint(fnv32(key))%uint(m.nShard)]
}

func (m *ConcurrentMap[K, V]) MSet(data map[K]V) {
	for key, value := range data {
		shard := m.GetShard(key)
		shard.Lock()
		shard.items.Put(key, value)
		shard.Unlock()
	}
}

func (m *ConcurrentMap[K, V]) Set(key K, value V) {
	shard := m.GetShard(key)
	shard.Lock()
	shard.items.Put(key, value)
	shard.Unlock()
}

// UpsertCb Callback to return new element to be inserted into the map
// It is called while lock is held, therefore it MUST NOT
// try to access other keys in same map, as it can lead to deadlock since
// Go sync.RWLock is not reentrant
type UpsertCb[K comparable, V any] func(exist bool, valueInMap V, newValue V) V

// Upsert Insert or Update - updates existing element or inserts a new one using UpsertCb
//
//nolint:nonamedreturns
func (m *ConcurrentMap[K, V]) Upsert(key K, value V, cb UpsertCb[K, V]) (res V) {
	shard := m.GetShard(key)
	shard.Lock()
	v, ok := shard.items.Get(key)
	res = cb(ok, v, value)
	shard.items.Put(key, res)
	shard.Unlock()

	return res
}

func (m *ConcurrentMap[K, V]) SetIfAbsent(key K, value V) bool {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	_, ok := shard.items.Get(key)
	if !ok {
		shard.items.Put(key, value)
	}
	shard.Unlock()

	return !ok
}

func (m *ConcurrentMap[K, V]) ComputeIfPresent(key K, mapFunc func(key K, value V) (V, bool)) (V, bool) {
	shard := m.GetShard(key)
	shard.Lock()
	value, ok := shard.items.Get(key)
	if ok {
		newValue, remove := mapFunc(key, value)
		if remove {
			shard.items.Delete(key)
			shard.Unlock()

			return value, true
		}

		shard.items.Put(key, newValue)
		shard.Unlock()

		return newValue, false
	}
	shard.Unlock()

	return value, false
}

func (m *ConcurrentMap[K, V]) ComputeIfAbsent(key K, mapFunc func(key K, value V) V) V {
	shard := m.GetShard(key)
	shard.Lock()
	value, ok := shard.items.Get(key)
	returnValue := value
	if !ok {
		newValue := mapFunc(key, value)
		shard.items.Put(key, newValue)
		returnValue = newValue
	}

	shard.Unlock()

	return returnValue
}

func (m *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	shard := m.GetShard(key)
	shard.RLock()
	val, ok := shard.items.Get(key)
	shard.RUnlock()

	return val, ok
}

func (m *ConcurrentMap[K, V]) Count() int {
	count := 0
	for i := 0; i < m.nShard; i++ {
		shard := m.shards[i]
		shard.RLock()
		count += shard.items.Count()
		shard.RUnlock()
	}

	return count
}

func (m ConcurrentMap[K, _]) Contains(key K) bool {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	_, ok := shard.items.Get(key)
	shard.RUnlock()

	return ok
}

func (m ConcurrentMap[K, _]) Remove(key K) {
	shard := m.GetShard(key)
	shard.Lock()
	shard.items.Delete(key)
	shard.Unlock()
}

// RemoveCb is a callback executed in a map.RemoveCb() call, while Lock is held
// If returns true, the element will be removed from the map
type RemoveCb[K comparable, V any] func(key K, v V, exists bool) bool

// RemoveCb locks the shard containing the key, retrieves its current value and calls the callback with those params
// If callback returns true and element exists, it will remove it from the map
// Returns the value returned by the callback (even if element was not present in the map)
func (m *ConcurrentMap[K, V]) RemoveCb(key K, cb RemoveCb[K, V]) bool {
	shard := m.GetShard(key)
	shard.Lock()
	v, ok := shard.items.Get(key)
	remove := cb(key, v, ok)
	if remove && ok {
		shard.items.Delete(key)
	}
	shard.Unlock()

	return remove
}

//nolint:nonamedreturns
func (m *ConcurrentMap[K, V]) Pop(key K) (v V, exists bool) {
	shard := m.GetShard(key)
	shard.Lock()
	v, exists = shard.items.Get(key)
	shard.items.Delete(key)
	shard.Unlock()

	return v, exists
}

// IsEmpty checks if map is empty.
func (m ConcurrentMap[_, _]) IsEmpty() bool {
	return m.Count() == 0
}

// Tuple Used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type Tuple[K comparable, V any] struct {
	Key K
	Val V
}

// Iter returns an iterator which could be used in a for range loop.
//
// Deprecated: using IterBuffered() will get a better performance
func (m *ConcurrentMap[K, V]) Iter() <-chan Tuple[K, V] {
	chans := snapshot(m)
	ch := make(chan Tuple[K, V])
	go fanIn(chans, ch)

	return ch
}

// IterBuffered returns a buffered iterator which could be used in a for range loop.
func (m *ConcurrentMap[K, V]) IterBuffered() <-chan Tuple[K, V] {
	chans := snapshot(m)
	total := 0
	for _, c := range chans {
		total += cap(c)
	}
	ch := make(chan Tuple[K, V], total)
	go fanIn(chans, ch)

	return ch
}

// Clear removes all items from map.
func (m ConcurrentMap[_, _]) Clear() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

// Returns an array of channels that contains elements in each shard,
// which likely takes a snapshot of `m`.
// It returns once the size of each buffered channel is determined,
// before all the channels are populated using goroutines.
//
//nolint:nonamedreturns
func snapshot[K comparable, V any](m *ConcurrentMap[K, V]) (chans []chan Tuple[K, V]) {
	//When you access map items before initializing.
	if len(m.shards) == 0 {
		panic(`cmap.ConcurrentMap is not initialized. Should run New() before usage.`)
	}
	nShard := m.nShard
	chans = make([]chan Tuple[K, V], nShard)
	wg := sync.WaitGroup{}
	wg.Add(nShard)
	// Foreach shard.
	for index, shard := range m.shards {
		go func(index int, shard *Shard[K, V]) {
			// Foreach key, value pair.
			shard.RLock()
			chans[index] = make(chan Tuple[K, V], shard.items.Count())
			wg.Done()
			shard.items.Iter(func(k K, v V) (stop bool) {
				chans[index] <- Tuple[K, V]{k, v}

				return false
			})

			shard.RUnlock()
			close(chans[index])
		}(index, shard)
	}
	wg.Wait()

	return chans
}

// fanIn reads elements from channels `chans` into channel `out`
func fanIn[K comparable, V any](chans []chan Tuple[K, V], out chan Tuple[K, V]) {
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan Tuple[K, V]) {
			for t := range ch {
				out <- t
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(out)
}

// Items returns all items as map[string]any
func (m *ConcurrentMap[K, V]) Items() map[K]V {
	tmp := make(map[K]V)

	// Insert items to temporary map.
	for item := range m.IterBuffered() {
		tmp[item.Key] = item.Val
	}

	return tmp
}

// IterCb Iterator callback,called for every key,value found in
// maps. RLock is held for all calls for a given shard
// therefore callback sess consistent view of a shard,
// but not across the shards
type IterCb[K comparable, V any] func(key K, v V) bool

// IterCb Callback based iterator, cheapest way to read
// all elements in a map.
func (m *ConcurrentMap[K, V]) IterCb(fn IterCb[K, V]) {
	for _, shard := range m.shards {
		shard.RLock()

		shard.items.Iter(func(k K, v V) bool {
			return fn(k, v)
		})

		shard.RUnlock()
	}
}

func (m *ConcurrentMap[K, V]) Keys() []K {
	count := m.Count()
	ch := make(chan K, count)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(m.nShard)
		for _, shard := range m.shards {
			go func(shard *Shard[K, V]) {
				shard.RLock()

				shard.items.Iter(func(k K, _ V) bool {
					ch <- k

					return false
				})

				shard.RUnlock()
				wg.Done()
			}(shard)
		}
		wg.Wait()
		close(ch)
	}()

	keys := make([]K, 0, count)
	for k := range ch {
		keys = append(keys, k)
	}

	return keys
}

// MarshalJSON Reviles ConcurrentMap "private" variables to json marshal.
func (m *ConcurrentMap[K, V]) MarshalJSON() ([]byte, error) {
	tmp := make(map[K]V)

	for item := range m.IterBuffered() {
		tmp[item.Key] = item.Val
	}

	return json.Marshal(tmp)
}

const (
	hash    = uint32(2166136261)
	prime32 = uint32(16777619)
)

func fnv32[K comparable](key K) uint32 {
	b := []byte(fmt.Sprint(key))
	s := string(b)
	keyLength := len(s)
	h := hash

	for i := 0; i < keyLength; i++ {
		h *= prime32
		h ^= uint32(s[i])
	}

	return hash
}
