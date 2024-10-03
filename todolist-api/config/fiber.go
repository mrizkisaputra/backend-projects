package config

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"time"
	"todolist-api/internal/models/dto"
)

func NewApp(v *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      v.GetString("APP.NAME"),
		ErrorHandler: errorHandle,
	})
	return app
}

func errorHandle(ctx *fiber.Ctx, err error) error {
	var e *fiber.Error
	if errors.As(err, &e) {
		apiResponse := dto.ApiResponseError[any]{
			Status:    e.Code,
			Timestamp: time.Now().Format("01-02-2006 15:04:05"),
			Message:   e.Message,
			Error:     nil,
		}
		return ctx.Status(apiResponse.Status).JSON(&apiResponse, "application/json")
	}

	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		apiResponse := dto.ApiResponseError[dto.ApiValidationError]{
			Status:    fiber.StatusBadRequest,
			Timestamp: time.Now().Local().Format("01-02-2006 15:04:05"),
			Message:   "Validation errors",
			Error:     validatedError(validationErr),
		}
		return ctx.Status(apiResponse.Status).JSON(&apiResponse, "application/json")
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON("Internal server error", "application/json")
}

func validatedError(err validator.ValidationErrors) []dto.ApiValidationError {
	var apiValidationErr []dto.ApiValidationError
	fieldTagMessage := map[string]map[string]string{
		"Name": {
			"required": "REQUIRED",
			"max":      "TO_LONG",
		},
		"Email": {
			"required": "REQUIRED",
			"email":    "EMAIL_FORMAT",
			"max":      "TO_LONG",
		},
		"Password": {
			"required": "REQUIRED",
			"min":      "TO_SHORT",
			"max":      "TO_LONG",
		},
		"AccessToken": {
			"required": "REQUIRED",
			"max":      "TO_LONG",
		},
		"Mark": {
			"oneof": "done or in-progress",
			"max":   "TO_LONG",
		},
	}
	for _, e := range err {
		if msg, ok := fieldTagMessage[e.Field()][e.Tag()]; ok {
			apiValidationErr = append(apiValidationErr, dto.ApiValidationError{
				Field:         e.Field(),
				RejectedValue: e.Value(),
				Message:       msg,
			})
		}
	}

	return apiValidationErr
}
