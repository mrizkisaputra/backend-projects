package http

import "github.com/gofiber/fiber/v2"

type UserControllerInterface interface {
	Register(ctx *fiber.Ctx) error

	Login(ctx *fiber.Ctx) error
}
