package forum

import (
	"github.com/gofiber/fiber/v2"
)

func Server() {
	var app = fiber.New()

	app.Static("/static", "./static")

    app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendFile("./templates/index.html")
	})
    app.Get("/home", func(c *fiber.Ctx) error {
        return c.Status(200).SendFile("./templates/home.html")
    })
    app.Get("/register", func(c *fiber.Ctx) error {
        return c.Status(200).SendFile("./templates/register.html")
    })
    app.Get("/login", func(c *fiber.Ctx) error {
        return c.Status(200).SendFile("./templates/login.html")
    })

    var auth = app.Group("/auth")

    auth.Post("/register", Register)

	app.Listen(":3000")
}
