package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Makeshift Database:
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {

	// Initialize some sample books
	books = []Book{
		{ID: "1", Title: "Book 1", Author: "Author 1"},
		{ID: "2", Title: "Book 2", Author: "Author 2"},
	}

	// --- Fiber App ---
	app := fiber.New()

	// API ROUTES

	// ---> Basic Example
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	// app.Get("/books/:id", getBook)
	// app.Post("/books", createBook)
	// app.Put("/books/:id", updateBook)
	// app.Delete("/books/:id", deleteBook)

	log.Println("Running server on http://localhost:3000")
	app.Listen(":3000")
	// --- Fiber App ---
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookID := c.Params("id", "0")

	// Find the book with the specified ID
	for _, book := range books {
		if book.ID == bookID {
			// Return the book as JSON
			return c.JSON(book)
		}
	}

	// Didn't find anything
	return c.SendString(fmt.Sprintf("Cannot find book with ID: %s", bookID))
}
