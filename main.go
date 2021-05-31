package main

import (
	// "database/sql"
	"fmt"
	// "github.com/gofiber/fiber/v2"
	// _ "github.com/mattn/go-sqlite3"
	// "log"
	"github.com/anatolethien/forum/pkg/forum_db"
	// "github.com/anatolethien/forum/pkg/forum_web"
)

func main() {
	// forum_db.CreateDatabase()
	// forum_db.AddUser(forum_db.User{Username: "totole", Password: "123"})
	// forum_db.ScanUser()
	// forum_web.Server()
	var Password = "1070df575a6b2f5a6f9baaa8d83ca4bc3174acd98fef141bfdd10895d50a283ea"
	fmt.Printf("%x\n", forum_db.EncryptPassword(Password))
}
