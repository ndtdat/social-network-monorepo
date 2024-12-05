package singlepod

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/cache/rueidis"
	"time"
)

func (s *Service) getMutexKey() string {
	return fmt.Sprintf("%s_MUTEX", s.Prefix)
}

func (s *Service) lockRedisKey() (*redsync.Mutex, error) {
	key := s.getMutexKey()
	singlePodMutex := s.redSync.NewMutex(key)
	if err := singlePodMutex.Lock(); err != nil {
		return nil, fmt.Errorf("cannot get lock for %s due to: %v", key, err)
	}

	return singlePodMutex, nil
}

func (s *Service) unlockRedisKey(singlePodMutex *redsync.Mutex) {
	if _, err := singlePodMutex.Unlock(); err != nil {
		s.logger.Error(fmt.Sprintf("cannot unlock for %s due to: %v", s.getMutexKey(), err))
	}
}

func (s *Service) markAsRunning(ctx context.Context) error {
	singlePodMutex, err := s.lockRedisKey()
	if err != nil {
		return err
	}

	defer func() {
		s.unlockRedisKey(singlePodMutex)
	}()

	prefix := s.Prefix
	if err := rueidis.Set(
		ctx, s.redisClient, prefix, prefix, s.Duration,
	); err != nil {
		return fmt.Errorf("cannot mark running checker to redis due to: %v", err)
	}

	return nil
}

func (s *Service) deleteRunning(ctx context.Context) error {
	singlePodMutex, err := s.lockRedisKey()
	if err != nil {
		return err
	}

	defer func() {
		s.unlockRedisKey(singlePodMutex)
	}()

	prefix := s.Prefix

	if err := rueidis.Del(ctx, s.redisClient, prefix); err != nil {
		return fmt.Errorf("cannot delete key(%v) due to: %v", prefix, err)
	}

	return nil
}

func (s *Service) CheckPodRunningAndMarkRunningIfPossible(ctx context.Context) (bool, error) {
	singlePodMutex, err := s.lockRedisKey()
	if err != nil {
		return false, err
	}
	defer func() {
		s.unlockRedisKey(singlePodMutex)
	}()

	value, err := rueidis.Get(ctx, s.redisClient, s.Prefix, s.clientSideCacheDuration)
	if err != nil {
		return false, fmt.Errorf("cannot get running checker from cache due to: %v", err)
	}

	isRunning := value != ""

	if !isRunning {
		// Mark is running
		prefix := s.Prefix
		if err = rueidis.Set(
			ctx, s.redisClient, prefix, prefix, s.Duration,
		); err != nil {
			return false, fmt.Errorf("cannot mark running checker to redis due to: %v", err)
		}

		s.Lock()
		s.isRunning = true
		s.Unlock()
	}

	return isRunning, nil
}

func (s *Service) MarkAsNotRunningIfPossible(ctx context.Context) error {
	if !s.isRunning {
		return nil
	}

	if err := s.deleteRunning(ctx); err != nil {
		return err
	}

	s.Lock()
	s.logger.Debug(fmt.Sprintf("Mark %s as not running", s.ID))
	s.isRunning = false
	s.Unlock()

	return nil
}

func (s *Service) KeepAlivePermission(ctx context.Context) {
	var (
		failToMarkRunning = false
		logger            = s.logger
	)

	for ctx.Err() == nil {
		if !failToMarkRunning {
			time.Sleep(s.keepAliveDuration)
		}

		s.Lock()
		if !s.isRunning {
			return
		}
		s.Unlock()

		if err := s.markAsRunning(ctx); err != nil {
			logger.Debug(fmt.Sprintf("cannot increase expiration of cron running key due to %v", err))
			time.Sleep(100 * time.Microsecond)
			failToMarkRunning = true

			continue
		}

		if failToMarkRunning {
			failToMarkRunning = false
		}
	}
}
