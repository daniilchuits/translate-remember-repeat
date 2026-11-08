package repo

import (
	"context"
	"fmt"
	"translate-remember-repeat/internal/domain"
)

func (manager *repoManager) InsertTranslationsPSQL(
	ctx context.Context,
	sentence string,
	resp *domain.ResponseOfDeepL,
) error {

	query := `
		INSERT INTO translations (sentence,deteched_language,translation) VALUES
		($1,$2,$3);
	`

	if _, err := manager.repo.ExecContext(
		ctx,
		query,
		sentence,
		resp.DetechedSourceLanguage(),
		resp.Text(),
	); err != nil {
		return fmt.Errorf("%w: %w", domain.InsertingPSQL, err)
	}
	return nil
}
