package config

import (
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
	}
}
