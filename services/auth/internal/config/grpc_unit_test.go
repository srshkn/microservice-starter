package config

import (
	"testing"
	"time"
)

func TestGRPCValidate(t *testing.T) {
	tests := []struct {
		name    string
		server  grpc
		wantErr bool
	}{
		{
			name: "valid",
			server: grpc{
				Host:            "localhost",
				Port:            50051,
				ShutdownTimeout: 5 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "empty host",
			server: grpc{
				Port:            50051,
				ShutdownTimeout: 5 * time.Second,
			},
			wantErr: true,
		},
		{
			name: "port too low",
			server: grpc{
				Host:            "localhost",
				Port:            0,
				ShutdownTimeout: 5 * time.Second,
			},
			wantErr: true,
		},
		{
			name: "port too high",
			server: grpc{
				Host:            "localhost",
				Port:            65536,
				ShutdownTimeout: 5 * time.Second,
			},
			wantErr: true,
		},
		{
			name: "invalid shutdown timeout",
			server: grpc{
				Host:            "localhost",
				Port:            50051,
				ShutdownTimeout: 0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.server.validate()

			if (err != nil) != tt.wantErr {
				t.Fatalf("validateServer() error = %v", err)
			}
		})
	}
}
