package usecases

import (
	"context"
	"fmt"
	"translate-remember-repeat/internal/domain"
)

type cliTestWriter interface {
	TestAnnouncement(ctx context.Context) error
	ZeroSentences(ctx context.Context) error
	WrongLen(ctx context.Context, must, have int) error
	CorrectAnswers(ctx context.Context, must, have int) error
	SeparateTest(ctx context.Context) error
	Separate(ctx context.Context) error
}

type sentLister interface {
	GetSentences(ctx context.Context) ([]domain.Sentence, error)
}

type postgresChecker interface {
	CheckWithPostgres(ctx context.Context, sentence string) (int, error)
}

type redisChecker interface {
	CheckWithRedis(ctx context.Context, sentence string) (int, error)
}

type testUsecase struct {
	cliTestWriter
	sentLister
	postgresChecker
	redisChecker
}

func NewTestUsecase(
	cliTestWriter cliTestWriter,
	sentLister sentLister,
	postgresChecker postgresChecker,
	redisChecker redisChecker,
) *testUsecase {
	return &testUsecase{
		cliTestWriter:   cliTestWriter,
		sentLister:      sentLister,
		postgresChecker: postgresChecker,
		redisChecker:    redisChecker,
	}
}

func (t *testUsecase) Test(ctx context.Context) error {

	if err := t.SeparateTest(ctx); err != nil {
		return err
	}

	if err := t.TestAnnouncement(ctx); err != nil {
		return err
	}

	sentences, err := t.GetSentences(ctx) // panic: runtime error: slice bounds out of range [:-3]
	if err != nil {
		return fmt.Errorf("Test execution error: %w", err)
	}

	if len(sentences) == 0 {
		if err = t.ZeroSentences(ctx); err != nil {
			return err
		}
		return nil
	} else if domain.TestSentences > len(sentences) {
		if err = t.WrongLen(ctx, domain.TestSentences, len(sentences)); err != nil {
			return err
		}
	}
	// output test if the len is less or if it's ok
	var count int

	for _, sentence := range sentences {
		switch sentence.Type {
		case domain.Postgres:
			n, err := t.CheckWithPostgres(ctx, sentence.Sent)
			if err != nil {
				return err
			}
			count += n
		case domain.Redis:
			n, err := t.CheckWithRedis(ctx, sentence.Sent)
			if err != nil {
				return err
			}
			count += n
		default:
			return domain.WrongType
		}
	}

	if err = t.CorrectAnswers(ctx, len(sentences), count); err != nil {
		return err
	}

	if err := t.Separate(ctx); err != nil {
		return err
	}

	return nil
}
