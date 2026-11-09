package redis

import (
	"bytes"
	"context"
	"encoding/json"
	"translate-remember-repeat/internal/domain"
)

func (manager *redisManager) GetTranslationRedis(
	ctx context.Context,
	sentence string,
) (string, error) {

	b, err := manager.client.Get(ctx, sentence).Bytes()
	if err != nil {
		return "", err
	}

	var resp domain.ResponseOfDeepL
	if err := json.NewDecoder(bytes.NewBuffer(b)).Decode(
		&resp,
	); err != nil {
		return "", err
	}
	return resp.Text(), nil
}
