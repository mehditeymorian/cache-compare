package mongodb

import (
	"context"
	"fmt"
	"github.com/mehditeymorian/cache-compare/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const collectionName = "cache"

type MongoDB struct {
	Client *mongo.Database
}

func New(ctx context.Context, cfg Config) (*mongo.Database, error) {
	// create mongodb connection
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Uri))
	if err != nil {
		return nil, fmt.Errorf("db new client error: %w", err)
	}

	// connect to the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), 10*time.Second)
		defer done()

		if err := client.Connect(ctx); err != nil {
			return nil, fmt.Errorf("db connection error: %v", err)
		}
	}
	// ping the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), 10*time.Second)
		defer done()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return nil, fmt.Errorf("db ping error: %w", err)
		}
	}

	return client.Database(cfg.Name), nil
}

func (m MongoDB) Set(ctx context.Context, key string, value string) error {
	data := model.NewKVPair(key, value)
	_, err := m.Client.Collection(collectionName).InsertOne(ctx, data)
	return err
}

func (m MongoDB) Get(ctx context.Context, key string) (string, error) {

	result := m.Client.Collection(collectionName).FindOne(ctx, bson.D{{"key", key}})

	var kvpair model.KVPair

	err := result.Decode(&kvpair)
	if err != nil {
		return "", err
	}

	return kvpair.Value, nil
}

func (m MongoDB) Del(ctx context.Context, key string) error {
	_, err := m.Client.Collection(collectionName).DeleteOne(ctx, bson.M{"key": key})
	return err
}
