package middlewares

import (
	"portal-berita/configs"
	"portal-berita/internal/adapters/handlers/responses"
	"portal-berita/libs/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type MiddlewaresService interface {
	CheckToken() fiber.Handler
}

type MiddlewareConfig struct {
	authJwt auth.JWTService
}

// CheckToken implements [MiddlewaresService].
func (m *MiddlewareConfig) CheckToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var errorResponse responses.ErrorResponseDefault

		authHandler := c.Get("Authorization")
		if authHandler == "" {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Missing Authorization header"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}

		tokenString := strings.Split(authHandler, "Bearer ")[1]
		claims, err := m.authJwt.VerifyAccessToken(tokenString)
		if err != nil {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Invalid token"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}

		c.Locals("user", claims)

		return c.Next()
	}
}

func NewMiddlewareService(cfg *configs.Config) MiddlewaresService {
	option := new(MiddlewareConfig)
	option.authJwt = auth.NewJWTService(cfg)

	return option
}
