package db

import (
	"database/sql"
	"fmt"
	"time"
)

func connectDB(db *sql.DB) (*sql.DB, error) {

	var err error
	for range 10 {
		err = db.Ping()
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}
	return nil, fmt.Errorf("Error pinging db: %w", err)
}
