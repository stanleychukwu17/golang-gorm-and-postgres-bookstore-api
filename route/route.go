package route

import (
	"log"
	"net/http"
	"strconv"

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

	// Parse the request body & bind it to the book struct
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

// deleting a book
func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	idStr := context.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 16)
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid ID received",
		})
	}

	// Execute a raw SQL DELETE query
	result := r.DB.Exec("DELETE FROM books WHERE id = ?", id)
	if result.Error != nil {
		return context.Status(http.StatusNotModified).JSON(map[string]interface{}{
			"message": result.Error,
		})
	}

	// return a success message
	return context.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "item deleted successfully",
	})
}

func (r *Repository) updateBook(context *fiber.Ctx) error {
	book_id := context.Params("id")
	book := schema.Book{}

	// check if the book record exists
	if err := r.DB.First(&book, book_id).Error; err != nil {
		return context.Status(http.StatusNotFound).JSON(map[string]interface{}{
			"message": "Book not found",
		})
	}

	// Parse the request body
	if err := context.BodyParser(&book); err != nil {
		return context.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	// update the book
	if err := r.DB.Save(&book).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "Failed to update user",
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book updated successfully",
	})
}

// sets up all the routes
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/createBook", r.CreateBook)
	api.Delete("/deleteBook/:id", r.DeleteBook)
	api.Patch("/updateBook/:id", r.updateBook)
	// api.get("/getBook/:id", r.GetBookByID)
	// api.get("/allBooks", r.GetAllBooks)
}
