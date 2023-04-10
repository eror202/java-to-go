package web

import (
	"github.com/gofiber/fiber/v2"
	"java-to-go/dto"
	"java-to-go/entity"
	"java-to-go/service"
)

func CreateBook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			entity.Book
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		book := &dto.BookDto{
			Title:     req.Title,
			Author:    req.Author,
			PageCount: req.PageCount,
		}
		id, err := service.CreateBook(book)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(201).JSON(id)
	}
}

func UpdateBook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		type request struct {
			entity.Book
		}
		req := request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		book := &dto.BookDto{
			Title:     req.Title,
			Author:    req.Author,
			PageCount: req.PageCount,
		}
		updateBook, err := service.UpdateBook(book, id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(201).JSON(updateBook)
	}
}

func GetBookById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		book, err := service.GetBookById(id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(200).JSON(book)
	}
}

func DeleteBookById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		_, err := service.DeleteBookById(id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(200).JSON("DELETED")
	}
}
