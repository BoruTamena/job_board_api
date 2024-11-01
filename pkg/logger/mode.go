package logger

import "go.uber.org/zap"

type Mode string

const (
	dev_mode        Mode = "Development"
	production_mode Mode = "Production"
)

func GetConfig(mode Mode) zap.Config {

	switch mode {

	case dev_mode:

		return zap.NewDevelopmentConfig()

	case production_mode:
		return zap.NewProductionConfig()

	default:
		return zap.NewDevelopmentConfig()
	}

}
