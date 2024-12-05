package config

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"time"
)

type Service struct {
	Microservices       Microservices       `mapstructure:"service"`
	InitFilePath        InitFilePath        `mapstructure:"initFilePath"`
	MonitorCampaignCron MonitorCampaignCron `mapstructure:"monitorCampaignCron"`
}

type InitFilePath struct {
	Campaign string `mapstructure:"campaign"`
}

type Microservices struct {
	Purchase config.GRPCClient `mapstructure:"purchase"`
}

type MonitorCampaignCron struct {
	Cron              config.Cronjob `mapstructure:"cron"`
	Cache             config.Cache   `mapstructure:"cache"`
	KeepAliveDuration time.Duration  `mapstructure:"keepAliveDuration"`
}
