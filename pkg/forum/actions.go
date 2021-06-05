package forum

import (
    "github.com/gofiber/fiber/v2"
    "fmt"
)

func Register(c *fiber.Ctx) error {
	var user = User{
		c.FormValue("name"),
		c.FormValue("email"),
		HashPassword(c.FormValue("password")),
	}
	AddUser(user)
	return c.SendString(fmt.Sprintf("User %s registered.", c.FormValue("name")))
}
