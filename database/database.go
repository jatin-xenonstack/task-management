package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	} else {
		fmt.Println("Table Created successfully")
	}
}
