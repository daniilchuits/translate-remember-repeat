package repo

import "context"

func (manager *repoManager) GetTranslationPSQL(ctx context.Context, sentence string) (string, error) {

	query := `
		SELECT translation
		FROM translations
		WHERE sentence = $1
	`

	var tanslation string
	if err := manager.repo.QueryRowContext(
		ctx,
		query,
		sentence,
	).Scan(
		&tanslation,
	); err != nil {
		return "", err
	}
	return tanslation, nil
}
