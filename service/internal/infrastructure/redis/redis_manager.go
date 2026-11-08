package redis

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisManager struct {
	client *redis.Client
}

func New() (*redisManager, error) {

	cnnStr := os.Getenv("REDIS_CNN_STR")
	if cnnStr == "" {
		return nil, errors.New("Empty 'REDIS_CNN_STR' variable in '.env'")
	}

	client := redis.NewClient(&redis.Options{
		Addr: cnnStr,
	})

	connCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(connCtx).Err(); err != nil {
		_ = client.Close()
		return nil, fmt.Errorf("Error pinging redis: %s", err)
	}
	return &redisManager{client: client}, nil
}

func (r *redisManager) Close() error {
	return r.client.Close()
}
