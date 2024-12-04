package log

import (
	"context"
	"fmt"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	common2 "github.com/ndtdat/social-network-monorepo/gokit/pkg/common"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/enum"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"strconv"
)

var (
	_logger *zap.Logger
)

func New(cfg *config.App) (*zap.Logger, []grpc_zap.Option, error) {
	var (
		zapLogger  *zap.Logger
		err        error
		sentryHook zap.Option
		zapCfg     zap.Config
		env        = cfg.Environment
	)

	zapCfg = zap.Config{
		Development:       true,
		Encoding:          "console",
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:       []string{"stderr"},
		DisableStacktrace: false,

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			TimeKey:       "time",
			LevelKey:      "level",
			CallerKey:     "caller",
			StacktraceKey: "stacker",
			EncodeCaller:  zapcore.FullCallerEncoder,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    SyslogTimeEncoder,
		},
	}

	switch env {
	case enum.Environment_LOCAL:
	case enum.Environment_DEVELOPMENT:
		zapCfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	case enum.Environment_PRODUCTION:
		zapCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		zapCfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	default:
		return nil, nil, fmt.Errorf("cannot find any log env for %s", env)
	}

	if sentryHook != nil {
		zapLogger, err = zapCfg.Build(sentryHook)
	} else {
		zapLogger, err = zapCfg.Build()
	}
	if err != nil {
		return nil, nil, err
	}

	_logger = zapLogger
	defer _logger.Sync() //nolint:errcheck

	zapOpts := []grpc_zap.Option{
		grpc_zap.WithLevels(func(code codes.Code) zapcore.Level {
			switch code { //nolint:exhaustive
			case codes.OK:
				return zapcore.InfoLevel

			case codes.Internal:
				return zapcore.ErrorLevel

			default:
				return zapcore.DebugLevel
			}
		}),
		grpc_zap.WithDecider(func(fullMethodName string, err error) bool {
			if err == nil && fullMethodName == common2.HealthCheckMethod {
				return false
			}

			return true
		}),
		grpc_zap.WithMessageProducer(
			func(ctx context.Context, msg string, level zapcore.Level, _ codes.Code, _ error, _ zapcore.Field) {
				traceID, spanID := getDDTraceInfo(ctx)
				ctxzap.Extract(ctx).Check(level, msg).Write(
					zap.String(common2.InternalCallHeader, util.FieldFromIncomingCtx(ctx, common2.InternalCallHeader)),
					zap.String(common2.IdentityIDHeader, util.FieldFromIncomingCtx(ctx, common2.IdentityIDHeader)),
					zap.String(common2.SessionIDHeader, util.FieldFromIncomingCtx(ctx, common2.SessionIDHeader)),
					zap.String(common2.ClientIPHeader, util.FieldFromIncomingCtx(ctx, common2.ClientIPHeader)),
					zap.String(common2.UserAgentHeader, util.FieldFromIncomingCtx(ctx, common2.UserAgentHeader)),
					zap.String(common2.TraceIDLogName, traceID),
					zap.String(common2.SpanIDLogName, spanID),
				)
			}),
	}

	return zapLogger, zapOpts, nil
}

func getDDTraceInfo(ctx context.Context) (string, string) {
	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		return "", ""
	}

	spanCtx := span.Context()

	return strconv.FormatUint(spanCtx.TraceID(), 10), strconv.FormatUint(spanCtx.SpanID(), 10)
}
