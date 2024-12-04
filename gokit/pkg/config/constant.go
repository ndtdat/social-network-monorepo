package config

const (
	BaseConfigPath         string = "./config"
	BaseOverrideConfigPath string = "./config/override"
	BaseConfigType         string = "yaml"

	AppConfigName                 string = "app"
	ServiceConfigName             string = "service"
	AuthConfigName                string = "auth"
	AccountConfigName             string = "account"
	DockerOverrideConfigName      string = "docker"
	DevelopmentOverrideConfigName string = "development"
	ProductionOverrideConfigName  string = "production"
)
