package config

import (
	enum2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/enum"
	"time"
)

type App struct {
	NATS              NATS              `mapstructure:"nats"`
	ServiceCredential ServiceCredential `mapstructure:"serviceCredential"`
	Temporal          Temporal          `mapstructure:"temporal"`
	BNB               BNB               `mapstructure:"bsc"`
	Swagger           Swagger           `mapstructure:"swagger"`
	Environment       enum2.Environment `mapstructure:"environment"`
	KMS               KMS               `mapstructure:"kms"`
	TronGrid          TronGrid          `mapstructure:"tronGrid"`
	TRON              TRON              `mapstructure:"tron"`
	HTTP              HTTP              `mapstructure:"http"`
	Tracing           Tracing           `mapstructure:"tracing"`
	Sentry            Sentry            `mapstructure:"sentry"`
	Celo              Celo              `mapstructure:"celo"`
	BSC               BSC               `mapstructure:"bsc"`
	Arbitrum          Arbitrum          `mapstructure:"arbitrum"`
	Fantom            Fantom            `mapstructure:"fantom"`
	Avalanche         Avalanche         `mapstructure:"avalanche"`
	Ethereum          Ethereum          `mapstructure:"ethereum"`
	Geth              Geth              `mapstructure:"geth"`
	Polygon           Polygon           `mapstructure:"polygon"`
	WebSocket         WebSocket         `mapstructure:"webSocket"`
	GRPC              GRPC              `mapstructure:"grpc"`
	JWT               JWT               `mapstructure:"jwt"`
	ElasticSearch     ElasticSearch     `mapstructure:"elasticSearch"`
	TypeElasticSearch ElasticSearch     `mapstructure:"typedElasticSearch"`
	Aptos             Aptos             `mapstructure:"aptos"`
	MySQL             MySQL             `mapstructure:"mysql"`
	Redis             Redis             `mapstructure:"redis"`
	ClickHouse        ClickHouse        `mapstructure:"clickhouse"`
	VertexAI          VertexAI          `mapstructure:"vertexAI"`
}

type VertexAI struct {
	ProjectID string `json:"projectID"`
	Location  string `json:"location"`
}

type Bucket struct {
	Hostname       string        `mapstructure:"hostname"`
	Name           string        `mapstructure:"name"`
	ExpireDuration time.Duration `mapstructure:"expireDuration"`
}

