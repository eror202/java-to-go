package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"java-to-go/web"
	"log"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Cookie",
		AllowMethods:     "GET,POST",
		AllowCredentials: true,
	}))

	api := app.Group("/api")
	web.UserRouter(api)
	log.Fatal(app.Listen(Config.GetString("port")))
}

var Config = viper.New()
