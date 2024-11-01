package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	logz *zap.Logger
	cfg  CfgOptions
}

func NewLogger(cfg CfgOptions) Logger {

	// get atomic
	atom := GetDefaultAtomicLevel(cfg.Atomiclevel)
	// get config
	config := GetConfig(cfg.Dev)
	config.Level = atom
	config.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:     "time",
		LevelKey:    "level",
		MessageKey:  "msg",
		EncodeTime:  zapcore.ISO8601TimeEncoder,
		EncodeLevel: zapcore.CapitalLevelEncoder,
	}

	config.OutputPaths = []string{cfg.Out_path} // output path

	config.ErrorOutputPaths = []string{cfg.Err_path} // error log path
	lg, err := config.Build()

	if err != nil {
		return nil
	}
	// level split
	return logger{
		logz: lg,
		cfg:  cfg,
	}
}

// log info

func (lg logger) Info(ctx context.Context, msg string, fields ...zap.Field) {

	lg.logz.Info(msg, fields...)
}

// log err
func (lg logger) Error(ctx context.Context, msg string, fields ...zap.Field) {}

// log warn
func (lg logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {}

// log panic
func (lg logger) Panic(ctx context.Context, msg string, fields ...zap.Field) {}
