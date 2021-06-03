package forum

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func CreateDatabase() {
	var database, _ = sql.Open("sqlite3", "./database.db")
	var statement, _ = database.Prepare("create table if not exists user (id integer primary key, name text, email text, password text)")
	statement.Exec()
}

func AddUser(user User) {
	var database, _ = sql.Open("sqlite3", "./database.db")
	var statement, _ = database.Prepare("insert into user (name, email, password) values (?, ?, ?)")
	statement.Exec(user.Name, user.Email, user.Password)
}

func ScanUser() {
	var database, _ = sql.Open("sqlite3", "./database.db")
	var rows, _ = database.Query("select name, email, password from user")
}
