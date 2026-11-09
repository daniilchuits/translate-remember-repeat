package test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"translate-remember-repeat/internal/domain"
)

type redisTester interface {
	GetAllKeys(ctx context.Context) ([]string, error)
}

type postgresTester interface {
	Select(ctx context.Context) ([]string, error)
}

type getKeysUsecase struct {
	redisTester
	postgresTester
}

func NewGetKeysUsecase(
	redisTester redisTester,
	postgresTester postgresTester,
) *getKeysUsecase {
	return &getKeysUsecase{
		redisTester:    redisTester,
		postgresTester: postgresTester,
	}
}

func (g *getKeysUsecase) GetSentences(ctx context.Context) ([]domain.Sentence, error) {

	var sentences []domain.Sentence

	redisKeys, err := g.GetAllKeys(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error redis gitting keys: %w", err)
	}
	switch {
	case len(redisKeys) > domain.TestSentences:
		newIdxs := (rand.Perm(len(redisKeys)))[:domain.TestSentences]
		for _, idx := range newIdxs {
			sentences = append(sentences, domain.Sentence{
				Sent: redisKeys[idx],
				Type: domain.Redis,
			})
		}
		return sentences, nil
	case len(redisKeys) == domain.TestSentences:

		for i := range redisKeys {
			sentences = append(sentences, domain.Sentence{
				Sent: redisKeys[i],
				Type: domain.Redis,
			})
		}
		return sentences, nil
	default:
	}

	for _, sentence := range redisKeys {
		sentences = append(sentences, domain.Sentence{
			Sent: sentence,
			Type: domain.Redis,
		})
	}

	postgresKeys, err := g.Select(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error postgres getting keys: %w", err)
	}

	keysLeft := domain.TestSentences - len(sentences)
	if len(postgresKeys) > keysLeft {
		newIdxs := (rand.Perm(len(postgresKeys)))[:keysLeft]
		for _, idx := range newIdxs {
			sentences = append(sentences, domain.Sentence{
				Sent: postgresKeys[idx],
				Type: domain.Postgres,
			})
		}
		return sentences, nil
	}
	for i := range postgresKeys {
		sentences = append(sentences, domain.Sentence{
			Sent: postgresKeys[i],
			Type: domain.Postgres,
		})
	}
	return sentences, nil
}
