package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	instance *sql.DB
	once     sync.Once
)

// GetDB returns the singleton database instance.
func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		instance, err = sql.Open("sqlite3", "./events.db")
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}
		// Set database connection pool limits
		instance.SetMaxOpenConns(10)
		instance.SetMaxIdleConns(5)

		// Initialize the database schema
		createTables(instance)
	})
	return instance
}

// createTables creates the necessary tables if they do not already exist.
func createTables(db *sql.DB) {
	createEventsTable := `
        CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime DATETIME NOT NULL,
            user_id INTEGER
        );
    `
	_, err := db.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}
