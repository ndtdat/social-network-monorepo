package log

import (
	"go.uber.org/zap"
)

type Generic struct {
	logger *zap.Logger
}

func NewGenericLogger(l *zap.Logger) *Generic {
	return &Generic{l}
}

func (l *Generic) Debug(msg string, keyvals ...any) {
	l.logger.Debug(msg, Decode(keyvals)...)
}

func (l *Generic) Info(msg string, keyvals ...any) {
	l.logger.Info(msg, Decode(keyvals)...)
}

func (l *Generic) Warn(msg string, keyvals ...any) {
	l.logger.Warn(msg, Decode(keyvals)...)
}

func (l *Generic) Error(msg string, keyvals ...any) {
	l.logger.Error(msg, Decode(keyvals)...)
}
