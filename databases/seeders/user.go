package seeders

import (
	"portal-berita/internal/core/domains/models"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), 14)
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
