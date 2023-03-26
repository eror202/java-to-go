package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"java-to-go/repository"
	"java-to-go/service"
)

type Handler struct {
	services *service.Service
}

func UserRouter(app fiber.Router, db *sqlx.DB) {
	repo := repository.NewRepository(db)
	services := service.NewService(repo)

	app.Post("/createUser", CreatePerson(services))
	app.Get("/getUser/:id", GetPersonById(services))
	app.Post("/updateUser/:id", UpdatePersonById(services))
	app.Get("deleteUser/:id", DeletePersonById(services))
}
