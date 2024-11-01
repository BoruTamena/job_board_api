package logger

import "go.uber.org/zap"

type Level string

const (
	// loging levels
	InfoLevel  Level = "info" // default
	ErrorLevel Level = "error"
	WarnLevel  Level = "warning"
	PanicLevel Level = "panic"
)

// get atomic level

func GetDefaultAtomicLevel(lvl Level) zap.AtomicLevel {

	atom := zap.NewAtomicLevel()

	switch lvl {
	case WarnLevel:
		atom.SetLevel(zap.WarnLevel)

	case ErrorLevel:
		atom.SetLevel(zap.ErrorLevel)

	case PanicLevel:
		atom.SetLevel(zap.PanicLevel)

	default:
		atom.SetLevel(zap.InfoLevel)

	}

	return atom

}
