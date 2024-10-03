package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"todolist-api/internal/models/dto"
	. "todolist-api/internal/usecase"
)

func NewAuthMiddleware(userUseCase UserUseCaseInterface, log *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := dto.VerifyUserRequest{AccessToken: ctx.Get("Authorization", "")}
		response, err := userUseCase.VerifyAccessToken(ctx.UserContext(), &request)
		if err != nil {
			log.Debugf("Failed find user by token: %v", request.AccessToken)
			return err
		}

		ctx.Locals("auth", response)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *dto.UserResponse {
	return ctx.Locals("auth").(*dto.UserResponse)
}
