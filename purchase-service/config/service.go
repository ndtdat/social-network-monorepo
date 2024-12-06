package config

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"time"
)

type Service struct {
	InitFilePath                    InitFilePath                    `mapstructure:"initFilePath"`
	MonitorVoucherConfigurationCron MonitorVoucherConfigurationCron `mapstructure:"monitorVoucherConfigurationCron"`
	MonitorUserVoucherCron          MonitorUserVoucherCron          `mapstructure:"monitorUserVoucherCron"`
	ProvisionVoucherCodeCron        ProvisionVoucherCodeCron        `mapstructure:"provisionVoucherCodeCron"`
	VoucherCfg                      VoucherCfg                      `mapstructure:"voucherCfg"`
}

type InitFilePath struct {
	VoucherCfg       string `mapstructure:"voucherCfg"`
	SubscriptionPlan string `mapstructure:"subscriptionPlan"`
}

type MonitorVoucherConfigurationCron struct {
	Cron              config.Cronjob `mapstructure:"cron"`
	Cache             config.Cache   `mapstructure:"cache"`
	KeepAliveDuration time.Duration  `mapstructure:"keepAliveDuration"`
}

type MonitorUserVoucherCron struct {
	Cron              config.Cronjob `mapstructure:"cron"`
	Cache             config.Cache   `mapstructure:"cache"`
	KeepAliveDuration time.Duration  `mapstructure:"keepAliveDuration"`
}

type ProvisionVoucherCodeCron struct {
	Cron              config.Cronjob `mapstructure:"cron"`
	Cache             config.Cache   `mapstructure:"cache"`
	KeepAliveDuration time.Duration  `mapstructure:"keepAliveDuration"`
	Target            int64          `mapstructure:"target"`
	MaxPerRun         int64          `mapstructure:"maxPerRun"`
}

type VoucherCfg struct {
	Length    int `mapstructure:"length"`
	ExpireDay int `mapstructure:"expireDay"`
}
