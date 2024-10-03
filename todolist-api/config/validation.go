package config

import "github.com/go-playground/validator/v10"

func NewValidation() *validator.Validate {
	return validator.New()
}
