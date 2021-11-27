package cmd

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mehditeymorian/cache-compare/internal/cache/redis"
	"github.com/mehditeymorian/cache-compare/internal/cache/tikv"
	"github.com/mehditeymorian/cache-compare/internal/config"
	"github.com/mehditeymorian/cache-compare/internal/http/handler"
	"log"
)

func Execute() {
	ctx := context.Background()

	cfg := config.New("config.yaml")
	log.Printf("config %+v\n", cfg)

	tikvClient, err := tikv.New(ctx, cfg.Tikv)
	if err != nil {
		log.Printf("error initializing tikv client: %v", err)
	}
	tikvCache := tikv.Tikv{Client: tikvClient}

	redisClient, err := redis.New(ctx, cfg.Redis)
	if err != nil {
		log.Printf("error initializing redis client: %v", err)
	}
	redisCache := redis.Redis{Client: redisClient}

	app := echo.New()

	handler.CacheHandler{Client: tikvCache}.Register("/tikv", app)
	handler.CacheHandler{Client: redisCache}.Register("/redis", app)

	if err := app.Start(":" + cfg.Http.Port); err != nil {
		log.Fatalf("failed to start HTTP server! %v", err)
	}

}
