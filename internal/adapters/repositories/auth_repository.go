package repositories

import (
	"context"
	"errors"

	"portal-berita/internal/core/domains/entity"
	"portal-berita/internal/core/domains/models"

	"github.com/gofiber/fiber/v3/log"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type AuthRepositories interface {
	GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.UserEntity, error)
}

type authRepositories struct {
	db *gorm.DB
}

// GetUserByEmail implements [AuthRepositories].
func (a *authRepositories) GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.UserEntity, error) {
	modelUser := models.User{}

	if err := a.db.WithContext(ctx).Where("email = ?", req.Email).First(&modelUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		code := "[REPOSITORY] GetUserByEmail - 1"
		log.Errorw(code, err)
		return nil, err
	}

	response := entity.UserEntity{
		Id:       modelUser.Id,
		Name:     modelUser.Name,
		Email:    modelUser.Email,
		Password: modelUser.Password,
	}

	return &response, nil
}

func NewAuthRepositories(db *gorm.DB) AuthRepositories {
	return &authRepositories{db: db}
}
