package apps

import (
	"portal-berita/configs"

	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := configs.NewConfig()
	_, err := cfg.ConnectPostgres()
	if err != nil {
		log.Fatal().Msgf("Error connecting database: %v", err)
		return
	}
}
