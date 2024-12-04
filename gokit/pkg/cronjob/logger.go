package cronjob

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/log"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(l *zap.Logger) *Logger {
	return &Logger{l}
}

//nolint:revive
func (c *Logger) Info(msg string, keyvals ...any) {
	// c.logger.Info(msg, log.Decode(keyvals)...)
}

func (c *Logger) Error(err error, msg string, keyvals ...any) {
	c.logger.Error(msg, append(log.Decode(keyvals), zap.Error(err))...)
}
