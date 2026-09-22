package grpc

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	authgrpc "auth/internal/grpc/auth"

	authv1 "starter.local/gen/auth/v1"
)

type config interface {
	GetHost() string
	GetPort() int
	GetShutdownTimeout() time.Duration
}

type RegisterFunc func(*grpc.Server)

type Server struct {
	server          *grpc.Server
	address         string
	shutdownTimeout time.Duration
	logger          *slog.Logger
}

func New(
	cfg config,
	logger *slog.Logger,
) *Server {
	address := net.JoinHostPort(
		cfg.GetHost(),
		strconv.Itoa(cfg.GetPort()),
	)

	authHandler := authgrpc.New()

	server := grpc.NewServer()

	authv1.RegisterAuthServiceServer(
		server,
		authHandler,
	)

	return &Server{
		address:         address,
		logger:          logger,
		server:          server,
		shutdownTimeout: cfg.GetShutdownTimeout(),
	}
}

func (s *Server) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("listen grpc: %w", err)
	}

	errCh := make(chan error, 1)

	go func() {
		s.logger.Info(
			"grpc server started",
			slog.String("address", s.address),
		)

		errCh <- s.server.Serve(listener)
	}()

	select {

	case err := <-errCh:
		if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			return fmt.Errorf("serve grpc: %w", err)
		}

		return nil

	case <-ctx.Done():
		s.shutdown()

		return nil
	}
}

func (s *Server) shutdown() {
	done := make(chan struct{})

	go func() {
		s.server.GracefulStop()
		close(done)
	}()

	timer := time.NewTimer(s.shutdownTimeout)
	defer timer.Stop()

	select {

	case <-done:
		s.logger.Info("grpc server stopped")

	case <-timer.C:
		s.logger.Warn("grpc graceful shutdown timeout")

		s.server.Stop()
	}
}
