package server

import (
	"github.com/BrownieBrown/cogi.git/internal/config"
	"log"
	"net/http"
)

func StartServer(config config.Config) {
	log.Println("Starting server...")

	err := http.ListenAndServe(config.API.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
