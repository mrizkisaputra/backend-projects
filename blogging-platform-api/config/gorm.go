package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(v *viper.Viper, log *logrus.Logger) *gorm.DB {
	dbUser := v.GetString("DB.USER")
	dbPassword := v.GetString("DB.PASSWORD")
	dbHost := v.GetString("DB.HOST")
	dbPort := v.GetInt("DB.PORT")
	dbName := v.GetString("DB.NAME")

	// connection to database
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName),
	}), &gorm.Config{
		Logger: logger.New(&logWriter{Logger: log}, logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		}),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.WithError(err).
			WithFields(logrus.Fields{"location": "../config/gorm.go"}).
			Fatal("Failed to connect to database")
	}

	return db
}

type logWriter struct {
	Logger *logrus.Logger
}

func (l *logWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
