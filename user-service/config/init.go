package config

import (
	config2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"go.uber.org/fx"
	"log"
)

var Module = fx.Options(
	fx.Provide(GetConfigs),
)

func GetConfigs() (*config2.App, *config2.Auth, *Service) {
	var (
		appConfig     *config2.App
		authConfig    *config2.Auth
		serviceConfig *Service
	)

	baseConfigPath := config2.BaseConfigPath
	baseConfigType := config2.BaseConfigType

	if err := config2.LoadConfig(
		baseConfigPath, config2.AppConfigName, baseConfigType, &appConfig,
	); err != nil {
		log.Fatal("Cannot load app config: ", err)
	}

	if err := config2.LoadConfig(
		baseConfigPath, config2.AuthConfigName, baseConfigType, &authConfig,
	); err != nil {
		log.Fatal("Cannot load auth config: ", err)
	}

	if err := config2.LoadConfig(
		baseConfigPath, config2.ServiceConfigName, baseConfigType, &serviceConfig,
	); err != nil {
		log.Fatal("Cannot load service config: ", err)
	}

	return appConfig, authConfig, serviceConfig
}
