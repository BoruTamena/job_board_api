package logger

import (
	"go.uber.org/zap"
)

type CfgOptions struct {

	// options
	LevelSplit bool // default false

	// a min log level. default info-lvl
	Atomiclevel string

	// log output path.default stdout
	Out_path string

	// url or folder path to write internal error to.
	// default stderr
	Err_path string

	// enable development mode, defaul value is false
	Dev bool
}

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
