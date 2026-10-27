package configs

import "github.com/spf13/viper"

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`

	JwtSecretKey string `json:"jwt_secret_key"`
	JwtIssuer    string `json:"jwt_issuer"`
}

type PsqlDB struct {
	DBHost        string `json:"db_host"`
	DBPort        string `json:"db_port"`
	DBName        string `json:"db_name"`
	DBPassword    string `json:"db_password"`
	DBUser        string `json:"db_user"`
	DBMaxOpenPool int    `json:"db_max_open_pool"`
	DBMaxIddle    int    `json:"db_max_iddle"`
}

type Config struct {
	App  App
	Psql PsqlDB
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort: viper.GetString("APP_PORT"),
			AppEnv:  viper.GetString("APP_ENV"),

			JwtSecretKey: viper.GetString("JWT_SECRET"),
			JwtIssuer:    viper.GetString("JWT_ISSUER"),
		},
		Psql: PsqlDB{
			DBName:        viper.GetString("DATABASE_NAME"),
			DBUser:        viper.GetString("DATABASE_USER"),
			DBHost:        viper.GetString("DATABASE_HOST"),
			DBPort:        viper.GetString("DATABASE_PORT"),
			DBPassword:    viper.GetString("DATABASE_PASSWORD"),
			DBMaxOpenPool: viper.GetInt("DATABASE_MAX_OPEN_POOL_CONNECTION"),
			DBMaxIddle:    viper.GetInt("DATABASE_MAX_IDDLE_CONNECTION"),
		},
	}
}
