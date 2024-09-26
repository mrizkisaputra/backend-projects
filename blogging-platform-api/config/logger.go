package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(v *viper.Viper) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "Jan 02 2005 15:04:05",
	})
	logger.SetLevel(logrus.Level(v.GetInt("LOGGER.LEVEL")))

	return logger
}
