package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"translate-remember-repeat/internal/domain"
)

func (manager *repoManager) SelectFromTranslationsPSQL(
	ctx context.Context,
	sentence string,
) (*domain.ResponseOfDeepL, error) {

	query := `
		SELECT
			deteched_language,
			translation
		FROM translations
		WHERE sentence = $1
	`

	var (
		detLang     string
		translation string
	)

	if err := manager.repo.QueryRowContext(
		ctx,
		query,
		sentence,
	).Scan(
		&detLang,
		&translation,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NoRowsFound
		}
		return nil, fmt.Errorf("%w: %w", domain.ReadingPSQL, err)
	}

	return domain.CreateResponse(
		translation, detLang,
	), nil
}
