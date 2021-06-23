package repository

import (
	"database/sql"

	sqlite "github.com/mattn/go-sqlite3"
)

// create db and tables
func OpenDB(driver, path string) (*sql.DB, error) {
	db, err := sql.Open(driver, path)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(100)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = createAllTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

// create all tables in db
func createAllTables(database *sql.DB) error {
	tx, err := database.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	PRAGMA FOREIGN_KEY=on;
	`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS user(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT,
		email TEXT UNIQUE,
		role TEXT,
		created_date DATE
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS session(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER UNIQUE,
		token TEXT UNIQUE,
		exp_time DATE,
		FOREIGN KEY (user_id) REFERENCES user (id)
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS category_posts(
		category_name TEXT,
		post_id INTEGER,
		FOREIGN KEY (post_id) REFERENCES post (id),
		FOREIGN KEY (category_name) REFERENCES category (name)
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS category(
		name TEXT UNIQUE PRIMARY KEY
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS post(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT,
		content TEXT,
		likes INTEGER,
		dislikes INTEGER,
		created_date DATE,
		updated_date DATE,
		FOREIGN KEY (user_id) REFERENCES user (id)
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS comment(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		post_id INTEGER,
		content TEXT,
		likes INTEGER,
		dislikes INTEGER,
		created_date DATE,
		updated_date DATE,
		FOREIGN KEY (user_id) REFERENCES user (id),
		FOREIGN KEY (post_id) REFERENCES post (id)
	);`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS comment_votes(
		user_id INTEGER,
		comment_id INTEGER,
		type TEXT,
		FOREIGN KEY (user_id) REFERENCES user (id),
		FOREIGN KEY (comment_id) REFERENCES comment (id)
	);
	`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS post_votes(
		user_id INTEGER,
		post_id INTEGER,
		type TEXT,
		FOREIGN KEY (user_id) REFERENCES user (id),
		FOREIGN KEY (post_id) REFERENCES post (id)
	);
	`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS post_images(
		post_id INTEGER,
		path TEXT,
		FOREIGN KEY (post_id) REFERENCES post (id)
	);
	`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(`
	INSERT INTO category (name) 
	VALUES (?)`, "Art"); err != nil {
		if sqliteErr, ok := err.(sqlite.Error); ok {
			if sqliteErr.Code != sqlite.ErrConstraint {
				tx.Rollback()
				return err
			}
		}
	}

	if _, err := tx.Exec(`
	INSERT INTO category (name) 
	VALUES (?)`, "Landscape"); err != nil {
		if sqliteErr, ok := err.(sqlite.Error); ok {
			if sqliteErr.Code != sqlite.ErrConstraint {
				tx.Rollback()
				return err
			}
		}
	}

	if _, err := tx.Exec(`
	INSERT INTO category (name) 
	VALUES (?)`, "Animals"); err != nil {
		if sqliteErr, ok := err.(sqlite.Error); ok {
			if sqliteErr.Code != sqlite.ErrConstraint {
				tx.Rollback()
				return err
			}
		}
	}

	if _, err := tx.Exec(`
	INSERT INTO category (name) 
	VALUES (?)`, "Nature"); err != nil {
		if sqliteErr, ok := err.(sqlite.Error); ok {
			if sqliteErr.Code != sqlite.ErrConstraint {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit()
}

