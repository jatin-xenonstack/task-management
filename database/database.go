package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create tasks table if not exists
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		due_date TEXT,
		status TEXT DEFAULT 'pending'
	);`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}
