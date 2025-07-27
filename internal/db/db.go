// https://medium.com/@peymaan.abedinpour/golang-crud-app-tutorial-step-by-step-guide-using-sqlite-a3ce08a4fc81
package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type Todo struct {
	ID    int
	Title string
}

var DB *sql.DB

const dbPath = "./data/app.db"

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err) // Log an error and stop the program if the database can't be opened
	}

	sqlStmt := `
 CREATE TABLE IF NOT EXISTS entries (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  source TEXT,
  user TEXT,
  password TEXT
 );
 CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  main_password TEXT
 );`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt)
	}
}
