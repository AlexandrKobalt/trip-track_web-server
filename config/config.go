package config

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator"
)

const (
	path = "config/config.json"
)

type Config struct {
}

func LoadConfig() (cfg *Config, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(jsonFile).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	err = validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
