package forum

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var Data = RegisterData{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: HashPassword(c.FormValue("password")),
	}
	if CheckUserName(Data) == false {
		return c.SendString("error: name already used")
	}
	if CheckUserEmail(Data) == false {
		return c.SendString("error: email already used")
	}
	AddUser(Data)
	return c.SendString(fmt.Sprintf("user %s registered", Data.Name))
}

func Login(c *fiber.Ctx) error {

}
