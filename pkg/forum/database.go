package forum

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	// "fmt"
)

func CreateDatabase() {
	var Database, _ = sql.Open("sqlite3", "./database.db")
	var statement, _ = Database.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, name TEXT, email TEXT, password TEXT)")
	statement.Exec()
}

func AddUser(Data RegisterData) {
	var Database, _ = sql.Open("sqlite3", "./database.db")
	var statement, _ = Database.Prepare("INSERT INTO user (name, email, password) VALUES (?, ?, ?)")
	statement.Exec(Data.Name, Data.Email, Data.Password)
}

func CheckUserName(Data RegisterData) bool {
	var Database, _ = sql.Open("sqlite3", "./database.db")
	var row = Database.QueryRow("SELECT name FROM user WHERE name=?", Data.Name)
	if err := row.Scan(); err != nil {
		if err == sql.ErrNoRows {
			return true
		}
	}
	return false
}

func CheckUserEmail(Data RegisterData) bool {
	var Database, _ = sql.Open("sqlite3", "./database.db")
	var row = Database.QueryRow("SELECT email FROM user WHERE email=?", Data.Email)
	if err := row.Scan(); err != nil {
		if err == sql.ErrNoRows {
			return true
		}
	}
	return false
}
