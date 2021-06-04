package forum

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"
	// "fmt"
)

func Static() {
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendFile("./skeleton/home.html")
    })

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.SendFile("./skeleton/register.html")
	})

	app.Get("/profile", func(c *fiber.Ctx) error {
		return c.SendFile("./skeleton/profile.html")
	})

	app.Get("/notifications", func(c *fiber.Ctx) error {
		return c.SendFile("./skeleton/notifications.html")
	})

	app.Post("/add_user", RegisterUser)

	app.Static("/", "./skeleton")

	app.Listen(":3000")
}

func RegisterUser(c *fiber.Ctx) error {
    var test = User{
        c.FormValue("name"),
        c.FormValue("email"),
        HashPassword(c.FormValue("password")),
    }
    AddUser(test)
	return c.SendString("Register")
}
