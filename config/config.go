package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Mode string `mapstructure:"mode"`
	} `mapstructure:"server"`

	RateLimit struct {
		RequestsPerMinute int `mapstructure:"requests_per_minute"`
	} `mapstructure:"rate_limit"`

	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`

	Bet struct {
		MaxAmount float64 `mapstructure:"max_amount"`
		MinAmount float64 `mapstructure:"min_amount"`
	} `mapstructure:"bet"`
}

var AppConfig Config

// LoadConfig loads both .env and .yaml configurations
func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	// Load config.yaml file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}

	// Overwrite settings from .env if they exist
	viper.AutomaticEnv()

	// Unmarshal config data
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %s", err)
	}
}