type ElasticSearch struct {
	Secret        Secret   `mapstructure:"secret"`
	Username      string   `mapstructure:"username"`
	Password      string   `mapstructure:"password"`
	Addresses     []string `mapstructure:"addresses"`
	RetryOnStatus []int    `mapstructure:"retryOnStatus"`
	MaxRetries    int      `mapstructure:"maxRetries"`
	DisableRetry  bool     `mapstructure:"disableRetry"`
	Debug         bool     `mapstructure:"debug"`
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

type WebSocket struct {
	Host             string     `mapstructure:"host"`
	Port             string     `mapstructure:"port"`
	Path             string     `mapstructure:"path"`
	HandshakeTimeout string     `mapstructure:"handshakeTimeout"`
	PingTimeout      string     `mapstructure:"pingTimeout"`
	WorkerPool       WorkerPool `mapstructure:"workerPool"`
	IsPublic         bool       `mapstructure:"isPublic"`
}

type MySQL struct {
	Secret                Secret               `mapstructure:"secret"`
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

//nolint:lll
type ClickHouse struct {
	Secret                      Secret                       `mapstructure:"secret"`
	ConnectionOpenStrategy      enum2.ConnectionOpenStrategy `mapstructure:"connectionOpenStrategy" json:"connectionOpenStrategy"`
	Password                    string                       `mapstructure:"password" json:"password"`
	DB                          string                       `mapstructure:"db" json:"db"`
	Username                    string                       `mapstructure:"username" json:"username"`
	Compress                    enum2.Compress               `mapstructure:"compress" json:"compress"`
	Migration                   DBMigration                  `mapstructure:"migration"`
	Hosts                       []string                     `mapstructure:"hosts" json:"hosts"`
	ClientInfoProduct           []ClickHouseProductInfo      `mapstructure:"clientInfoProduct" json:"clientInfoProduct"`
	DialTimeout                 time.Duration                `mapstructure:"dialTimeout" json:"dialTimeout"`
	ConnectionMaxLifetime       time.Duration                `mapstructure:"connectionMaxLifetime" json:"connectionMaxLifetime"`
	MaxCompressionBufferInBytes uint32                       `mapstructure:"maxCompressionBufferInBytes" json:"maxCompressionBufferInBytes"`
	BlockBufferSize             uint32                       `mapstructure:"blockBufferSize" json:"blockBufferSize"`
	MaxOpenConnections          uint32                       `mapstructure:"maxOpenConnections" json:"maxOpenConnections"`
	MaxIdleConnections          uint32                       `mapstructure:"maxIdleConnections" json:"maxIdleConnections"`
	Debug                       bool                         `mapstructure:"debug" json:"debug"`
}

type ClickHouseProductInfo struct {
	Name    string `mapstructure:"name" json:"name"`
	Version string `mapstructure:"version" json:"version"`
}

type DBMigration struct {
	SourceURL string `mapstructure:"sourceURL"`
	Enabled   bool   `mapstructure:"enabled"`
}

type Redis struct {
	Secret            Secret                 `mapstructure:"secret"`
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

type KMS struct {
	KeyRingID string `mapstructure:"keyRingID"`
}

type Log struct {
	Env string `mapstructure:"env"`
}

type Swagger struct {
	Path string `mapstructure:"path"`
}

type Sentry struct {
	DSN     string `mapstructure:"dsn"`
	Enabled bool   `mapstructure:"enabled"`
}

type BNB struct {
	MasterPrivateKey PrivateKey `mapstructure:"masterPrivateKey" json:"masterPrivateKey"`
}

type Secret struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Key     string `mapstructure:"key"`
}

type SecretEventQueue struct {
	ProjectID      string `mapstructure:"projectID"`
	SubscriptionID string `mapstructure:"subscriptionID"`
	ScheduleSpec   string `mapstructure:"scheduleSpec"`
	Enabled        bool   `mapstructure:"enabled"`
}

type PubSub struct {
	ProjectID      string `mapstructure:"projectID"`
	SubscriptionID string `mapstructure:"subscriptionID"`
	TopicID        string `mapstructure:"topicID"`
}

type PrivateKey struct {
	Secret Secret `mapstructure:"secret"`
	File   string `mapstructure:"file"`
}

type PublicKey struct {
	Secret Secret `mapstructure:"secret"`
	File   string `mapstructure:"file"`
}

type ServiceCredential struct {
	APIKey    string `mapstructure:"apiKey"`
	APISecret string `mapstructure:"apiSecret"`
	Name      string `mapstructure:"name"`
	Secret    Secret `mapstructure:"secret"`
}

type WorkerPool struct {
	MaxWorker int `mapstructure:"maxWorker"`
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

type TemporalConfig struct {
	TimeZone string `mapstructure:"timeZone"`
}

type NATS struct {
	TLS            NATSTLS            `mapstructure:"tls"`
	Secret         Secret             `mapstructure:"secret"`
	Host           string             `mapstructure:"host"`
	Port           string             `mapstructure:"port"`
	Name           string             `mapstructure:"name"`
	Authentication NATSAuthentication `mapstructure:"authentication"`
}

type NATSAuthentication struct {
	Token string `mapstructure:"token"`
}

type NATSTLS struct {
	ClientCert string `mapstructure:"clientCert"`
	ClientKey  string `mapstructure:"clientKey"`
	RootCA     string `mapstructure:"rootCA"`
}

type Temporal struct {
	Host               string `mapstructure:"host"`
	Port               string `mapstructure:"port"`
	Namespace          string `mapstructure:"namespace"`
	NamespaceRetention string `mapstructure:"namespaceRetention"`
	CertPath           string `mapstructure:"certPath"`
	KeyPath            string `mapstructure:"keyPath"`
}

type TemporalCron struct {
	ID           string `mapstructure:"id"`
	TimeZone     string `mapstructure:"timeZone"`
	CronSchedule string `mapstructure:"cronSchedule"`
	Namespace    string `mapstructure:"namespace"`
	Disabled     bool   `mapstructure:"disabled"`
}

type TronGrid struct {
	Schema                           string `mapstructure:"schema"`
	Host                             string `mapstructure:"host"`
	Port                             string `mapstructure:"port"`
	GetNowBlock                      string `mapstructure:"nowBlock"`
	TransactionInfoByID              string `mapstructure:"transactionInfoById"`
	TransactionInfoByContractAddress string `mapstructure:"transactionInfoByContractAddress"`
	GRPCEndpoint                     string `mapstructure:"gRPCEndpoint"`
	APIKey                           string `mapstructure:"apiKey"`
	Contract                         string `mapstructure:"contract"`
	Amount                           string `mapstructure:"amount"`
	Minimum                          string `mapstructure:"minimum"`
	Fee                              Fee    `mapstructure:"fee"`
}

type GethProvider struct {
	Name      string   `mapstructure:"name"`
	Endpoint  string   `mapstructure:"endpoint"`
	RequestID string   `mapstructure:"requestID"`
	APIKey    APIKey   `mapstructure:"apiKey"`
	Keys      []string `mapstructure:"keys"`
	Enabled   bool     `mapstructure:"enabled"`
}

type APIKey struct {
	Header      string `mapstructure:"header"`
	Required    bool   `mapstructure:"required"`
	AppendToURL bool   `mapstructure:"appendToURL"`
}

type Geth struct {
	Fee                       Fee            `mapstructure:"fee"`
	Providers                 []GethProvider `mapstructure:"providers"`
	MaxRetry                  uint32         `mapstructure:"maxRetry"`
	IsConcurrentBlockCrawling bool           `mapstructure:"isConcurrentBlockCrawling"`
}

type Fee struct {
	Contract  string `mapstructure:"contract"`
	Amount    string `mapstructure:"amount"`
	Threshold string `mapstructure:"threshold"`
	Decimals  string `mapstructure:"decimals"`
	BatchSize int    `mapstructure:"batchSize"`
}

type TRON struct {
	TronGrid TronGrid `mapstructure:"tronGrid"`
}

type BSC struct {
	Geth Geth `mapstructure:"geth"`
}

type Ethereum struct {
	Geth Geth `mapstructure:"geth"`
}

type Aptos struct {
	Providers []GethProvider `mapstructure:"providers"`
	GoSDK     Geth           `mapstructure:"goSDK"`
	MaxRetry  uint32         `mapstructure:"maxRetry"`
	//nolint:revive
	ChainId uint8 `mapstructure:"chainId"`
}

type Polygon struct {
	Geth Geth `mapstructure:"geth"`
}

type Avalanche struct {
	Geth Geth `mapstructure:"geth"`
}

type Arbitrum struct {
	Geth Geth `mapstructure:"geth"`
}

type Fantom struct {
	Geth Geth `mapstructure:"geth"`
}

type Celo struct {
	Geth Geth `mapstructure:"geth"`
}

type SonyFlake struct {
	NumInstance int `mapstructure:"numInstance"`
}
