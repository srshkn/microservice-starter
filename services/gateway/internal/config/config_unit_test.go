package config

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		wantErr string
	}{
		{
			name: "success",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
			},
		},
		{
			name: "parse error",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "invalid",
				serverShutdownTimeoutEnv: "10s",
			},
			wantErr: "parse config",
		},
		{
			name: "validation error",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
			},
			wantErr: "validate config Server",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.env {
				t.Setenv(key, value)
			}

			cfg, _, err := New()

			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.wantErr)
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf(
						"expected error containing %q, got %q",
						tt.wantErr,
						err.Error(),
					)
				}

				return
			}

			if err != nil {
				t.Fatalf("New() unexpected error: %v", err)
			}

			if cfg == nil {
				t.Fatal("New() returned nil config")
			}
		})
	}
}
