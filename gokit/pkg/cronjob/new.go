package cronjob

import (
	"context"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"time"
)

const defaultTaskTimeout = time.Minute * 2

type Cronjob struct {
	logger       *zap.Logger
	cfg          *config.Cronjob
	handler      *Handler
	cron         *cron.Cron
	ID           string
	lastProgress int64
	taskTimeout  time.Duration
}

func NewCronjob(
	logger *zap.Logger, cfg *config.Cronjob, handler *Handler,
) *Cronjob {
	taskTimout := cfg.TaskTimeout
	if taskTimout == 0 {
		taskTimout = defaultTaskTimeout
	}

	return &Cronjob{
		logger:       logger,
		ID:           cfg.ID,
		cfg:          cfg,
		handler:      handler,
		cron:         cron.New(cron.WithChain(cron.SkipIfStillRunning(NewLogger(logger)))),
		lastProgress: util.CurrentUnix(),
		taskTimeout:  taskTimout,
	}
}

func (c *Cronjob) Stop() error {
	if c != nil {
		c.logger.Info(fmt.Sprintf("Shutdown cronjob %s", c.ID))

		return c.retire()
	}

	return nil
}

func (c *Cronjob) addTask() (cron.EntryID, error) {
	return c.cron.AddFunc(c.cfg.Spec, c.execute)
}

func (c *Cronjob) Start() error {
	if c.cfg.Disabled {
		return nil
	}

	id := c.ID
	if _, err := c.addTask(); err != nil {
		return fmt.Errorf("add task failed for %s cronjob due to %v", id, err)
	}

	c.logger.Info(fmt.Sprintf("Add task completed for %v cronjob", id))

	c.cron.Start()
	c.lastProgress = util.CurrentUnix()

	return nil
}

func (c *Cronjob) retire() error {
	c.cron.Stop()
	id := c.ID
	retireHandler := c.handler.retire
	if retireHandler != nil {
		c.logger.Info(fmt.Sprintf("Cronjob %s retired", id))

		return retireHandler(context.TODO(), c)
	}

	return nil
}

func (c *Cronjob) execute() {
	logger := c.logger
	defer func() {
		if err := recover(); err != nil {
			logger.Error(fmt.Sprintf("Panice when execute cron %s due to %v", c.ID, err))
		}
	}()

	handler := c.handler
	parameterHandler := handler.parameter
	id := c.ID
	ctx, cancelFunc := context.WithTimeout(context.TODO(), c.taskTimeout)
	defer cancelFunc()

	maxNonProgressSec := c.cfg.MaxNonProgressSec
	if util.CurrentUnix()-c.lastProgress > maxNonProgressSec {
		logger.Info(
			fmt.Sprintf("Shuting down cron %s due to not being active after %d (s)", id, maxNonProgressSec),
		)
		if err := c.retire(); err != nil {
			logger.Info(err.Error())
		} else {
			return
		}
	}

	if parameterHandler == nil {
		logger.Debug(fmt.Sprintf("parameterHandler is nil for %s", id))

		return
	}

	params, err := parameterHandler(ctx)
	if err != nil {
		logger.Warn(fmt.Sprintf("Error when getting parameters for %s due to %v", id, err))

		return
	}

	results, err := handler.execute(ctx, params)
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot get results with param %v for %s due to %v", params, id, err))

		return
	}

	resultHandler := handler.result
	if resultHandler != nil {
		err = resultHandler(ctx, results, params)
		if err != nil {
			logger.Warn(fmt.Sprintf("Error handling results for %s: %v", id, err))

			return
		}

		c.lastProgress = util.CurrentUnix()

		return
	}

	logger.Debug(fmt.Sprintf("Result eventhandler is nil for %s", id))
}
