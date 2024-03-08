package main

import (
	"librestays/internal/db"
	handler "librestays/internal/handlers"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading enviroment variables")
	}

	time.Sleep(5 * time.Second) // wait for database to start

	db.ConnectDatabase()
	db.InitDB()
	app := fiber.New()
	app.Static("/static", "./ui/static")
	app.Get("/", handler.HomeGet)
	app.Get("/login", handler.LoginGet)
	app.Post("/login", handler.LoginPost)
	app.Get("/register", handler.RegisterGet)
	app.Post("/register", handler.RegisterPost)

	err = app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
