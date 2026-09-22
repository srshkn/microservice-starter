package logging

import (
	"log/slog"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		format      string
		wantHandler any
		wantErr     string
	}{
		{
			name:        "dev",
			format:      "dev",
			wantHandler: &slog.TextHandler{},
		},
		{
			name:        "prod",
			format:      "prod",
			wantHandler: &slog.JSONHandler{},
		},
		{
			name:    "invalid format",
			format:  "invalid",
			wantErr: `unsupported logger format: "invalid"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := configStub{
				format: tt.format,
			}

			logger, err := New(cfg)

			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q", tt.wantErr)
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("expected error containing %q, got %q", tt.wantErr, err)
				}

				if logger != nil {
					t.Fatal("expected nil logger")
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if logger == nil {
				t.Fatal("expected logger, got nil")
			}

			switch tt.wantHandler.(type) {
			case *slog.TextHandler:
				if _, ok := logger.Handler().(*slog.TextHandler); !ok {
					t.Fatalf(
						"expected *slog.TextHandler, got %T",
						logger.Handler(),
					)
				}

			case *slog.JSONHandler:
				if _, ok := logger.Handler().(*slog.JSONHandler); !ok {
					t.Fatalf(
						"expected *slog.JSONHandler, got %T",
						logger.Handler(),
					)
				}
			}
		})
	}
}

type configStub struct {
	format string
}

func (c configStub) GetFormat() string {
	return c.format
}
