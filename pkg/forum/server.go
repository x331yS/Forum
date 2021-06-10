package forum

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"
)

func Server() {
	var app = fiber.New()

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendFile("./public/index.html")
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.Status(200).SendFile("./public/home.html")
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Status(200).SendFile("./public/register.html")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Status(200).SendFile("./public/login.html")
	})

	var auth = app.Group("/auth")

	auth.Post("/register", Register)
	auth.Post("/login", Login)

	app.Listen(":3000")
}
