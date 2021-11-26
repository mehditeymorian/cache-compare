package tikv

import (
	"context"
)

type Tikv struct {
}

func (t Tikv) Set(ctx context.Context, key string, value string) error {
	panic("implement me")
}

func (t Tikv) Get(ctx context.Context, key string) (string, error) {
	panic("implement me")
}

func (t Tikv) Del(ctx context.Context, key string) error {
	panic("implement me")
}
