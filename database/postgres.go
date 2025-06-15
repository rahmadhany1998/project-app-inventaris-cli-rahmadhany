package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDB, open connection to PostgreSQL database
func ConnectDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=inventory sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	// Test Connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database failed: %w", err)
	}
	return db, nil
}
