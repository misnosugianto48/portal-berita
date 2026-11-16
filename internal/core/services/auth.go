package services

import (
	"context"
	"fmt"
	"portal-berita/configs"
	"portal-berita/internal/adapters/repositories"
	"portal-berita/internal/core/domains/entity"
	"portal-berita/libs/auth"
	"portal-berita/libs/conv"
	"time"

	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt/v5"
)

type AuthServices interface {
	GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.JWTAccessToken, error)
}

type authServices struct {
	authRepositories repositories.AuthRepositories
	cfg              *configs.Config
	jwtToken         auth.JWTConfig
}

// GetUserByEmail implements [AuthServices].
func (a *authServices) GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.JWTAccessToken, error) {
	var result *entity.UserEntity
	var err error
	var code string

	result, err = a.authRepositories.GetUserByEmail(ctx, req)
	if err != nil {
		code = "[SERVICE] GetByUserEmail - 1"
		log.Errorw(code, err)
		return nil, err
	}

	if validPass := conv.ComparePassword(req.Password, result.Password); !validPass {
		code = "[SERVICE] GetByUserEmail - 2"
		log.Errorw(code, "Invalid Password")
		return nil, fmt.Errorf("invalid password")
	}

	jwtData := entity.JWTClaims{UserId: float64(result.Id), RegisteredClaims: jwt.RegisteredClaims{
		NotBefore: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		ID:        string(result.Id),
	}}

	accessToken, expiresAt, err := a.jwtToken.GenerateAccessToken(&jwtData)
	if err != nil {
		code = "[SERVICE] GetByUserEmail - 3"
		log.Errorw(code, err)
		return nil, err
	}

	response := entity.JWTAccessToken{
		AccessToken: accessToken,
		ExpiresAt:   time.Unix(expiresAt, 0),
	}

	return &response, nil
}

func NewAuthService(authRepositories repositories.AuthRepositories, cfg *configs.Config, jwtToken auth.JWTConfig) AuthServices {
	return &authServices{authRepositories: authRepositories, cfg: cfg, jwtToken: jwtToken}
}
