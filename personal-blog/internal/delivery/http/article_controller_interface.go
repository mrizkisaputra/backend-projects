package http

import "github.com/gofiber/fiber/v2"

type ArticleControllerInterface interface {
	Home(ctx *fiber.Ctx) error
	PageAdminHome(ctx *fiber.Ctx) error
	PageNewArticle(ctx *fiber.Ctx) error
	PageLogin(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	PageEditArticle(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
