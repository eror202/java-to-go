package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"java-to-go/web"
	"log"
)

func main() {
	Config.SetConfigFile("config.yaml")
	Config.ReadInConfig()

	db, err := ConnectDB(
		Config.GetString("db.host"),
		Config.GetString("db.port"),
		Config.GetString("db.dbname"),
		Config.GetString("db.username"),
		Config.GetString("db.password"),
		Config.GetString("db.sslmode"),
	)
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	log.Println("Database connection success!")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Cookie",
		AllowMethods:     "GET,POST,OPTIONS",
		AllowCredentials: true,
	}))

	api := app.Group("/api")
	web.UserRouter(api, db)
	log.Fatal(app.Listen(Config.GetString("port")))
}

var Config = viper.New()

func ConnectDB(host, port, dbname, user, password, sslmode string) (*sqlx.DB, error) {
	databaseURL := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host,
		port,
		dbname,
		user,
		password,
		sslmode,
	)
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
