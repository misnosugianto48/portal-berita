package apps

import (
	"portal-berita/configs"
	"portal-berita/libs/auth"

	"github.com/rs/zerolog/log"
)

func RunServer() error {
	cfg := configs.NewConfig()
	_, err := cfg.ConnectPostgres()
	if err != nil {
		log.Fatal().Msgf("Error connecting database: %v", err)
	}
	cfg.LoadAwsConfig()
	auth.NewJWTService(cfg)
	log.Info().Msg("postgres connected")
	return nil
}