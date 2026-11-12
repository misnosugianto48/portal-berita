package seeders

import (
	"portal-berita/internal/core/domains/models"
	"portal-berita/libs/conv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	hashedPassword, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg("Error hashing password")
	}

	admin := models.User{
		Name:     "Admin",
		Email:    "admin@email.com",
		Password: string(hashedPassword),
	}

	if err := db.FirstOrCreate(&admin, models.User{Email: "admin@email.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("Error seeding admin")
	} else {
		log.Info().Msg("Admin role seeding successfuly!")
	}
}
