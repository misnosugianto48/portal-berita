package handlers

import (
	"portal-berita/internal/core/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

type AuthHandlersInterface interface {
	Login(c *fiber.Ctx) error
}

type authHandlers struct {
	authServices services.AuthServices
}

// Login implements [AuthHandlersInterface].
func (a *authHandlers) Login(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewAuthHandlers(authService services.AuthServices) AuthHandlersInterface {
	return &authHandlers{authServices: authService}
}
