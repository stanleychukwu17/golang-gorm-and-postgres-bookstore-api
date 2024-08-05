package main

import (
	"fmt"
	"log"
	"os"

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

	// db, err := storage.NewConnection(config)
	// if err != nil {
	// 	log.Fatalf("Error connecting to database: %v", err)
	// }

	// Create new Fiber instance
	app := fiber.New()

	// Create GET route on path "/"
	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Hello, World!")
	})

	// Start server
	app.Listen(fmt.Sprintf(":%s", port))

	// Print message to console
	fmt.Println("Server running on port", port)
}
