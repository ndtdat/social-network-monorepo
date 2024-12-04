package common

import "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

const (
	XTraceIDHeader             = tracer.DefaultTraceIDHeader
	XSpanIDHeader              = tracer.DefaultParentIDHeader
	AuthorizationHeader        = "authorization"
	IdentityIDHeader           = "x-identity-id"
	SessionIDHeader            = "x-session-id"
	IdentityRolesHeader        = "x-identity-roles"
	IdentityMetadataHeader     = "x-identity-metadata"
	IdentityAPIKeyHeader       = "x-identity-api-key"
	SecWebSocketProtocolHeader = "sec-websocket-protocol"
	InternalCallHeader         = "x-internal-call"
	TokenExpiredHeader         = "token-expired"
	HeaderSeparator            = ","
	AcceptLanguage             = "accept-language"
	XForwardedForHeader        = "x-forwarded-for"
	ClientIPHeader             = "client-ip"
	UserAgentHeader            = "user-agent"
	GRPCWebHeader              = "x-grpc-web"
	RicherErrorDetailsHeader   = "details"
	DomainHeader               = "x-currency"
	DeviceIDHeader             = "x-device-id"
	RateLimitRetryAfterHeader  = "retry-after"
	RateLimitMaxAllowedHeader  = "max-allowed"
	RateLimitTypeHeader        = "rate-limit-type"
)

const (
	TraceIDLogName = "dd.trace_id"
	SpanIDLogName  = "dd.span_id"
)
