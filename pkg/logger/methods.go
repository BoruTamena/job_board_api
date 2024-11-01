package logger

import "go.uber.org/zap"

type Logger interface {
	Info(string, zap.Field)
	Error()
}
