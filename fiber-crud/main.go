package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// s-- Makeshift Database --
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

// e-- Makeshift Database --

func main() {

	// Initialize some sample books
	books = []Book{
		{ID: "1", Title: "Book 1", Author: "Author 1"},
		{ID: "2", Title: "Book 2", Author: "Author 2"},
	}

	// --- Fiber App ---
	app := fiber.New()

	// s-- API ROUTES --
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// --- Fiber App ---
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	// app.Delete("/books/:id", deleteBook)

	log.Println("Running server on http://localhost:3000")
	// e-- API ROUTES --

	app.Listen(":3000")
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

func createBook(c *fiber.Ctx) error {
	// Create a new book object
	var book Book
	if err := c.BodyParser(&book); err != nil {
		return err
	}

	// Add the new book to the books slice
	books = append(books, book)

	// Return the created book as JSON
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookID := c.Params("id", "0")

	var updatedBook Book
	if err := c.BodyParser(&updatedBook); err != nil {
		return err
	}

	// Find the book
	for i := 0; i < len(books); i++ {
		log.Printf("%s == %s", books[i].ID, bookID)
		if books[i].ID == bookID {
			books[i] = updatedBook
			return c.JSON(updatedBook)
		}
	}

	// Didn't find anything
	return c.SendString(fmt.Sprintf("Cannot find book with ID: %s", bookID))
}
