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

type CloudflareR2 struct {
	CloudflareR2BucketName string `json:"cloudflare_r2_bucket_name"`
	CloudflareR2ApiKey     string `json:"cloudflare_r2_api_key"`
	CloudflareR2ApiSecret  string `json:"cloudflare_r2_api_secret"`
	CloudflareR2Token      string `json:"cloudflare_r2_token"`
	CloudflareR2AccountId  string `json:"cloudflare_r2_account_id"`
	CloudflareR2PublicUrl  string `json:"cloudflare_r2_public_url"`
}

type Config struct {
	App          App
	Psql         PsqlDB
	CloudflareR2 CloudflareR2
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort: viper.GetString("APP_PORT"),
			AppEnv:  viper.GetString("APP_ENV"),

			JwtSecretKey: viper.GetString("JWT_SECRET_KEY"),
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
		CloudflareR2: CloudflareR2{

			CloudflareR2BucketName: viper.GetString("CLOUDFLARE_R2_BUCKET_NAME"),
			CloudflareR2ApiKey:     viper.GetString("CLOUDFLARE_R2_API_KEY"),
			CloudflareR2ApiSecret:  viper.GetString("CLOUDFLARE_R2_API_SECRET"),
			CloudflareR2Token:      viper.GetString("CLOUDFLARE_R2_TOKEN"),
			CloudflareR2AccountId:  viper.GetString("CLOUDFLARE_R2_ACCOUNT_ID"),
			CloudflareR2PublicUrl:  viper.GetString("CLOUDFLARE_R2_PUBLIC_URL"),
		},
	}
}
