package auth

import (
	"fmt"
	"portal-berita/configs"
	"portal-berita/internal/core/domains/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateAccessToken(data *entity.JWTClaims) (string, int64, error)
	VerifyAccessToken(token string) (*entity.JWTClaims, error)
}

type JWTConfig struct {
	signingKey string
	issuer     string
}

// GenerateAccessToken implements [JWTService].
func (o *JWTConfig) GenerateAccessToken(data *entity.JWTClaims) (string, int64, error) {
	now := time.Now().Local()
	expiredAt := now.Add(time.Hour * 24)
	data.ExpiresAt = jwt.NewNumericDate(expiredAt)
	data.Issuer = o.issuer
	data.NotBefore = jwt.NewNumericDate(now)
	createAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	accessToken, err := createAccessToken.SignedString([]byte(o.signingKey))

	if err != nil {
		return "", 0, err
	}

	return accessToken, expiredAt.Unix(), nil
}

// VerifyAccessToken implements [JWTService].
func (o *JWTConfig) VerifyAccessToken(token string) (*entity.JWTClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &entity.JWTClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return []byte(o.signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("Token is not valid")
	}

	claims, ok := parsedToken.Claims.(*entity.JWTClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}

	return claims, nil
}

func NewJWTService(cfg *configs.Config) JWTService {
	option := new(JWTConfig)
	option.signingKey = cfg.App.JwtSecretKey
	option.issuer = cfg.App.JwtIssuer

	return option
}