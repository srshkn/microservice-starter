package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/conf/v3"
	"github.com/joho/godotenv"
)

const (
	configPrefixEnv = "CONFIG_PREFIX"
)

type config struct {
	Server *server
}

func New() (*config, string, error) {
	var cfg config

	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, "", fmt.Errorf("load .env: %w", err)
	}

	prefix := os.Getenv(configPrefixEnv)

	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		return nil, help, fmt.Errorf("parse config: %w", err)
	}

	if err := cfg.Server.validate(); err != nil {
		return nil, "", fmt.Errorf("validate config Server: %w", err)
	}

	return &cfg, "", nil
}
