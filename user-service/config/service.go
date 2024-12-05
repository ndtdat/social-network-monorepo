package config

import "github.com/ndtdat/social-network-monorepo/gokit/pkg/config"

type Service struct {
	Microservices Microservices `mapstructure:"service"`
}

type Microservices struct {
	Purchase config.GRPCClient `mapstructure:"purchase"`
}
