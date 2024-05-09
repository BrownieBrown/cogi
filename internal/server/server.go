package server

import (
	"log"
	"net/http"
)

func StartServer(config Config) {
	log.Println("Starting server...")

	err := http.ListenAndServe(config.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
