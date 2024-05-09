package main

import (
	"database/sql"
	"github.com/BrownieBrown/cogi.git/internal/config"
	"github.com/BrownieBrown/cogi.git/internal/database/postgres"
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
	conn, err := postgres.Client(cfg)
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	server.StartServer(cfg)
}
