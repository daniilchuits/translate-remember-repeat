package test

import (
	"context"
	"fmt"
	"math/rand/v2"
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

const (
	testSentences = 3
)

func (g *getKeysUsecase) GetSentences(ctx context.Context) ([]string, error) {

	var sentences []string

	redisKeys, err := g.GetAllKeys(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error redis gitting keys: %w", err)
	}
	switch {
	case len(redisKeys) > testSentences:
		newIdxs := (rand.Perm(len(redisKeys)))[:testSentences]
		for _, idx := range newIdxs {
			sentences = append(sentences, redisKeys[idx])
		}
		return sentences, nil
	case len(redisKeys) == testSentences:
		return redisKeys, nil
	default:
	}

	for _, sentence := range redisKeys {
		sentences = append(sentences, sentence)
	}

	postgresKeys, err := g.Select(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error postgres getting keys: %w", err)
	}

	keysLeft := len(sentences) - testSentences
	if len(postgresKeys) > keysLeft {
		newIdxs := (rand.Perm(len(postgresKeys)))[:keysLeft]
		for _, idx := range newIdxs {
			sentences = append(sentences, postgresKeys[idx])
		}
		return sentences, nil
	}
	for i := range postgresKeys {
		sentences = append(sentences, postgresKeys[i])
	}
	return sentences, nil
}
