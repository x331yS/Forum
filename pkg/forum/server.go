package forum

import (
    "github.com/gofiber/fiber/v2"
)

func Static() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Listen(":3000")
}
