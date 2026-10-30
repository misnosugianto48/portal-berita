package apps

import (
	"portal-berita/configs"

	"github.com/rs/zerolog/log"
)

func RunServer() error {
	cfg := configs.NewConfig()
	_, err := cfg.ConnectPostgres()
	if err != nil {
		log.Fatal().Msgf("Error connecting database: %v", err)
	}

	log.Info().Msg("postgres connected")
	return nil
}
