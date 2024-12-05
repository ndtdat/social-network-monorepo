package singlepod

import (
	"context"
	"fmt"
	"time"
)

type Starter func(ctx context.Context) error

func (s *Service) TryToStart(ctx context.Context, starter Starter) {
	logger := s.logger
	id := s.ID

	if err := s.startOnSinglePod(ctx, starter); err != nil {
		logger.Debug(fmt.Sprintf("cannot execute single pod %s cron due to: %v", id, err))
		time.Sleep(5 * time.Second)
		s.TryToStart(ctx, starter)
	}
}

func (s *Service) startOnSinglePod(ctx context.Context, starter Starter) error {
	id := s.ID
	// Check another pod is running
	running, err := s.CheckPodRunningAndMarkRunningIfPossible(ctx)
	if err != nil {
		return err
	}
	if running {
		return fmt.Errorf("%s cron is running on other pod", id)
	}

	if err = starter(ctx); err != nil {
		return fmt.Errorf("cannot start %s cron due to %v", id, err)
	}

	//Keep live the redis key to prevent other pod run
	go func() {
		s.KeepAlivePermission(ctx)
	}()

	return nil
}
