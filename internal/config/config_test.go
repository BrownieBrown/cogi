package config

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up
	expectedPort := ":8080"

	// Execute
	config := LoadConfig()

	// Assert
	if config.API.Port != expectedPort {
		t.Errorf("LoadConfig() failed, expected %v, got %v", expectedPort, config.API.Port)
	}
}
