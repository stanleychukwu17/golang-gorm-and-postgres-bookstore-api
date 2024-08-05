package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (db *gorm.DB, err error) {

	got_err := godotenv.Load()
	if got_err != nil {
		log.Fatalf("Error loading .env file: %v", got_err)
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_TIMEZONE := os.Getenv("DB_TIMEZONE")

	// connect to the postgres database
	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v", DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_TIMEZONE)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return
}
