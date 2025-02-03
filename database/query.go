package database

// Create tasks table if not exists

var query string = `
CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT,
	due_date TEXT,
	status TEXT 
);`

// fmt.Println(query)
