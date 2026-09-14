package config

import (
	"fmt"
)

const (
	loggerFormatEnv string = "LOGGER_FORMAT"
)

type logger struct {
	Format string `conf:"required"`
}

func (l *logger) GetFormat() string {
	return l.Format
}

func (l *logger) validate() error {
	if l.Format != "dev" && l.Format != "prod" {
		return fmt.Errorf(
			"environment variable %q must be one of: dev, local, prod",
			loggerFormatEnv)
	}

	return nil
}
