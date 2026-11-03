package auth

import (
	"fmt"
	"portal-berita/configs"
	"portal-berita/internal/core/domains/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JsonWebToken interface {
	GenerateToken(data *entity.JsonWebTokenData) (string, int64, error)
	VerifyAccessToken(token string) (*entity.JsonWebTokenData, error)
}

type Options struct {
	signingKey string
	issuer     string
}

// GenerateToken implements [JsonWebToken].
func (o *Options) GenerateToken(data *entity.JsonWebTokenData) (string, int64, error) {
	now := time.Now().Local()
	expiredAt := now.Add(time.Hour * 24)
	data.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expiredAt)
	data.RegisteredClaims.Issuer = o.issuer
	data.RegisteredClaims.NotBefore = jwt.NewNumericDate(now)
	createAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	accessToken, err := createAccessToken.SignedString([]byte(o.signingKey))

	if err != nil {
		return "", 0, err
	}

	return accessToken, expiredAt.Unix(), nil
}

// VerifyAccessToken implements [JsonWebToken].
func (o *Options) VerifyAccessToken(token string) (*entity.JsonWebTokenData, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &entity.JsonWebTokenData{}, func(t *jwt.Token) (interface{}, error) {
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

	claims, ok := parsedToken.Claims.(*entity.JsonWebTokenData)
	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}

	return claims, nil
}

func CreateToken(cfg *configs.Config) JsonWebToken {
	option := new(Options)
	option.signingKey = cfg.App.JwtSecretKey
	option.issuer = cfg.App.JwtIssuer

	return option
}
