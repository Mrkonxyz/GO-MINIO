package config

import "github.com/spf13/viper"

var AppConfig Config

type Config struct {
	BucketURL         string `mapstructure:"BUCKET_URL"`
	StorageEndpoint   string `mapstructure:"STORAGE_ENDPOINT"`
	StorageAccessKey  string `mapstructure:"STORAGE_ACCESS_KEY"`
	StorageSecretKey  string `mapstructure:"STORAGE_SECRET_KEY"`
	StorageBucketName string `mapstructure:"STORAGE_BUCKET_NAME"`
}

func LoadConFig(path string) (config Config) {
	viper.SetDefault("BUCKET_URL", "http://localhost:9000/")
	viper.SetDefault("STORAGE_ENDPOINT", "localhost:9000")
	viper.SetDefault("STORAGE_ACCESS_KEY", "miniouser")
	viper.SetDefault("STORAGE_SECRET_KEY", "Pa22W0rd")
	viper.SetDefault("STORAGE_BUCKET_NAME", "news-dev")

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	AppConfig = config

	return AppConfig
}
