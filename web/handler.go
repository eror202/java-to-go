package web

import (
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {

	app.Post("/createUser", CreatePerson())
	app.Get("/getUser/:id", GetPersonById())
	app.Post("/updateUser/:id", UpdatePersonById())
	app.Get("deleteUser/:id", DeletePersonById())

	app.Post("/createBook", CreateBook())
	app.Get("/getBook/:id", GetBookById())
	app.Post("/updateBook/:id", UpdateBook())
	app.Get("deleteBook/:id", DeleteBookById())
}
