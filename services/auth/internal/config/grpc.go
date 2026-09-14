package config

import (
	"fmt"
	"time"
)

const (
	grpcHostEnv            string = "GRPC_HOST"
	grpcPortEnv            string = "GRPC_PORT"
	grpcShutdownTimeoutEnv string = "GRPC_SHUTDOWN_TIMEOUT"
)

type grpc struct {
	Host            string        `conf:"required"`
	Port            int           `conf:"required"`
	ShutdownTimeout time.Duration `conf:"required"`
}

func (s *grpc) GetHost() string {
	return s.Host
}

func (s *grpc) GetPort() int {
	return s.Port
}

func (s *grpc) GetShutdownTimeout() time.Duration {
	return s.ShutdownTimeout
}

func (s *grpc) validate() error {
	switch {

	case s.Host == "":
		return fmt.Errorf(
			"environment variable %q is required",
			grpcHostEnv,
		)

	case s.Port < 1 || s.Port > 65535:
		return fmt.Errorf(
			"environment variable %q must be between 1 and 65535",
			grpcPortEnv,
		)

	case s.ShutdownTimeout <= 0:
		return fmt.Errorf(
			"environment variable %q must be greater than 0",
			grpcShutdownTimeoutEnv,
		)
	}
	return nil
}
