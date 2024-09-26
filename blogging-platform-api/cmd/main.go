package main

import (
	"blogging-platform-api/config"
	"fmt"
)

func main() {
	viperConfig := config.NewViper()
	loggerConfig := config.NewLogger(viperConfig)
	fiberConfig := config.NewFiber(viperConfig)
	databaseConfig := config.NewDatabase(viperConfig, loggerConfig)
	validationConfig := config.NewValidation()

	config.NewBootstrap(config.Bootstrap{
		App:      fiberConfig,
		Log:      loggerConfig,
		DB:       databaseConfig,
		Validate: validationConfig,
	})

	host := viperConfig.GetString("SERVER.HOST")
	port := viperConfig.GetInt("SERVER.PORT")
	panic(fiberConfig.Listen(fmt.Sprintf("%s:%d", host, port)))
}
