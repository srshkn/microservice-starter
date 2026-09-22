package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/conf/v3"
	"github.com/joho/godotenv"
)

const (
	configPrefixEnv = "SERVICE"
)

type config struct {
	service string
	Server  *server
	Logger  *logger
	Swagger *swagger
	CORS    *cors
}

func New() (*config, string, error) {
	var cfg config

	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, "", fmt.Errorf("load .env: %w", err)
	}

	prefix := os.Getenv(configPrefixEnv)

	cfg.service = prefix

	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		return nil, help, fmt.Errorf("parse config: %w", err)
	}

	if err := cfg.Server.validate(); err != nil {
		return nil, "", fmt.Errorf("validate config Server: %w", err)
	}

	if err := cfg.Logger.validate(); err != nil {
		return nil, "", fmt.Errorf("validate config Logger: %w", err)
	}

	if err := cfg.CORS.validate(); err != nil {
		return nil, "", fmt.Errorf("validate config CORS: %w", err)
	}

	return &cfg, "", nil
}

func (c *config) GetServiceName() string {
	return c.service
}
