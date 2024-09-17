package config

import (
	"github.com/sirupsen/logrus"
	"time"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.Level(6))
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	return log
}
