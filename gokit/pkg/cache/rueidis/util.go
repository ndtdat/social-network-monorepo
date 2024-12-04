package rueidis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/log"
	"github.com/redis/rueidis"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
	"time"
)

type Setter[T any] func(context.Context) (*T, error)

func Set(ctx context.Context, r rueidis.Client, k, v string, duration time.Duration) error {
	return r.Do(ctx, r.B().Set().Key(k).Value(v).ExSeconds(int64(duration.Seconds())).Build()).Error()
}

func Incrby(ctx context.Context, r rueidis.Client, key string, increment int64) error {
	return r.Do(ctx, r.B().Incrby().Key(key).Increment(increment).Build()).Error()
}

func Decrby(ctx context.Context, r rueidis.Client, key string, decrement int64) error {
	return r.Do(ctx, r.B().Decrby().Key(key).Decrement(decrement).Build()).Error()
}

func Del(ctx context.Context, r rueidis.Client, k string) error {
	return r.Do(ctx, DelCmd(ctx, r, k)).Error()
}

func Get(
	ctx context.Context, r rueidis.Client, key string, clientSideCacheDuration time.Duration,
) (string, error) {
	var resp rueidis.RedisResult
	if clientSideCacheDuration > 0 {
		resp = r.DoCache(ctx, r.B().Get().Key(key).Cache(), clientSideCacheDuration)
	} else {
		resp = r.Do(ctx, r.B().Get().Key(key).Build())
	}

	if err := resp.Error(); err != nil {
		if rueidis.IsRedisNil(err) {
			return "", nil
		}

		return "", err
	}

	value, _ := resp.ToString()

	return value, nil
}

func SetProtoMessage[T proto.Message](
	ctx context.Context, r rueidis.Client, key string, value T, duration time.Duration,
) error {
	v, err := proto.Marshal(value)
	if err != nil {
		return err
	}

	return Set(ctx, r, key, rueidis.BinaryString(v), duration)
}

func GetProtoMessage[T proto.Message](
	ctx context.Context, r rueidis.Client, key string, result T, clientSideCacheDuration time.Duration,
) (*T, error) {
	value, err := Get(ctx, r, key, clientSideCacheDuration)
	if err != nil {
		return nil, err
	}
	if value == "" {
		return nil, nil
	}

	err = proto.Unmarshal([]byte(value), result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func SetStruct[T any](
	ctx context.Context, r rueidis.Client, key string, value T, duration time.Duration,
) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.Do(ctx, r.B().Set().Key(key).Value(rueidis.BinaryString(v)).
		ExSeconds(int64(duration.Seconds())).Build()).Error()
}

func GetStruct[T any](
	ctx context.Context, r rueidis.Client, key string, clientSideCacheDuration time.Duration,
) (*T, error) {
	value, err := Get(ctx, r, key, clientSideCacheDuration)
	if err != nil {
		return nil, err
	}
	if value == "" {
		return nil, err
	}

	var result T
	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func ComputeIfAbsent[T any](
	ctx context.Context, r rueidis.Client, key string, setter Setter[T], d time.Duration,
) (*T, error) {
	result, err := GetStruct[T](ctx, r, key, d)
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

func DelByPrefix(ctx context.Context, r rueidis.Client, prefix string) error {
	var (
		entry rueidis.ScanEntry
		err   error
	)

	for more := true; more; more = entry.Cursor != 0 {
		if entry, err = r.Do(
			ctx, r.B().
				Scan().
				Cursor(entry.Cursor).
				Match(fmt.Sprintf("%s*", prefix)).
				Build(),
		).AsScanEntry(); err != nil {
			return err
		}

		if err = DelMulti(ctx, r, entry.Elements); err != nil {
			return err
		}
	}

	return nil
}

func DelMulti(ctx context.Context, r rueidis.Client, ks []string) error {
	delCmds := []rueidis.Completed{r.B().Multi().Build()}

	for _, k := range ks {
		delCmds = append(delCmds, DelCmd(ctx, r, k))
	}
	delCmds = append(delCmds, r.B().Exec().Build())

	for _, rep := range r.DoMulti(ctx, delCmds...) {
		if err := rep.Error(); err != nil {
			return err
		}
	}

	return nil
}

func DelCmd(_ context.Context, r rueidis.Client, k string) rueidis.Completed {
	return r.B().Del().Key(k).Build()
}

func SAdd(ctx context.Context, r rueidis.Client, key string, value []string) error {
	return r.Do(ctx, r.B().Sadd().Key(key).Member(value...).Build()).Error()
}

func SMembers(ctx context.Context, r rueidis.Client, key string) ([]string, error) {
	members, err := r.Do(ctx, r.B().Smembers().Key(key).Build()).ToArray()
	if err != nil {
		return nil, err
	}

	var results []string
	for _, v := range members {
		// TODO check if err is Nil
		vStr, err := v.ToString()
		if err != nil {
			return nil, err
		}
		results = append(results, vStr)
	}

	return results, nil
}

func SRem(ctx context.Context, r rueidis.Client, key string, members []string) error {
	return r.Do(ctx, r.B().Srem().Key(key).Member(members...).Build()).Error()
}

func Expire(ctx context.Context, r rueidis.Client, key string, seconds int64) error {
	return r.Do(ctx, r.B().Expire().Key(key).Seconds(seconds).Build()).Error()
}

func ExpireAt(ctx context.Context, r rueidis.Client, key string, timestamp int64) error {
	return r.Do(ctx, r.B().Expireat().Key(key).Timestamp(timestamp).Build()).Error()
}

func ExpireTime(ctx context.Context, r rueidis.Client, key string) error {
	return r.Do(ctx, r.B().Expiretime().Key(key).Build()).Error()
}

type updater func() error

func DelByPrefixesIfSuccessful(_ context.Context, r rueidis.Client, prefixes []string, updater updater) error {
	err := updater()
	if err == nil && len(prefixes) > 0 {
		go func() {
			if delErr := DelByPrefixes(context.TODO(), r, prefixes); delErr != nil {
				log.Logger().Warn(fmt.Sprintf("Cannot del by prefix for %v", prefixes))
			}
		}()
	}

	return err
}

func DelByPrefixes(ctx context.Context, r rueidis.Client, prefixes []string) error {
	errGroup, _ := errgroup.WithContext(ctx)
	errGroup.SetLimit(-1)

	for _, prefix := range prefixes {
		safePrefix := prefix
		errGroup.Go(func() error {
			var (
				entry rueidis.ScanEntry
				err   error
			)

			for more := true; more; more = entry.Cursor != 0 {
				if entry, err = r.Do(
					ctx, r.B().
						Scan().
						Cursor(entry.Cursor).
						Match(fmt.Sprintf("%s*", safePrefix)).
						Build(),
				).AsScanEntry(); err != nil {
					return err
				}

				if err = DelMulti(ctx, r, entry.Elements); err != nil {
					return err
				}
			}

			return nil
		})
	}

	return errGroup.Wait()
}
