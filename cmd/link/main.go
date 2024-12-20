package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adii1203/link/internal/handlers/link"
	"github.com/adii1203/link/internal/initializers"
	"github.com/adii1203/link/internal/middlewares"
)

func main() {
	store, err := initializers.New()
	if err != nil {
		log.Fatal("db initialization error")
	}

	router := http.NewServeMux()

	router.Handle("POST /api/links", middlewares.ValidatePayload(link.New(store)))
	router.Handle("GET /api/links/metadata", link.Metadata())
	router.Handle("GET /{slug}", middlewares.IsCrawler(link.Redirect(store)))

	port := os.Getenv("PORT")
	server := http.Server{
		Addr:    os.Getenv(":" + port),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("error server starting")
	}
}
