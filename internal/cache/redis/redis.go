package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/common/log"
	"time"
)

type Redis struct {
	Client *redis.Client
}

func New(ctx context.Context, cfg Config) (*redis.Client, error) {
	timeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	log.Info("Connecting to redis...")
	if err := client.Ping(timeout).Err(); err != nil {
		return nil, fmt.Errorf("error connecting to redis: %v", err)
	}

	return client, nil
}

func (r Redis) Set(ctx context.Context, key string, value string) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

func (r Redis) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r Redis) Del(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}
