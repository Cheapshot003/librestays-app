//internal/handler/handler.go

package handler

import (
	"librestays/internal/db"

	"github.com/gofiber/fiber/v2"
)

func HomeGet(c *fiber.Ctx) error {
	return c.SendFile("./ui/html/index.html")
}

func LoginGet(c *fiber.Ctx) error {
	return c.SendFile("./ui/html/login.html")
}

func LoginPost(c *fiber.Ctx) error {

	if db.AuthenticateUser(c.FormValue("username"), c.FormValue("password")) {
		return c.SendFile("./ui/html/login_success.html")
	} else {
		return c.SendString("ERROR")
	}
}

func RegisterGet(c *fiber.Ctx) error {
	return c.SendFile("./ui/html/register.html")
}

func RegisterPost(c *fiber.Ctx) error {
	err := db.CreateUser(c.FormValue("username"), c.FormValue("password"))
	if err != nil {
		return c.SendString("Error creating User")
	} else {
		return c.SendString("User created!")
	}
}
