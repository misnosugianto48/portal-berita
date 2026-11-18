package handlers

import (
	"errors"
	"fmt"
	"strings"

	"portal-berita/internal/adapters/handlers/requests"
	"portal-berita/internal/adapters/handlers/responses"
	"portal-berita/internal/core/domains/entity"
	"portal-berita/internal/core/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

type AuthHandlersInterface interface {
	Login(c fiber.Ctx) error
}

type authHandlers struct {
	authServices services.AuthServices
}

// Login implements [AuthHandlersInterface].
func (a *authHandlers) Login(c fiber.Ctx) error {
	var code string
	var errorResponse responses.ErrorResponseDefault

	req := requests.LoginRequest{}
	res := responses.SuccessAuthResponse{}

	if err := c.Bind().JSON(&req); err != nil {
		code = "[HANDLER] Login - 1"
		log.Error().Err(err).Msg(code)
		errorResponse.Meta.Status = false
		errorResponse.Meta.Message = bindErrorMessage(err)

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

// bindErrorMessage unwraps validation errors into "Field: tag" messages;
// anything else (e.g. malformed JSON) is returned as-is.
func bindErrorMessage(err error) string {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		msgs := make([]string, 0, len(validationErrors))
		for _, e := range validationErrors {
			msgs = append(msgs, fmt.Sprintf("%s: failed on '%s'", e.Field(), e.Tag()))
		}
		return strings.Join(msgs, ", ")
	}
	return err.Error()
}

func NewAuthHandlers(authServices services.AuthServices) AuthHandlersInterface {
	return &authHandlers{authServices: authServices}
}