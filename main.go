package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stanleychukwu17/golang-gorm-and-postgres-bookstore-api/route"
	"github.com/stanleychukwu17/golang-gorm-and-postgres-bookstore-api/schema"
	"github.com/stanleychukwu17/golang-gorm-and-postgres-bookstore-api/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Access environment variables
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatalf("PORT environment variable is required but not set")
	}

	db, err := storage.NewConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	} else {
		fmt.Println("Connected to database", db)
	}

	// Create new Fiber instance
	app := fiber.New()

	// Create GET route on path "/"
	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Hello, World!")
	})

	// set_up a new Repository
	r := &route.Repository{
		DB: db,
	}

	// Automatically migrate your schema
	err = db.AutoMigrate(&schema.User{}, &schema.Book{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	r.SetupRoutes(app)

	// Start server
	app.Listen(fmt.Sprintf(":%s", port))

	// Print message to console
	fmt.Println("Server running on port", port)
}
