package logger

import (
	"context"

	"go.uber.org/zap"
)

type Logger interface {
	// log info
	Info(ctx context.Context, msg string, fields ...zap.Field)
	// log err
	Error(ctx context.Context, msg string, fields ...zap.Field)
	// log warn
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	// log panic
	Panic(ctx context.Context, msg string, fields ...zap.Field)
}
