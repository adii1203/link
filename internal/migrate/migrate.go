package main

import (
	"log"

	"github.com/adii1203/link/internal/initializers"
	"github.com/adii1203/link/internal/models"
)

func main() {

	pgx, err := initializers.New()
	if err != nil {
		log.Fatal("error while initializing database")
	}

	err = pgx.Db.AutoMigrate(&models.Link{})
	if err != nil {
		log.Fatal("db migration error")
	}
}
