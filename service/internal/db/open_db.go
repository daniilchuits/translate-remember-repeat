package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func openDB() (*sql.DB, error) {

	cnnStr := os.Getenv("POSTGRES_CNN_STR")
	if cnnStr == "" {
		return nil, fmt.Errorf("Empty 'POSTGRES_CNN_STR'")
	}

	db, err := sql.Open("postgres", cnnStr)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to db: %w", err)
	}

	return connectDB(db)
}
