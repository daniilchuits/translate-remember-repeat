package repo

import "context"

func (manager *repoManager) Select(ctx context.Context) ([]string, error) {

	query := `
		SELECT sentence FROM translations
	`

	var sentences []string

	rows, err := manager.repo.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sentence string
		if err := rows.Scan(&sentence); err != nil {
			return nil, err
		}
		sentences = append(sentences, sentence)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return sentences, nil
}
