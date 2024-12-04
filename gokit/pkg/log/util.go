package log

import (
	"context"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/common"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const RFC3339Millis = "2006-01-02T15:04:05.000Z07:00"

func Logger() *zap.Logger {
	return _logger
}

func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return _logger
	}

	traceID, spanID := getDDTraceInfo(ctx)

	return _logger.With(
		zap.String(common.TraceIDLogName, traceID),
		zap.String(common.SpanIDLogName, spanID),
	)
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]", t.UTC().Format(RFC3339Millis)))
}

func Decode(keyvals []any) []zap.Field {
	var fields []zap.Field
	nPair := len(keyvals)
	for i := 0; i < nPair/2; i++ {
		fields = append(fields, zap.Reflect(keyvals[i*2].(string), keyvals[i*2+1]))
	}

	return fields
}
