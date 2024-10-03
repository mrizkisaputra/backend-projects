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
	dbPass := v.GetString("DB.PASSWORD")
	dbHost := v.GetString("DB.HOST")
	dbPort := v.GetInt("DB.PORT")
	dbName := v.GetString("DB.NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.WithError(err).Error("Failed to connect to database")
	}

	return db
}
