package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *redis.Client
}

func New(ctx context.Context, cfg Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

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
