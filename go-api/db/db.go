package db

import (
	"database/sql"
	"fmt"
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

func query(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
		panic(fmt.Sprintf("Error creating table: %v", err))
	}
}

func createUsersTable(db *sql.DB) {
	createUsersTable := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT NOT NULL UNIQUE,
			password TEXT NULL
        );
    `
	query(db, createUsersTable)
}

func createEventsTable(db *sql.DB) {

	createEventsTable := `
        CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime DATETIME NOT NULL,
            user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
        );
    `
	query(db, createEventsTable)
}

func createTables(db *sql.DB) {
	createUsersTable(db)
	createEventsTable(db)
}
