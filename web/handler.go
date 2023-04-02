package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"java-to-go/repository"
	"java-to-go/service"
)

func UserRouter(app fiber.Router, db *sqlx.DB) {
	personRepo := repository.NewPersRep(db)
	personServices := service.NewPersonService(personRepo)

	app.Post("/createUser", CreatePerson(personServices))
	app.Get("/getUser/:id", GetPersonById(personServices))
	app.Post("/updateUser/:id", UpdatePersonById(personServices))
	app.Get("deleteUser/:id", DeletePersonById(personServices))

	bookRepo := repository.NewBookRep(db)
	bookServices := service.NewBookService(bookRepo)

	app.Post("/createBook", CreateBook(bookServices))
	app.Get("/getUser/:id", GetBookById(bookServices))
	app.Post("/updateUser/:id", UpdateBook(bookServices))
	app.Get("deleteUser/:id", DeleteBookById(bookServices))
}
