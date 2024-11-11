package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_USER          string `mapstructure:"DB_USER"`
	DB_PASSWORD      string `mapstructure:"DB_PASSWORD"`
	DB_ROOT_PASSWORD string `mapstructure:"DB_ROOT_PASSWORD"`
	DB_HOST          string `mapstructure:"DB_HOST"`
	DB_PORT          string `mapstructure:"DB_PORT"`
	DB_NAME          string `mapstructure:"DB_NAME"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env") // Look for a file named .env
	viper.AutomaticEnv()        // Enable reading from environment variables

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
