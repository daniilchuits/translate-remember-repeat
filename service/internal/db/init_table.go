package db

import (
	"context"
	"fmt"
)

func InitTranslationsTable(ctx context.Context) error {

	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
		CREATE TABLE IF NOT EXISTS translations(
			id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			sentence TEXT UNIQUE NOT NULL,
			deteched_language TEXT NOT NULL,
			translation TEXT NOT NULL
		);
	`

	if _, err = db.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("Error initializing table: %w", err)
	}
	return nil
}
