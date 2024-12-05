package manager

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/ndtdat/social-network-monorepo/common/pkg/singlepod"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/cronjob"
	"github.com/ndtdat/social-network-monorepo/purchase-service/config"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/cron/monitoruservoucher"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/cron/monitorvoucherconfiguration"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/cron/provisionvouchercode"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/voucher"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/service/self/voucherpool"
	"github.com/redis/rueidis"
	"go.uber.org/zap"
)

type Manager struct {
	ctx          context.Context
	cfg          *config.Service
	logger       *zap.Logger
	singlePodMap map[string]*singlepod.Service
	Crons        []cronjob.CronContainer
}

//nolint:funlen,maintidx
func NewManager(
	ctx context.Context, cfg *config.Service, logger *zap.Logger, redisClient rueidis.Client,
	redSync *redsync.Redsync, voucherService *voucher.Service, voucherPoolService *voucherpool.Service,
) (*Manager, error) {
	var (
		m = &Manager{
			ctx:    ctx,
			logger: logger,
			cfg:    cfg,
		}
		checkerMap = make(map[string]*singlepod.Service)

		monitorVoucherConfigurationCfg     = cfg.MonitorVoucherConfigurationCron
		monitorVoucherConfigurationCache   = monitorVoucherConfigurationCfg.Cache
		monitorVoucherConfigurationCronCfg = monitorVoucherConfigurationCfg.Cron

		monitorUserVoucherCfg     = cfg.MonitorUserVoucherCron
		monitorUserVoucherCache   = monitorUserVoucherCfg.Cache
		monitorUserVoucherCronCfg = monitorUserVoucherCfg.Cron

		provisionVoucherCodeCfg     = cfg.ProvisionVoucherCodeCron
		provisionVoucherCodeCache   = provisionVoucherCodeCfg.Cache
		provisionVoucherCodeCronCfg = provisionVoucherCodeCfg.Cron
	)

	/* =================================================================== */
	// MONITOR VOUCHER CONFIGURATION
	/* =================================================================== */
	// 1. Checker
	monitorVoucherConfigurationID := monitorVoucherConfigurationCronCfg.ID
	monitorVoucherConfigurationChecker := singlepod.NewService(
		logger, monitorVoucherConfigurationID, redisClient, redSync, monitorVoucherConfigurationCache.Prefix,
		monitorVoucherConfigurationCache.Duration, monitorVoucherConfigurationCache.ClientSideCacheDuration,
		monitorVoucherConfigurationCfg.KeepAliveDuration,
	)

	// 2. Cron
	monitorVoucherConfigurationCron, err := monitorvoucherconfiguration.NewCron(
		ctx, logger, &monitorVoucherConfigurationCfg, monitorVoucherConfigurationChecker, voucherService,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create %s cronjob due to: %v", monitorVoucherConfigurationID, err)
	}

	checkerMap[monitorVoucherConfigurationID] = monitorVoucherConfigurationChecker

	/* =================================================================== */
	// MONITOR USER VOUCHER
	/* =================================================================== */
	// 1. Checker
	monitorUserVoucherID := monitorUserVoucherCronCfg.ID
	monitorUserVoucherChecker := singlepod.NewService(
		logger, monitorUserVoucherID, redisClient, redSync, monitorUserVoucherCache.Prefix,
		monitorUserVoucherCache.Duration, monitorUserVoucherCache.ClientSideCacheDuration,
		monitorUserVoucherCfg.KeepAliveDuration,
	)

	// 2. Cron
	monitorUserVoucherCron, err := monitoruservoucher.NewCron(
		ctx, logger, &monitorUserVoucherCfg, monitorUserVoucherChecker, voucherService,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create %s cronjob due to: %v", monitorUserVoucherID, err)
	}

	checkerMap[monitorUserVoucherID] = monitorUserVoucherChecker

	/* =================================================================== */
	// PROVISION VOUCHER CODE
	/* =================================================================== */
	// 1. Checker
	provisionVoucherCodeID := provisionVoucherCodeCronCfg.ID
	provisionVoucherCodeChecker := singlepod.NewService(
		logger, provisionVoucherCodeID, redisClient, redSync, provisionVoucherCodeCache.Prefix,
		provisionVoucherCodeCache.Duration, provisionVoucherCodeCache.ClientSideCacheDuration,
		provisionVoucherCodeCfg.KeepAliveDuration,
	)

	// 2. Cron
	provisionVoucherCodeCron, err := provisionvouchercode.NewCron(
		ctx, logger, &provisionVoucherCodeCfg, provisionVoucherCodeChecker, voucherPoolService,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create %s cronjob due to: %v", provisionVoucherCodeID, err)
	}

	checkerMap[provisionVoucherCodeID] = provisionVoucherCodeChecker

	/* =================================================================== */
	// CONTAINER
	/* =================================================================== */
	m.Crons = []cronjob.CronContainer{
		monitorVoucherConfigurationCron, monitorUserVoucherCron, provisionVoucherCodeCron,
	}
	m.singlePodMap = checkerMap

	return m, nil
}

func (m *Manager) Start() error {
	for _, c := range m.Crons {
		if err := c.Start(); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) Stop() error {
	for _, c := range m.Crons {
		if err := c.Stop(); err != nil {
			m.logger.Error(fmt.Sprintf("Cannot stop cron %s due to %v", c.ID(), err))
		}
	}

	singlepod.DeleteSingPods(m.logger, m.singlePodMap)

	return nil
}

func (m *Manager) SkitCronjobManager() {}
