package main

import (
	"librestays/internal/db"
	handler "librestays/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db.ConnectDatabase()
	app := fiber.New()
	app.Static("/static", "./ui/static")
	app.Get("/", handler.HomeGet)
	app.Get("/login", handler.LoginGet)
	app.Post("/login", handler.LoginPost)
	app.Get("/register", handler.RegisterGet)
	app.Post("/register", handler.RegisterPost)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
