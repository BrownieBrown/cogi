package config

import (
	"log"
	"os"
)

type Config struct {
	API APIConfig
}
type APIConfig struct {
	Port string
}

func LoadConfig() Config {
	apiConfig := loadAPIConfig()
	return Config{
		API: apiConfig,
	}
}

func loadAPIConfig() APIConfig {
	port := os.Getenv("API_PORT")
	if port == "" {
		log.Fatal("API_PORT is not set")
	}

	return APIConfig{Port: port}
}
