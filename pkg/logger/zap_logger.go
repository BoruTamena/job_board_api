package logger

import (
	"go.uber.org/zap"
)

type logger struct {
	logz *zap.Logger
	cfg  CfgOptions
}

func NewLogger(cfg CfgOptions) Logger {

	// get config

	config := GetConfig(cfg.Dev)

	// level split
	return logger{
		cfg: cfg,
	}
}

func (l logger) Info(msg string)  {}
func (l logger) Error(msg string) {}
