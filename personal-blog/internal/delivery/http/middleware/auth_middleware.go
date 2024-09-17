package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func NewAuthMiddleware(store *session.Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		if sess.Get("authenticated") == nil {
			return ctx.Redirect("/login")
		}

		return ctx.Next()
	}
}
