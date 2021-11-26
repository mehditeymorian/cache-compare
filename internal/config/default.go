package config

import "github.com/mehditeymorian/cache-compare/internal/cache/tikv"

func Default() Config {
	return Config{
		Tikv: tikv.Config{
			Address: "http://tikv-pd:2379",
		},
	}
}
