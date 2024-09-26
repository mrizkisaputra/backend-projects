package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewLoggerMiddleware(log *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		// setelah eksekusi handler route di eksekusi
		if err != nil {
			if e, ok := err.(*fiber.Error); ok {
				ctx.Status(e.Code) // Set status code dari error Fiber
			} else {
				ctx.Status(fiber.StatusInternalServerError)
			}
		}

		log.WithFields(logrus.Fields{
			"method":        ctx.Method(),
			"path":          ctx.Path(),
			"status-code":   ctx.Response().StatusCode(),
			"user-agent":    ctx.Get("User-Agent"),
			"response-size": len(ctx.Response().Body()),
			"error":         err,
		}).Info("request completed")
		return err
	}
}
