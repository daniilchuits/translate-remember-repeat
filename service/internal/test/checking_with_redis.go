package test

import (
	"context"
	"fmt"
)

type redis interface {
	GetTranslationRedis(ctx context.Context, sentence string) (string, error)
}

type redisChecker struct {
	redis
	cli
}

func NewRedisChecker(
	redis redis,
	cli cli,
) *redisChecker {
	return &redisChecker{
		redis: redis,
		cli:   cli,
	}
}

func (r *redisChecker) CheckWithRedis(ctx context.Context, sentence string) (int, error) {

	usersTranslation, err := r.WriteSentanceWaitTranslation(ctx, sentence)
	if err != nil {
		return 0, fmt.Errorf("Error getting user's translation: %w", err)
	}

	translation, err := r.GetTranslationRedis(ctx, sentence)
	if err != nil {
		return 0, fmt.Errorf("Error getting redis translation: %w", err)
	}

	if usersTranslation == translation {
		if err = r.CorrectAnswer(ctx); err != nil {
			return 0, fmt.Errorf("Error logging to cli: %w", err)
		}
		return 1, nil
	}
	if err = r.WrongAnswer(ctx, translation); err != nil {
		return 0, fmt.Errorf("Error logging to cli: %w", err)
	}
	return 0, nil
}
