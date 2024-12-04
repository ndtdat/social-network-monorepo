package redissync

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

func NewService(client redis.UniversalClient) (*redsync.Redsync, error) {
	return redsync.New(goredis.NewPool(client)), nil
}
