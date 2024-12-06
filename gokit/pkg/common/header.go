package common

const (
	AuthorizationHeader      = "authorization"
	IdentityIDHeader         = "x-identity-id"
	SessionIDHeader          = "x-session-id"
	IdentityRolesHeader      = "x-identity-roles"
	IdentityMetadataHeader   = "x-identity-metadata"
	IdentityAPIKeyHeader     = "x-identity-api-key"
	InternalCallHeader       = "x-internal-call"
	HeaderSeparator          = ","
	AcceptLanguage           = "accept-language"
	XForwardedForHeader      = "x-forwarded-for"
	ClientIPHeader           = "client-ip"
	UserAgentHeader          = "user-agent"
	GRPCWebHeader            = "x-grpc-web"
	RicherErrorDetailsHeader = "details"
	DomainHeader             = "x-currency"
	DeviceIDHeader           = "x-device-id"
)

const (
	TraceIDLogName = "dd.trace_id"
	SpanIDLogName  = "dd.span_id"
)
