package main

import (
	"todolist-api/config"
)

func main() {
	v := config.NewViper()
	log := config.NewLogger(v)
	validate := config.NewValidation()
	app := config.NewApp(v)
	db := config.NewDatabase(v, log)

	config.NewBootstrap(&config.BootstrapConfig{
		App:      app,
		DB:       db,
		Log:      log,
		Validate: validate,
	})

	panic(app.Listen(":3000"))
}
