package main

import (
	"log"
	"os"

	"github.com/dogukanozdemir/go-postgresql/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment file!")
	}

	postgres_uri := os.Getenv("POSTGRES_URI")
	// host := os.Getenv("DB_HOST")
	// name := os.Getenv("DB_NAME")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASS")
	// port := os.Getenv("DB_PORT")
	// ssl := os.Getenv("DB_SSL")
	db, err := gorm.Open(postgres.Open(postgres_uri), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

}
