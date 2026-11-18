package configs

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	err := v.validate.Struct(out)
	if err == nil {
		return nil
	}

	valErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	var messages []string
	for _, e := range valErrors {
		messages = append(messages, formatError(e))
	}

	return fmt.Errorf("%s", strings.Join(messages, ", "))
}

func formatError(e validator.FieldError) string {
	field := e.Field()
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must not exceed %s characters", field, e.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

// NewStructValidator adapts go-playground/validator to fiber's StructValidator.
func NewStructValidator() fiber.StructValidator {
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &structValidator{validate: v}
}
