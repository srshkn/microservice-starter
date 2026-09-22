package logging

import (
	"fmt"
	"log/slog"
	"os"
)

type config interface {
	GetFormat() string
}

func New(cfg config) (*slog.Logger, error) {
	var handler slog.Handler

	switch cfg.GetFormat() {

	case "dev":
		handler = slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			},
		)

	case "prod":
		handler = slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		)
	default:
		return nil, fmt.Errorf("unsupported logger format: %q", cfg.GetFormat())
	}

	return slog.New(handler), nil
}
