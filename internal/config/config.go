package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/mehditeymorian/cache-compare/internal/cache/tikv"
	"github.com/mehditeymorian/cache-compare/internal/http"
	"log"
)

type Config struct {
	Http http.Config `koanf:"http"`
	Tikv tikv.Config `koanf:"tikv"`
}

func New(filename string) Config {
	k := koanf.New(".")

	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default config: %v", err)
	}

	if err := k.Load(file.Provider(filename), yaml.Parser()); err != nil {
		log.Printf("error loading file config: %v", err)
	}

	var cfg Config
	err := k.Unmarshal("", &cfg)
	if err != nil {
		log.Fatalf("erro unmarshaling: %v", err)
	}

	return cfg
}
