package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"translate-remember-repeat/internal/domain"
)

func (manager *redisManager) Save(
	ctx context.Context,
	sentence string,
	resp *domain.ResponseOfDeepL,
) error {

	body, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("%w: %w", domain.Marshaling, err)
	}

	if err = manager.client.Set(ctx, sentence, body, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("%w: %w", domain.SettingRedis, err)
	}
	return nil
}
