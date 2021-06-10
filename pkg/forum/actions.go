package forum

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user = User{
		c.FormValue("name"),
		c.FormValue("email"),
		HashPassword(c.FormValue("password")),
	}
	// Check if user exists
	// Check if email used
	AddUser(user)
	return c.SendString(fmt.Sprintf("User %s registered.", c.FormValue("name")))
}

// Work in progress
func Login(c *fiber.Ctx) error {
	var user = User{
		c.FormValue("name"),
		c.FormValue("email"),
		c.FormValue("password"),
	}

	var token = jwt.New(jwt.SigningMethodHS256)

    var claims = token.Claims.(jwt.MapClaims)
    claims[""]

	return c.SendString(fmt.Sprintf("name: %s\nemail: %s\npassword: %s\n", user.Name, user.Email, user.Password))
}
