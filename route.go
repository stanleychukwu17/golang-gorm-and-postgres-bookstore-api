package main

import (
	"log"
	"net/http"

	"github.com/stanleychukwu17/golang-gorm-and-postgres-bookstore-api/schema"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Repository struct
type Repository struct {
	DB *gorm.DB
}

// CreateBook function
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := schema.Book{}

	// Parse the request body, bind it to the book struct
	// the context.BodyParser(&book), is parsing the request body from json to go struct.. Fiber does this internally, be default, Golang does not understand json
	if err := context.BodyParser(&book); err != nil {
		// Log the error and respond with a 400 Bad Request
		log.Println("Error parsing request body:", err)

		return context.Status(http.StatusUnprocessableEntity).JSON(
			map[string]interface{}{
				"message": "Invalid request body",
			},
		)
	}

	// Save the book to the database
	err := r.DB.Create(&book).Error
	if err != nil {
		// Log the error and respond with a 500 Internal Server Error
		log.Println("Error creating book:", err)

		return context.Status(http.StatusInternalServerError).JSON(
			map[string]interface{}{
				"message": "Failed to create book",
			},
		)
	}

	// Respond with a 201 Created status and a success message
	context.Status(http.StatusCreated).JSON(map[string]interface{}{
		"message": "Book created successfully",
	})

	return nil
}

func (r *Repository) setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/createBooks", r.CreateBook)
	// api.delete("/deleteBooks/:id", r.DeleteBook)
	// api.get("/getOneBook/:id", r.GetBookByID)
	// api.get("/allBooks", r.GetAllBooks)
}
