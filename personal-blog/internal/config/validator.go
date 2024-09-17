package config

import "github.com/go-playground/validator/v10"

func NewValidate() *validator.Validate {
	return validator.New()
}
