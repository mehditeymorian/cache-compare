package mongodb

import (
	"context"
)

type MongoDB struct {
}

func (m MongoDB) Set(ctx context.Context, key string, value string) error {
	panic("implement me")
}

func (m MongoDB) Get(ctx context.Context, key string) (string, error) {
	panic("implement me")
}

func (m MongoDB) Del(ctx context.Context, key string) error {
	panic("implement me")
}
