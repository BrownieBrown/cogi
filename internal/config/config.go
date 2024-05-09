package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	API      APIConfig
	Postgres PostgresConfig
}
type APIConfig struct {
	Port string
}

type PostgresConfig struct {
	Host             string
	Port             string
	User             string
	Driver           string
	Password         string
	Database         string
	ConnectionString string
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

func loadPostgresConfig() PostgresConfig {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		log.Fatal("POSTGRES_HOST is not set")
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		log.Fatal("POSTGRES_PORT is not set")
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		log.Fatal("POSTGRES_USER is not set")
	}

	driver := os.Getenv("POSTGRES_DRIVER")
	if driver == "" {
		log.Fatal("POSTGRES_DRIVER is not set")
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		log.Fatal("POSTGRES_PASSWORD is not set")
	}

	db := os.Getenv("POSTGRES_DB")
	if db == "" {
		log.Fatal("POSTGRES_DB is not set")
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)

	return PostgresConfig{
		Host:             host,
		Port:             port,
		User:             user,
		Driver:           driver,
		Password:         password,
		Database:         db,
		ConnectionString: connectionString,
	}
}
