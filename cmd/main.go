package main

import (
	"github.com/BrownieBrown/cogi.git/internal/config"
	"github.com/BrownieBrown/cogi.git/internal/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()
	server.StartServer(cfg)
}
