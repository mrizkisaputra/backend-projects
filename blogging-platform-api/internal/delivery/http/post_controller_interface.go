package http

import "github.com/gofiber/fiber/v2"

type PostControllerInterface interface {
	Create(ctx *fiber.Ctx) error

	Get(ctx *fiber.Ctx) error

	Delete(ctx *fiber.Ctx) error

	Update(ctx *fiber.Ctx) error

	List(ctx *fiber.Ctx) error
}
