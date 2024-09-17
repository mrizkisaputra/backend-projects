package main

import (
	"embed"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
	"personal-blog/internal/config"
)

//go:embed web/templates/*.gohtml
var TEMPLATES embed.FS

func main() {

	templateEngine := config.NewEngine(http.FS(TEMPLATES), ".gohtml")
	app := config.NewFiber(templateEngine)
	log := config.NewLogger()
	validate := config.NewValidate()
	StoreSession := session.New()

	config.NewBootstrap(&config.BootstrapConfig{
		App:          app,
		Log:          log,
		Validate:     validate,
		DirFile:      "./articles",
		StoreSession: StoreSession,
	})

	log.Fatal(app.Listen("localhost:8080"))
}
