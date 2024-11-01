package logger

import (
	"go.uber.org/zap"
)

type logger struct {
	logz *zap.Logger
	cfg  CfgOptions
}

func NewLogger(cfg CfgOptions) Logger {

	// level split
	return logger{}
}

func (l logger) Info(msg string)  {}
func (l logger) Error(msg string) {}
