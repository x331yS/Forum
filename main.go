package main

import (
	// "database/sql"
	// "fmt"
	// "github.com/gofiber/fiber/v2"
	// _ "github.com/mattn/go-sqlite3"
	// "log"
	"github.com/anatolethien/forum/pkg/forum_db"
	"github.com/anatolethien/forum/pkg/forum_web"
)

func main() {
	forum_db.CreateDatabase()
	// forum_db.AddUser(forum_db.User{Username: "totole", Password: "123"})
	// forum_db.ScanUser()
	forum_web.Server()
}
