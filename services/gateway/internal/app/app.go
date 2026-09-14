package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"

	"gateway/internal/middleware"
	"gateway/internal/swagger"

	v1GenAPI "gateway/internal/generated/v1"
	handlerApp "gateway/internal/http/v1/handler"
)

type configServer interface {
	GetHost() string
	GetPort() int
	GetShutdownTimeout() time.Duration
}

type serverApp struct {
	server          *http.Server
	logger          *slog.Logger
	shutdownTimeout time.Duration
}

func New(
	serverCfg configServer,
	logger *slog.Logger,
	swaggerCfg swagger.Config,

) *serverApp {
	mux := http.NewServeMux()

	handler := handlerApp.New(
		handlerApp.NewMeta(),
	)

	strictHandler := v1GenAPI.NewStrictHandlerWithOptions(
		handler,
		[]v1GenAPI.StrictMiddlewareFunc{
			middleware.Logging(logger),
		},
		v1GenAPI.StrictHTTPServerOptions{},
	)

	apiHandler := v1GenAPI.HandlerWithOptions(
		strictHandler,
		v1GenAPI.StdHTTPServerOptions{
			BaseURL: "/api/v1",
		},
	)

	swagger.Register(mux, swaggerCfg)

	mux.Handle("/", apiHandler)

	address := net.JoinHostPort(
		serverCfg.GetHost(),
		strconv.Itoa(serverCfg.GetPort()),
	)

	return &serverApp{
		server: &http.Server{
			Addr:              address,
			Handler:           mux,
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       15 * time.Second,
			WriteTimeout:      15 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
		logger:          logger,
		shutdownTimeout: serverCfg.GetShutdownTimeout(),
	}

}

func (s *serverApp) gracefulStop(serverErr <-chan error) error {
	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		s.shutdownTimeout,
	)
	defer cancel()

	s.logger.Info(
		"shutting down HTTP server",
		slog.Duration("timeout", s.shutdownTimeout),
	)

	if err := s.server.Shutdown(shutdownCtx); err != nil {
		s.logger.Error(
			"graceful shutdown failed",
			slog.Any("error", err),
		)

		if closeErr := s.server.Close(); closeErr != nil {
			s.logger.Error(
				"force close HTTP server failed",
				slog.Any("error", closeErr),
			)
		}

		return fmt.Errorf("shutdown HTTP server: %w", err)
	}

	err := <-serverErr
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("HTTP server stopped: %w", err)
	}

	s.logger.Info("HTTP server stopped")

	return nil
}

func (s *serverApp) Run(ctx context.Context) error {
	serverErr := make(chan error, 1)

	go func() {
		serverErr <- s.server.ListenAndServe()
	}()

	s.logger.Info("starting server",
		slog.String("address", s.server.Addr),
	)

	s.logger.Info("server endpoints",
		slog.String("api", "http://"+s.server.Addr),
		slog.String("swagger", "http://"+s.server.Addr+"/docs/"),
	)

	select {

	case err := <-serverErr:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err

	case <-ctx.Done():
		s.logger.Info("shutdown signal received")
		return s.gracefulStop(serverErr)

	}
}

func (s *serverApp) Handler() http.Handler {
	return s.server.Handler
}
