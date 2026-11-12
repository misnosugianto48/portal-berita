package apps

import (
	"context"
	"os"
	"os/signal"
	"portal-berita/configs"
	"portal-berita/libs/auth"
	"portal-berita/libs/middlewares"
	"portal-berita/libs/paginations"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

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
	middlewares.NewMiddlewareService(cfg)
	paginations.NewPaginator()

	log.Info().Msg("postgres connected")

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Portal App v1.0.1",
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip} ${status} - ${latency} ${method} ${path}\n",
	}))

	_ = app.Group("/api")

	if cfg.App.AppPort == "" {
		cfg.App.AppPort = os.Getenv("APP_PORT")
	}

	go func() {
		err := app.Listen(":" + cfg.App.AppPort)
		if err != nil {
			log.Fatal().Msgf("error starting server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit

	log.Info().Msg("shutdown signal received, waiting for server to close")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatal().Msgf("error shutting down server: %v", err)
	}

	return nil
}