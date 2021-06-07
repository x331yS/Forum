package forum

import (
	"database/sql"
	// "github.com/gofiber/storage/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func CreateDatabase() {
	var database, _ = sql.Open("sqlite3", "./database.db")
	var statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, name TEXT, email TEXT, password TEXT)")
	statement.Exec()
}

func AddUser(user User) {
	var database, _ = sql.Open("sqlite3", "./database.db")
	var statement, _ = database.Prepare("INSERT INTO user (name, email, password) VALUES (?, ?, ?)")
	statement.Exec(user.Name, user.Email, user.Password)
}

func NameIsValid(name string) bool {
    if name == "" {
        return false
    }
    var database, _ = sql.Open("sqlite3", "./database.db")
	var rows, _ = database.Query(fmt.Sprintf("SELECT name FROM user WHERE name = '%s'", name))
	fmt.Println(rows)
	// select name from user where name = 'loickordi'
	return true
}


// func FiberDb() {
// 	// Initialize default config
// 	store := sqlite3.New()
// 	store.Set
// }

// func ScanUser() {
// 	var database, _ = sql.Open("sqlite3", "./database.db")
// 	var rows, _ = database.Query("select name, email, password from user")
// }
