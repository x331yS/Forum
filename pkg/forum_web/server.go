package forum_web

import (
  // "github.com/anatolethien/forum/pkg/forum_db"
  "github.com/gofiber/fiber/v2"
  "fmt"
)

func Server()  {
		app := fiber.New()

		app.Static("/", "./public")

		app.Listen(":3000")
}

func response() string {
  return fmt.Sprintf("Hello %s!", "you")
}
