package provisionvouchercode

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/cronjob"
)

func (c *Cron) startOnSinglePod(_ context.Context) error {
	c.cron = cronjob.NewCronjob(c.logger, &c.cfg.Cron,
		cronjob.NewHandler(
			func(context.Context) (cronjob.InputParams, error) {
				return nil, nil
			},
			func(ctx context.Context, _ cronjob.InputParams) (cronjob.Result, error) {
				return nil, c.voucherPoolService.Provision(ctx)
			},
			func(_ context.Context, _ cronjob.Result, _ cronjob.InputParams) error {
				return nil
			},
			func(_ context.Context, _ *cronjob.Cronjob) error {
				return nil
			},
		),
	)

	return c.cron.Start()
}
