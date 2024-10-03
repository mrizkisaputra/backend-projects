package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(v *viper.Viper) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "Jan 02 2006 15:04:05",
	})
	log.SetLevel(logrus.Level(v.GetInt("LOG.LEVEL")))

	return log
}
