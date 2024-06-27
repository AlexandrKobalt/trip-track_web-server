package config

import (
	"encoding/json"
	"os"

	grpcclient "github.com/AlexandrKobalt/trip-track_file-server/pkg/grpc/client"
	"github.com/AlexandrKobalt/trip-track_web-server/pkg/duration"
	"github.com/AlexandrKobalt/trip-track_web-server/pkg/fiberapp"
	"github.com/go-playground/validator"
)

const (
	path = "config/config.json"
)

type Config struct {
	StartTimeout duration.Seconds
	StopTimeout  duration.Seconds

	FiberApp       fiberapp.Config
	FileServerGRPC grpcclient.Config
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
