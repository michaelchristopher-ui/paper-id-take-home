package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func regenerateSQL(dbFileName string) {
	// Remove the existing database file if it exists
	if _, err := os.Stat(dbFileName); err == nil {
		err = os.Remove(dbFileName)
		if err != nil {
			log.Fatalf("Error removing existing database: %v", err)
		}
	}

	// Create a new database
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Create a table (example)
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER
		);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	fmt.Println("Database created and table initialized successfully.")
}
