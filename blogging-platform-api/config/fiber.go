package config

import (
	"blogging-platform-api/internal/models/dto"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"net/http"
)

func NewFiber(v *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      v.GetString("APP.NAME"),
		ErrorHandler: errorHandler,
	})

	return app
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	var apiResponseError = new(dto.APIResponseError)

	var e *fiber.Error
	if ok := errors.As(err, &e); ok {
		apiResponseError.Status = e.Code
		apiResponseError.Errors = e.Message
	}

	var errValidation validator.ValidationErrors
	if ok := errors.As(err, &errValidation); ok {
		apiResponseError.Status = http.StatusBadRequest
		apiResponseError.Errors = postValidationError(errValidation)
	}

	return ctx.Status(apiResponseError.Status).JSON(apiResponseError, "application/json")
}

func postValidationError(errValidation validator.ValidationErrors) *dto.PostValidationErrResponse {
	errValidationResponse := new(dto.PostValidationErrResponse)
	for _, err := range errValidation {
		switch err.Field() {
		case "Title":
			if err.Tag() == "required" {
				errValidationResponse.Title = append(errValidationResponse.Title, "REQUIRED")
			}
			if err.Tag() == "min" {
				errValidationResponse.Title = append(errValidationResponse.Title, "TOO_SHORT")
			}
			if err.Tag() == "max" {
				errValidationResponse.Title = append(errValidationResponse.Title, "TOO_LONG")
			}
		case "Content":
			if err.Tag() == "required" {
				errValidationResponse.Content = append(errValidationResponse.Content, "REQUIRED")
			}
		case "Category":
			if err.Tag() == "required" {
				errValidationResponse.Category = append(errValidationResponse.Category, "REQUIRED")
			}
			if err.Tag() == "min" {
				errValidationResponse.Category = append(errValidationResponse.Category, "TOO_SHORT")
			}
			if err.Tag() == "max" {
				errValidationResponse.Category = append(errValidationResponse.Category, "TOO_LONG")
			}
			if err.Tag() == "alpha" {
				errValidationResponse.Category = append(errValidationResponse.Category, "MUST_BE_ALPHABET")
			}
		}
	}
	return errValidationResponse
}
