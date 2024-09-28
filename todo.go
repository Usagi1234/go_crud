package main

import (
	"errors"
	"strconv"

	"github.com/Usagi1234/go_crud/models"
	"github.com/gofiber/fiber/v2"
)

func getTodoIndexById(id string) (int, error) {
	for i, t := range todos {
		if t.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("Todo not found")
}

// getbook
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

// getTodo
func getTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

// getbook
func getBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("ไม่เจอจ้า")
}

// getTodo
func getTodo(c *fiber.Ctx) error {
	todoid := c.Params("id")
	for _, todo := range todos {
		if todo.ID == todoid {
			return c.JSON(todo)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("ไม่เจอจ้า")
}

// Book
func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	books = append(books, *book)
	return c.JSON(book)
}

// addTodo
func addTodo(c *fiber.Ctx) error {
	newTodo := new(models.Todo)
	if err := c.BodyParser(newTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	todos = append(todos, *newTodo)
	return c.JSON(newTodo)
}

// updateBook
func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("ไม่เจอหนังสือที่ต้องการ")
}

// updateTodo
func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todoUpdate := new(models.Todo)
	if err := c.BodyParser(todoUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Item = todoUpdate.Item
			return c.JSON(todos[i])
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Todo not found")
}

// deleteBook
func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, book := range books {
		if book.ID == bookId {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("ไม่เจอหนังสือที่ต้องการ")
}

// deleteTodo
func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	index, err := getTodoIndexById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	}

	todos = append(todos[:index], todos[index+1:]...)
	return c.Status(fiber.StatusOK).SendString("Todo deleted")
}
