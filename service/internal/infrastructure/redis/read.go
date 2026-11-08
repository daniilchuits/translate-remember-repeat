package redis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"translate-remember-repeat/internal/domain"
)

func (manager *redisManager) Read(
	ctx context.Context,
	sentence string,
) (*domain.ResponseOfDeepL, error) {

	val, err := manager.client.Get(ctx, sentence).Bytes()
	if err != nil {
		return nil, fmt.Errorf("Error getting from redis: %w", err)
	}

	var resp domain.ResponseOfDeepL

	if err = json.NewDecoder(bytes.NewBuffer(val)).Decode(&resp); err != nil {
		return nil, fmt.Errorf("%w from redis: %w", domain.Decoding, err)
	}
	return &resp, nil
}
