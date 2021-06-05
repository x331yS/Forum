package forum

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"
)

func Server() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/home.html")
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendFile("./public/home.html")
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.SendFile("./public/register.html")
	})

	app.Post("/add_user", Register)

	app.Listen(":3000")
}
