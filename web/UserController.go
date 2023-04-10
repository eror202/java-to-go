package web

import (
	"github.com/gofiber/fiber/v2"
	"java-to-go/dto"
	"java-to-go/entity"
	"java-to-go/service"
	request2 "java-to-go/web/request"
	response2 "java-to-go/web/response"
)

func CreatePerson() fiber.Handler {

	return func(c *fiber.Ctx) error {
		type requestGl struct {
			request2.PersonBookRequest
		}
		var bookId []string
		req := &requestGl{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		us := &dto.PersonDto{
			FullName: req.PersonRequest.FullName,
			Title:    req.PersonRequest.Title,
			Age:      req.PersonRequest.Age,
		}
		id, err := service.CreatePerson(us)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}

		for _, value := range req.PersonBookRequest.BookRequest {
			bo := &dto.BookDto{
				Title:     value.Title,
				Author:    value.Author,
				PageCount: value.PageCount,
				PersonId:  id,
			}
			bId, err := service.CreateBook(bo)
			if err != nil {
				return c.Status(422).JSON(err.Error())
			}
			bookId = append(bookId, bId)
		}
		response := &response2.PersonBookResponse{
			PersonId: id,
			BookId:   bookId,
		}
		return c.Status(201).JSON(response)
	}
}

func GetPersonById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bookId []dto.BookDto
		id := c.Params("id")
		person, err := service.GetPersonById(id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		bookId, err = service.GetBookByPersonId(id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}

		response := &response2.PersonBookResponseObject{
			Person: *person,
			Book:   bookId,
		}
		return c.Status(200).JSON(response)
	}
}

func UpdatePersonById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		type request struct {
			entity.Person
		}
		req := &request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		user := &dto.PersonDto{
			FullName: req.FullName,
			Title:    req.Title,
			Age:      req.Age,
		}
		person, err := service.UpdatePerson(user, id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(201).JSON(person)
	}
}

func DeletePersonById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		_, err := service.DeletePersonById(id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		_, err = service.DeleteBookByPersonId(id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(200).JSON("deleted")
	}
}
