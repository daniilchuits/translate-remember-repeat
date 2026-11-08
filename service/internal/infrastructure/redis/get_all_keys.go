package redis

import "context"

func (manager *redisManager) GetAllKeys(ctx context.Context) ([]string, error) {

	results, err := manager.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	return results, nil
}
