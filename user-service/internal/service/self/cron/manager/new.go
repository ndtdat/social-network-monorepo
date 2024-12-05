package manager

import (
	"context"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/common/pkg/singlepod"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/cronjob"
	"github.com/ndtdat/social-network-monorepo/user-service/config"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/campaign"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/service/self/cron/monitorcampaign"

	"github.com/go-redsync/redsync/v4"
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
	redSync *redsync.Redsync, campaignService *campaign.Service,
) (*Manager, error) {
	var (
		m = &Manager{
			ctx:    ctx,
			logger: logger,
			cfg:    cfg,
		}
		checkerMap = make(map[string]*singlepod.Service)

		monitorCampaignCfg     = cfg.MonitorCampaignCron
		monitorCampaignCache   = monitorCampaignCfg.Cache
		monitorCampaignCronCfg = monitorCampaignCfg.Cron
	)

	/* =================================================================== */
	// MONITOR CAMPAIGN
	/* =================================================================== */
	// 1. Checker
	monitorCampaignID := monitorCampaignCronCfg.ID
	monitorCampaignChecker := singlepod.NewService(
		logger, monitorCampaignID, redisClient, redSync, monitorCampaignCache.Prefix, monitorCampaignCache.Duration,
		monitorCampaignCache.ClientSideCacheDuration, monitorCampaignCfg.KeepAliveDuration,
	)

	// 2. Cron
	monitorCampaignCron, err := monitorcampaign.NewCron(
		ctx, logger, &monitorCampaignCfg, monitorCampaignChecker, campaignService,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create %s cronjob due to: %v", monitorCampaignID, err)
	}

	checkerMap[monitorCampaignID] = monitorCampaignChecker

	/* =================================================================== */
	// CONTAINER
	/* =================================================================== */
	m.Crons = []cronjob.CronContainer{
		monitorCampaignCron,
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
