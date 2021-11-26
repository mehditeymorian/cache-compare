package cmd

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mehditeymorian/cache-compare/internal/cache/tikv"
	"github.com/mehditeymorian/cache-compare/internal/config"
	"github.com/mehditeymorian/cache-compare/internal/http/handler"
	"log"
)

func Execute() {
	ctx := context.Background()

	cfg := config.New("config.yaml")

	tikvClient, err := tikv.New(ctx, cfg.Tikv)
	if err != nil {
		log.Printf("error initializing tikv client: %v", err)
	}
	tikvCache := tikv.Tikv{Client: tikvClient}

	app := echo.New()

	handler.Tikv{Client: tikvCache}.Register(app)

	if err := app.Start(":" + cfg.Http.Port); err != nil {
		log.Fatalf("failed to start HTTP server! %v", err)
	}

}
