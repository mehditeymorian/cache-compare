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
			Address: "http://ADDRESS:2379",
		},
		Redis: redis.Config{
			Address:  "ADDRESS:6379",
			Password: "",
			DB:       0,
		},
		MongoDB: mongodb.Config{
			Uri:      "mongodb://ADDRESS:27017",
			Name:     "cache-compare",
			Username: "USERNAME",
			Password: "PASSWORD",
		},
	}
}
