package configs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/rs/zerolog/log"
)

func (cfg Config) LoadAwsConfig() aws.Config {
	config, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.CloudflareR2.CloudflareR2ApiKey, cfg.CloudflareR2.CloudflareR2ApiSecret, "",
		)), awsConfig.WithRegion("auto"))
	if err != nil {
		log.Fatal().Err(err).Msgf("Unable to load aws config: %v", err)
	}
	log.Info().Msg("AWS config load successfully!")

	return config
}
