package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
	"todolist-api/internal/models/dto"
	. "todolist-api/internal/usecase"
)

type userController struct {
	UserUseCase UserUseCaseInterface
	Log         *logrus.Logger
}

func NewUserController(usecase UserUseCaseInterface, log *logrus.Logger) UserControllerInterface {
	return &userController{
		UserUseCase: usecase,
		Log:         log,
	}
}

func (u *userController) Register(ctx *fiber.Ctx) error {
	if contentType := ctx.Get("Content-Type"); contentType != "application/json" {
		u.Log.Error("content-type is not application/json")
		return fiber.NewError(fiber.StatusBadRequest, "Content-Type must be application/json")
	}

	requestBody := new(dto.RegisterUserRequestBody)
	if err := ctx.BodyParser(requestBody); err != nil {
		u.Log.WithError(err).Error("error parsing register request body")
		return fiber.ErrInternalServerError
	}

	responseEntity, err := u.UserUseCase.Register(ctx.UserContext(), requestBody)
	if err != nil {
		return err
	}

	responseApi := dto.ApiResponse{
		Status: fiber.StatusCreated,
		Data:   *responseEntity,
		Paging: nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(&responseApi, "application/json")
}

func (u *userController) Login(ctx *fiber.Ctx) error {
	if contentType := ctx.Get("Content-Type"); contentType != "application/json" {
		u.Log.Error("content-type is not application/json")
		return fiber.NewError(fiber.StatusBadRequest, "Content-Type must be application/json")
	}

	requestBody := new(dto.LoginUserRequestBody)
	if err := ctx.BodyParser(&requestBody); err != nil {
		u.Log.WithError(err).Error("error parsing login request body")
		return fiber.ErrInternalServerError
	}

	responseToken, err := u.UserUseCase.Login(ctx.UserContext(), requestBody)
	if err != nil {
		return err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   responseToken.Token,
		Expires: time.Now().Add(24 * time.Hour),
	})
	return ctx.Status(fiber.StatusOK).JSON(responseToken, "application/json")
}
