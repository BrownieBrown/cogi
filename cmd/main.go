package main

import (
	"github.com/BrownieBrown/cogi.git/internal/config"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()
	log.Println("Starting server...")

	log.Println("Server started on port", cfg.API.Port)

	err = http.ListenAndServe(cfg.API.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
