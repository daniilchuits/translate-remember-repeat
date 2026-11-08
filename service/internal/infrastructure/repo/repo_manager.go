package repo

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type repoManager struct {
	repo *sql.DB
}

func OpenDB() (*repoManager, error) {

	cnnStr := os.Getenv("POSTGRES_CNN_STR")
	if cnnStr == "" {
		return nil, fmt.Errorf("Empty 'POSTGRES_CNN_STR' variable in '.env'")
	}
	db, err := sql.Open("postgres", cnnStr)
	if err != nil {
		return nil, fmt.Errorf("Error sql.Open: %w", err)
	}

	if err = pingWithRetries(db); err != nil {
		return nil, fmt.Errorf("Error pinging db: %w", err)
	}

	return new(db), nil
}

func pingWithRetries(db *sql.DB) error {
	var err error
	for range 10 {

		err = db.Ping()
		if err == nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return err
}

func new(db *sql.DB) *repoManager {
	return &repoManager{repo: db}
}

func (repo *repoManager) Close() error {
	return repo.repo.Close()
}
