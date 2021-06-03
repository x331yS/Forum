package forum

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() {
	var database, _ = sql.Open("sqlite3", "./database.db")
	var statement, _ = database.Prepare("create table if not exists user (id integer primary key, username text, password text)")
	statement.Exec()
}

func AddUser(user User) {
	var database, _ = sql.Open("sqlite3", "./database.db")
	var statement, _ = database.Prepare("insert into user (username, password) values (?, ?)")
	statement.Exec(user.Username, user.Password)
}

func ScanUser() {
	var database, _ = sql.Open("sqlite3", "./database.db")
	var rows, _ = database.Query("select username, password from user")
	fmt.Println(rows)
}
