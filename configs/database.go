package configs

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectPostgres() (*Postgres, error) {
	dbConnect := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Psql.DBUser, cfg.Psql.DBPassword, cfg.Psql.DBHost, cfg.Psql.DBPort, cfg.Psql.DBName)

	db, err := gorm.Open(postgres.Open(dbConnect), &gorm.Config{})
	if err != nil {
		log.Error().
			Err(err).
			Msg("[ConnectionPostgres-1] failed to connect to database" + cfg.Psql.DBHost)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {

		log.Error().
			Err(err).
			Msg("[ConnectionPostgres-2] failed to get database connection")
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.Psql.DBMaxOpenPool)
	sqlDB.SetMaxIdleConns(cfg.Psql.DBMaxIddle)

	return &Postgres{DB: db}, nil
}
