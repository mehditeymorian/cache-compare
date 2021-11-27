package cmd

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mehditeymorian/cache-compare/internal/cache/mongodb"
	"github.com/mehditeymorian/cache-compare/internal/cache/redis"
	"github.com/mehditeymorian/cache-compare/internal/cache/tikv"
	"github.com/mehditeymorian/cache-compare/internal/config"
	"github.com/mehditeymorian/cache-compare/internal/http/handler"
	"log"
)

func Execute() {
	ctx := context.Background()

	// reading config
	cfg := config.New("config.yaml")
	log.Printf("config %+v\n", cfg)

	// initializing tikv client
	tikvClient, err := tikv.New(ctx, cfg.Tikv)
	if err != nil {
		log.Printf("error initializing tikv client: %v", err)
	}
	tikvCache := tikv.Tikv{Client: tikvClient}

	// initializing redis client
	redisClient, err := redis.New(ctx, cfg.Redis)
	if err != nil {
		log.Printf("error initializing redis client: %v", err)
	}
	redisCache := redis.Redis{Client: redisClient}

	// initializing mongoDB client
	mongoClient, err := mongodb.New(ctx, cfg.MongoDB)
	if err != nil {
		log.Printf("error initializing MongoDB client: %v", err)
	}
	mongoCache := mongodb.MongoDB{Client: mongoClient}

	// creating http app
	app := echo.New()
	// registering handlers
	handler.CacheHandler{Client: tikvCache}.Register("/tikv", app)
	handler.CacheHandler{Client: redisCache}.Register("/redis", app)
	handler.CacheHandler{Client: mongoCache}.Register("/mongo", app)

	// starting server
	if err := app.Start(":" + cfg.Http.Port); err != nil {
		log.Fatalf("failed to start HTTP server! %v", err)
	}

}
