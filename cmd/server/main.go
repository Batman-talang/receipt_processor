package main

import (
	"log"
	"net/http"
	"os"
	_ "receipt_processor/docs"
	"receipt_processor/internal/api"
	"receipt_processor/internal/services"
	"receipt_processor/internal/storage"
)

// @title                   Receipt Processor
// @version                 1.0
// @description             This is a simple RESTful service
//
// @contact.name            -
// @contact.url             -
// @contact.email           -
//
// @license.name            MIT
//
// @BasePath                /
// @schemes                 http https
// @description
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	store := storage.New()
	svc := services.New(store)
	router := api.NewRouter(svc)

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
