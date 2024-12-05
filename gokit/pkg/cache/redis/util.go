package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
	"time"
)

type Setter[T any] func(context.Context) (*T, error)

func SetStruct[T any](
	ctx context.Context, r redis.UniversalClient, key string, value T, duration time.Duration,
) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.Set(ctx, key, v, duration).Err()
}

func GetStruct[T any](ctx context.Context, r redis.UniversalClient, key string) (*T, error) {
	value, err := r.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, err
	}

	var result T
	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func SetProtoMessage[T proto.Message](
	ctx context.Context, r redis.UniversalClient, key string, value T, duration time.Duration,
) error {
	v, err := proto.Marshal(value)
	if err != nil {
		return err
	}

	return r.Set(ctx, key, v, duration).Err()
}

func GetProtoMessage[T proto.Message](ctx context.Context, r redis.UniversalClient, key string, result T) (*T, error) {
	value, err := r.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, err
	}

	err = proto.Unmarshal([]byte(value), result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func ComputeIfAbsent[T any](
	ctx context.Context, r redis.UniversalClient, key string, setter Setter[T], d time.Duration,
) (*T, error) {
	result, err := GetStruct[T](ctx, r, key)
	if err != nil {
		return nil, err
	}
	if result != nil {
		return result, nil
	}

	value, err := setter(ctx)
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, nil
	}

	_ = SetStruct[T](ctx, r, key, *value, d)

	return value, nil
}

func SAdd(
	ctx context.Context, r redis.UniversalClient, key string, values ...string,
) error {
	if len(values) == 0 {
		return fmt.Errorf("values is empty")
	}

	return r.SAdd(ctx, key, values).Err()
}

func SInterStruct(ctx context.Context, r redis.UniversalClient, key string) ([]string, error) {
	results, err := r.SInter(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, err
	}

	return results, nil
}

func SIsMember(ctx context.Context, r redis.UniversalClient, key string, value string) (bool, error) {
	presence, err := r.SIsMember(ctx, key, value).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}

		return false, err
	}

	return presence, nil
}
