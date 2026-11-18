package handlers

import (
	"portal-berita/internal/adapters/handlers/requests"
	"portal-berita/internal/adapters/handlers/responses"
	"portal-berita/internal/core/domains/entity"
	"portal-berita/internal/core/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
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
	var err error
	var code string
	var errorResponse responses.ErrorResponseDefault

	req := requests.LoginRequest{}
	res := responses.SuccessAuthResponse{}

	if err := c.BodyParser(&req); err != nil {
		code = "[HANDLER] Login - 1"
		log.Error().Err(err).Msg(code)
		errorResponse.Meta.Status = false
		errorResponse.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	if err = validate.Struct(req); err != nil {
		code = "[HANDLER] Login - 2"
		log.Error().Err(err).Msg(code)
		errorResponse.Meta.Status = false
		errorResponse.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	reqLogin := entity.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := a.authServices.GetUserByEmail(c.Context(), reqLogin)
	if err != nil {
		code = "[HANDLER] Login - 3"
		log.Error().Err(err).Msg(code)
		errorResponse.Meta.Status = false
		errorResponse.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
	}

	res.Meta.Status = true
	res.Meta.Message = "Login successful"
	res.AccessToken = result.AccessToken
	res.ExpiresAt = result.ExpiresAt.Unix()

	return c.JSON(res)
}

func NewAuthHandlers(authServices services.AuthServices) AuthHandlersInterface {
	return &authHandlers{authServices: authServices}
}
