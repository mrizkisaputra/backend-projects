package test

import (
	"blogging-platform-api/config"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var viperConfig *viper.Viper
var loggerConfig *logrus.Logger
var app *fiber.App
var db *gorm.DB
var validationConfig *validator.Validate

func init() {
	viperConfig = config.NewViper()
	loggerConfig = config.NewLogger(viperConfig)
	app = config.NewFiber(viperConfig)
	db = config.NewDatabase(viperConfig, loggerConfig)
	validationConfig = config.NewValidation()

	config.NewBootstrap(config.Bootstrap{
		App:      app,
		Log:      loggerConfig,
		DB:       db,
		Validate: validationConfig,
	})
}
