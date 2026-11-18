package configs

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type structValidator struct {
	validate *validator.Validate
}

// Validate implements [fiber.StructValidator].
func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

// NewStructValidator adapts go-playground/validator to fiber's StructValidator.
func NewStructValidator() fiber.StructValidator {
	return &structValidator{validate: validator.New()}
}
