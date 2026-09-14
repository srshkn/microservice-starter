package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"gateway/internal/app"
	"gateway/internal/config"
	"gateway/internal/logging"
)

func main() {

	// -------------------------------------------------------------------------
	// Context

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// -------------------------------------------------------------------------
	// Configuration

	cfg, _, err := config.New()
	if err != nil {
		slog.Error(
			"failed to load configuration",
			slog.Any("error", err),
		)
		return
	}

	// -------------------------------------------------------------------------
	// Logger

	logger, err := logging.New(cfg.Logger)
	if err != nil {
		slog.Error(
			"failed to configuration logger",
			slog.Any("error", err),
		)
		return
	}

	// -------------------------------------------------------------------------
	// Server

	serverApp := app.New(
		cfg.Server,
		logger,
		cfg.Swagger,
	)

	// -------------------------------------------------------------------------
	// Run

	err = serverApp.Run(ctx)
	if err != nil {
		logger.Error(
			"server stopped unexpectedly",
			slog.Any("error", err),
		)
	}
}
