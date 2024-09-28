package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Usagi1234/go_crud/models"
)

// Book struct to hold book data
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book // Slice to store books

var todos = []models.Todo{
	{ID: "1", Item: "Getting a space suit", Completed: false},
	{ID: "2", Item: "Getting a spaceship", Completed: false},
	{ID: "3", Item: "Going to Mars", Completed: false},
}

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Get("/todos", getTodos)          // Read
	app.Get("/todos/:id", getTodo)       // Read
	app.Post("/todos", addTodo)          // Create
	app.Put("/todos/:id", updateTodo)    // Update
	app.Delete("/todos/:id", deleteTodo) // Delete

	app.Listen(":8080")
}
