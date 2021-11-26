package tikv

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tikv/client-go/v2/config"
	"github.com/tikv/client-go/v2/rawkv"
)

type Tikv struct {
	Client *rawkv.Client
}

// New return new client for TiKV
func New(ctx context.Context, cfg Config) (*rawkv.Client, error) {
	client, err := rawkv.NewClient(ctx, []string{cfg.Address}, config.DefaultConfig().Security)
	if err != nil {
		return nil, errors.Errorf("error while connecting to Tikv server: %v", err)
	}
	return client, nil
}

// Set a new key value pair in cache store.
func (t Tikv) Set(ctx context.Context, key string, value string) error {
	keyBytes := []byte(key)
	valueBytes := []byte(value)
	return t.Client.Put(context.Background(), keyBytes, valueBytes)
}

// Get value for key or empty if key doesn't exist.
func (t Tikv) Get(ctx context.Context, key string) (string, error) {
	keyBytes := []byte(key)
	bytes, err := t.Client.Get(context.Background(), keyBytes)
	value := string(bytes)
	return value, err
}

// Del delete key-value pair from store.
func (t Tikv) Del(ctx context.Context, key string) error {
	keyBytes := []byte(key)
	return t.Client.Delete(context.Background(), keyBytes)
}
