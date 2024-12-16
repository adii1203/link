package main

import (
	"github.com/adii1203/link/internal/handlers/link"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /api/links", link.New())

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("error server starting")
	}
}
