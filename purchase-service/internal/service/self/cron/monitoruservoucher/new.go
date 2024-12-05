package monitoruservoucher

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/common/pkg/singlepod"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/cronjob"
	"github.com/ndtdat/social-network-monorepo/purchase-service/config"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/voucher"
	"go.uber.org/zap"
)

type Cron struct {
	ctx    context.Context
	logger *zap.Logger
	cfg    *config.MonitorUserVoucherCron
	cron   *cronjob.Cronjob

	singlePodChecker *singlepod.Service
	voucherService   *voucher.Service
}

func NewCron(
	ctx context.Context, logger *zap.Logger, cfg *config.MonitorUserVoucherCron,
	singlePodChecker *singlepod.Service, voucherService *voucher.Service,
) (*Cron, error) {
	s := &Cron{
		ctx:              ctx,
		logger:           logger,
		cfg:              cfg,
		singlePodChecker: singlePodChecker,
		voucherService:   voucherService,
	}

	return s, nil
}

func (c *Cron) Start() error {
	go func() {
		c.singlePodChecker.TryToStart(c.ctx, c.startOnSinglePod)
	}()

	return nil
}

func (c *Cron) Stop() error {
	return c.cron.Stop()
}

func (c *Cron) SkitCron() {}

func (c *Cron) ID() string {
	return c.cron.ID
}
