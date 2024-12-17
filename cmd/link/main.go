package main

import (
	"log"
	"net/http"

	"github.com/adii1203/link/internal/handlers/link"
	"github.com/adii1203/link/internal/initializers"
)

func main() {
	store, err := initializers.New()
	if err != nil {
		log.Fatal("db initialization error")
	}

	router := http.NewServeMux()

	router.HandleFunc("POST /api/links", link.New(store))

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("error server starting")
	}
}
