package main

import (
	"embed"
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
	"net/http"
)

//go:embed web/*.html
var WEB embed.FS

////go:embed assets/css/*.css
//var ASSETS_CSS embed.FS
//
////go:embed assets/js/*.js
//var ASSETS_JS embed.FS

var (
	host = kingpin.Flag("host", "host for the run app").Short('h').Required().String()
	port = kingpin.Flag("port", "port for the run app").Short('p').Required().Int()
)

func main() {
	kingpin.Parse()

	engine := html.NewFileSystem(http.FS(WEB), ".html")
	app := fiber.New(fiber.Config{Views: engine})
	handlers(app)
	handlerStatic(app)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", *host, *port)))
}

func handlers(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("web/index", nil)
	})

	app.Get("/_length.html", func(ctx *fiber.Ctx) error {
		return ctx.Render("web/_length", nil)
	})

	app.Get("/_weight.html", func(ctx *fiber.Ctx) error {
		return ctx.Render("web/_weight", nil)
	})

	app.Get("/_temperature.html", func(ctx *fiber.Ctx) error {
		return ctx.Render("web/_temperature", nil)
	})
}

func handlerStatic(app *fiber.App) {
	app.Static("/static/assets/css/output.css", "./assets/css/output.css")
	app.Static("/static/assets/js/jquery-3.7.1.min.js", "./assets/js/jquery-3.7.1.min.js")
	app.Static("/static/assets/js/main.js", "./assets/js/main.js")
	app.Static("/static/assets/js/length_converter.js", "./assets/js/length_converter.js")
	app.Static("/static/assets/js/weight_converter.js", "./assets/js/weight_converter.js")
	app.Static("/static/assets/js/temperature_converter.js", "./assets/js/temperature_converter.js")
}
