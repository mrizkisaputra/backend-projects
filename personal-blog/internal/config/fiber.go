package config

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func NewFiber(engine *html.Engine) (app *fiber.App) {
	app = fiber.New(fiber.Config{
		AppName:      "Personal Blog - WebApp",
		Views:        engine,
		ErrorHandler: errorHandler,
	})
	return app
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if ok := errors.As(err, &e); ok {
		code = e.Code
	}

	return ctx.SendStatus(code)
}
