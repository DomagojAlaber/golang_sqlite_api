package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		name        VARCHAR(50) NOT NULL,
		description TEXT NOT NULL,
		location    TEXT NOT NULL,
		dateTime    DATETIME NOT NULL,
		user_id     INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic(fmt.Sprintf("Error, could not create events table: %v", err))
	}
}
