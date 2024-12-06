package config

import (
	enum2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/enum"
	"time"
)

type App struct {
	Environment enum2.Environment `mapstructure:"environment"`
	HTTP        HTTP              `mapstructure:"http"`
	Tracing     Tracing           `mapstructure:"tracing"`
	GRPC        GRPC              `mapstructure:"grpc"`
	JWT         JWT               `mapstructure:"jwt"`
	MySQL       MySQL             `mapstructure:"mysql"`
	Redis       Redis             `mapstructure:"redis"`
}

type GRPC struct {
	Host                  string `mapstructure:"host"`
	Port                  string `mapstructure:"port"`
	MaxConnectionAge      string `mapstructure:"maxConnectionAge"`
	XDS                   XDS    `mapstructure:"xds"`
	MaxConcurrentStreams  uint32 `mapstructure:"maxConcurrentStreams"`
	AuthenticationEnabled bool   `mapstructure:"authenticationEnabled"`
	AuthorizationEnabled  bool   `mapstructure:"authorizationEnabled"`
	ReflectionEnabled     bool   `mapstructure:"reflectionEnabled"`
}

type XDS struct {
	Host    string `mapstructure:"host"`
	Enabled bool   `mapstructure:"enabled"`
}

type HTTP struct {
	Host               string `mapstructure:"host"`
	Port               string `mapstructure:"port"`
	Path               string `mapstructure:"path"`
	DisableGRPCGateway bool   `mapstructure:"disableGRPCGateway"`
}

type MySQL struct {
	Host                  string               `mapstructure:"host" json:"host"`
	IsolationLevel        enum2.IsolationLevel `mapstructure:"isolationLevel"`
	Port                  string               `mapstructure:"port" json:"port"`
	DB                    string               `mapstructure:"db" json:"db"`
	User                  string               `mapstructure:"user" json:"user"`
	Password              string               `mapstructure:"password" json:"password"`
	Migration             DBMigration          `mapstructure:"migration"`
	ConnectionMaxLifetime time.Duration        `mapstructure:"connectionMaxLifetime" json:"connectionMaxLifetime"`
	MaxOpenConnections    uint32               `mapstructure:"maxOpenConnections" json:"maxOpenConnections"`
	MaxIdleConnections    uint32               `mapstructure:"maxIdleConnections" json:"maxIdleConnections"`
	CreateBatchSize       uint32               `mapstructure:"createBatchSize" json:"createBatchSize"`
	GroupConcatMaxLen     uint32               `mapstructure:"groupConcatMaxLen"`
	InterpolateParams     bool                 `mapstructure:"interpolateParams" json:"interpolateParams"`
	PreparedStatement     bool                 `mapstructure:"preparedStatement" json:"preparedStatement"`
	LogMode               int                  `mapstructure:"logMode"`
	OverwriteGormError    bool                 `mapstructure:"overwriteGormError"`
}

type DBMigration struct {
	SourceURL string `mapstructure:"sourceURL"`
	Enabled   bool   `mapstructure:"enabled"`
}

type Redis struct {
	Host              string                 `mapstructure:"host" json:"host"`
	Port              string                 `mapstructure:"port" json:"port"`
	Password          string                 `mapstructure:"password" json:"password"`
	ClientSideCaching RedisClientSideCaching `mapstructure:"clientSideCaching" json:"clientSideCaching"`
	DB                int                    `mapstructure:"db" json:"db"`
	DisableRetry      bool                   `mapstructure:"disableRetry" json:"disableRetry"`
}

type RedisClientSideCaching struct {
	Prefixes           []string `mapstructure:"prefixes" json:"prefixes"`
	CacheSizeMegaBytes int      `mapstructure:"cacheSizeMegaBytes" json:"cacheSizeMegaBytes"`
	Enabled            bool     `mapstructure:"enabled" json:"enabled"`
	BroadcastMode      bool     `mapstructure:"broadcastMode" json:"broadcastMode"`
}

type Cache struct {
	Prefix                  string        `mapstructure:"prefix"`
	Duration                time.Duration `mapstructure:"duration"`
	ClientSideCacheDuration time.Duration `mapstructure:"clientSideCacheDuration"`
	MaxNumKey               uint32        `mapstructure:"maxNumKey"`
	Local                   bool          `mapstructure:"local"`
}

type JWT struct {
	PrivateKey           PrivateKey       `mapstructure:"privateKey" json:"privateKey"`
	PublicKey            PublicKey        `mapstructure:"publicKey" json:"publicKey"`
	Algorithm            enum2.Algorithm  `mapstructure:"algo" json:"algo"`
	Issuer               string           `mapstructure:"issuer" json:"issuer"`
	SecretEventQueue     SecretEventQueue `mapstructure:"secretEventQueue"`
	AccessTokenDuration  time.Duration    `mapstructure:"accessTokenDuration" json:"accessTokenDuration"`
	RefreshTokenDuration time.Duration    `mapstructure:"refreshTokenDuration" json:"refreshTokenDuration"`
}

type Tracing struct {
	SocketPath  string `mapstructure:"socketPath"`
	ServiceName string `mapstructure:"serviceName"`
	Enabled     bool   `mapstructure:"enabled"`
}

type Log struct {
	Env string `mapstructure:"env"`
}

type SecretEventQueue struct {
	ProjectID      string `mapstructure:"projectID"`
	SubscriptionID string `mapstructure:"subscriptionID"`
	ScheduleSpec   string `mapstructure:"scheduleSpec"`
	Enabled        bool   `mapstructure:"enabled"`
}

type PrivateKey struct {
	File string `mapstructure:"file"`
}

type PublicKey struct {
	File string `mapstructure:"file"`
}

type GRPCClient struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	TLS  bool   `mapstructure:"tls"`
}

type Cronjob struct {
	ID                string        `mapstructure:"id"`
	Spec              string        `mapstructure:"spec"`
	MaxNonProgressSec int64         `mapstructure:"maxNonProgressSec"`
	TaskTimeout       time.Duration `mapstructure:"taskTimeout"`
	Disabled          bool          `mapstructure:"disabled"`
}
