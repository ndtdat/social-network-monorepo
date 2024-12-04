package config

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/set"
	"strings"

	"github.com/dolthub/swiss"
)

type Auth struct {
	Rule AuthRule `mapstructure:"rule"`
}

type AuthRule struct {
	ServicePath  string      `mapstructure:"servicePath"`
	AuthRoutes   []AuthRoute `mapstructure:"authRoutes"`
	PublicRoutes []string    `mapstructure:"publicRoutes"`
}

type AuthRoute struct {
	Endpoint        string   `mapstructure:"endpoint"`
	AccessibleRoles []string `mapstructure:"accessibleRoles"`
	APIKeyHeader    string   `mapstructure:"apiKeyHeader"`
}

func (ac *Auth) AuthRouteMap() *swiss.Map[string, []string] {
	accessibleRoles := swiss.NewMap[string, []string](42)

	for _, route := range ac.Rule.AuthRoutes {
		accessibleRoles.Put(
			strings.Join([]string{ac.Rule.ServicePath, route.Endpoint}, "/"), route.AccessibleRoles,
		)
	}

	return accessibleRoles
}

func (ac *Auth) PublicRouteMap() *set.Set[string] {
	accessibleRoles := set.New[string]()

	for _, route := range ac.Rule.PublicRoutes {
		accessibleRoles.Add(route)
	}

	return accessibleRoles
}

func (ac *Auth) APIKeyHeaderMap() *swiss.Map[string, string] {
	apiKeyHeader := swiss.NewMap[string, string](42)

	for _, route := range ac.Rule.AuthRoutes {
		apiKeyHeader.Put(
			strings.Join([]string{ac.Rule.ServicePath, route.Endpoint}, "/"), route.APIKeyHeader,
		)
	}

	return apiKeyHeader
}
