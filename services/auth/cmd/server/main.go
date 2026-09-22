package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"auth/internal/config"

	grpcserver "auth/internal/grpc"

	"microservice-starter/pkg/logging"
)

func main() {

	// -------------------------------------------------------------------------
	// Context

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
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

	logger, err := logging.New(cfg.GetServiceName(), cfg.Logger)
	if err != nil {
		slog.Error(
			"failed to configuration logger",
			slog.Any("error", err),
		)
		return
	}

	// -------------------------------------------------------------------------
	// Server

	server := grpcserver.New(
		cfg.GRPC,
		logger,
	)

	// -------------------------------------------------------------------------
	// Run

	if err := server.Run(ctx); err != nil {
		logger.Error(
			"grpc server failed",
			slog.String("error", err.Error()),
		)

		os.Exit(1)
	}
}
