package redis

import (
	"context"
)

type Redis struct {
}

func (r Redis) Set(ctx context.Context, key string, value string) error {
	panic("implement me")
}

func (r Redis) Get(ctx context.Context, key string) (string, error) {
	panic("implement me")
}

func (r Redis) Del(ctx context.Context, key string) error {
	panic("implement me")
}
