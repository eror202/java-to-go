package web

import (
	"github.com/gofiber/fiber/v2"
	"java-to-go/entity"
	"java-to-go/service"
)

func CreateBook(services *service.BookServiceImpl) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			entity.Book
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		book := &entity.Book{
			Title:     req.Title,
			Author:    req.Author,
			PageCount: req.PageCount,
		}
		id, err := services.CreateBook(book)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(201).JSON(id)
	}
}

func UpdateBook(services *service.BookServiceImpl) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		type request struct {
			entity.Book
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		book := &entity.Book{
			Title:     req.Title,
			Author:    req.Author,
			PageCount: req.PageCount,
		}
		updateBook, err := services.UpdateBook(book, id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(201).JSON(updateBook)
	}
}

func GetBookById(services *service.BookServiceImpl) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		book, err := services.GetBookById(id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(200).JSON(book)
	}
}

func DeleteBookById(services *service.BookServiceImpl) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		_, err := services.DeleteBookById(id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(200).JSON("DELETED")
	}
}
