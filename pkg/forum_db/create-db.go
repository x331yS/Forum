package forum_db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

// create - insert
// read - select
// update - update
// delete - delete

type User struct {
  Username string
  Password string
}

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

// func ex1() {
// 	var database, _ = sql.Open("sqlite3", "./ex1.db")
//
// 	var statement, _ = database.Prepare("create table if not exists user (id integer primary key, firstname text, lastname text)")
// 	statement.Exec()
//
// 	statement, _ = database.Prepare("insert into user (firstname, lastname) values (?, ?)")
// 	statement.Exec("Anatole", "Thien")
//
// 	var rows, _ = database.Query("select id, firstname, lastname from user")
// 	var id int
// 	var firstname string
// 	var lastname string
// 	for rows.Next() {
// 		rows.Scan(&id, &firstname, &lastname)
// 		fmt.Printf("%d : %s %s\n", id, firstname, lastname)
// 	}
// }
