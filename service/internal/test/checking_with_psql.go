package test

import (
	"context"
	"fmt"
)

type repo interface {
	GetTranslationPSQL(ctx context.Context, sentence string) (string, error)
}

type cli interface {
	WriteSentanceWaitTranslation(ctx context.Context, sentence string) (string, error)
	WrongAnswer(ctx context.Context, must string) error
	CorrectAnswer(ctx context.Context) error
}

type postgresChecker struct {
	repo
	cli
}

func NewPostgresChecker(
	repo repo,
	cli cli,
) *postgresChecker {
	return &postgresChecker{
		repo: repo,
		cli:  cli,
	}
}

func (p *postgresChecker) CheckWithPostgres(ctx context.Context, sentence string) (int, error) {

	usersTranslation, err := p.WriteSentanceWaitTranslation(ctx, sentence)
	if err != nil {
		return 0, fmt.Errorf("Error getting user's translation: %w", err)
	}

	translation, err := p.GetTranslationPSQL(ctx, sentence)
	if err != nil {
		return 0, fmt.Errorf("Error checking postgres translation: %w", err)
	}

	if usersTranslation == translation {
		if err := p.CorrectAnswer(ctx); err != nil {
			return 0, fmt.Errorf("Error logging to cli: %w", err)
		}
		return 1, nil
	}
	if err = p.WrongAnswer(ctx, translation); err != nil {
		return 0, fmt.Errorf("Error logging to cli: %w", err)
	}
	return 0, nil
}
