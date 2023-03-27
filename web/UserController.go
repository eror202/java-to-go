package web

import (
	"github.com/gofiber/fiber/v2"
	"java-to-go/entity"
	"java-to-go/service"
)

func CreatePerson(services *service.PersonServ) fiber.Handler {

	return func(c *fiber.Ctx) error {
		type request struct {
			entity.Person
		}
		req := &request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		user := &entity.Person{
			FullName: req.FullName,
			Title:    req.Title,
			Age:      req.Age,
		}
		id, err := services.CreatePerson(user)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(201).JSON(id)
	}
}

func GetPersonById(services *service.PersonServ) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		person, err := services.GetPersonById(id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(200).JSON(person)
	}
}

func UpdatePersonById(services *service.PersonServ) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		type request struct {
			entity.Person
		}
		req := &request{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		user := &entity.Person{
			FullName: req.FullName,
			Title:    req.Title,
			Age:      req.Age,
		}
		person, err := services.UpdatePerson(user, id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(201).JSON(person)
	}
}

func DeletePersonById(services *service.PersonServ) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		_, err := services.DeletePersonById(id)
		if err != nil {
			return c.Status(422).JSON(err.Error())
		}
		return c.Status(200).JSON("deleted")
	}
}
