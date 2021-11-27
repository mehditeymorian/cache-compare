package config

import (
	"github.com/mehditeymorian/cache-compare/internal/cache/mongodb"
	"github.com/mehditeymorian/cache-compare/internal/cache/redis"
	"github.com/mehditeymorian/cache-compare/internal/cache/tikv"
	"github.com/mehditeymorian/cache-compare/internal/http"
)

func Default() Config {
	return Config{
		Http: http.Config{
			Port: "8080",
		},
		Tikv: tikv.Config{
			Address: "http://tikv-pd:2379",
		},
		Redis: redis.Config{
			Address:  "redis:6379",
			Password: "",
			DB:       0,
		},
		MongoDB: mongodb.Config{
			Uri:  "mongodb://localhost:27017",
			Name: "cache-compare",
		},
	}
}
