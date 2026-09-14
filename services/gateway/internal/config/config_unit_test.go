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
		// success

		{
			name: "success dev logger",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
		},
		{
			name: "success prod logger",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "prod",
				swaggerEnabledEnv:        "true",
			},
		},
		{
			name: "success ip host",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "127.0.0.1",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
		},
		{
			name: "success port minimum",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "1",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
		},
		{
			name: "success port maximum",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "65535",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
		},
		{
			name: "success shutdown timeout milliseconds",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "500ms",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
		},
		{
			name: "success shutdown timeout minutes",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "1m",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
		},
		{
			name: "success swagger enabled",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
		},
		{
			name: "success swagger disabled",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "false",
			},
		},

		// parse errors

		{
			name: "invalid port string",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "invalid",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "parse config",
		},
		{
			name: "invalid port float",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080.5",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "parse config",
		},
		{
			name: "invalid shutdown timeout",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "invalid",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "parse config",
		},
		{
			name: "shutdown timeout without unit",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "parse config",
		},
		{
			name: "invalid swagger enabled",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "invalid",
			},
			wantErr: "parse config",
		},
		{
			name: "empty swagger enabled",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "",
			},
			wantErr: "parse config",
		},

		// validation errors

		{
			name: "empty host",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "validate config Server",
		},
		{
			name: "port zero",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "0",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "validate config Server",
		},
		{
			name: "port negative",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "-1",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "validate config Server",
		},
		{
			name: "port above maximum",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "65536",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "validate config Server",
		},
		{
			name: "zero shutdown timeout",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "0s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "validate config Server",
		},
		{
			name: "negative shutdown timeout",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "-1s",
				loggerFormatEnv:          "dev",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "validate config Server",
		},
		{
			name: "empty logger format",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "validate config Logger",
		},
		{
			name: "invalid logger format",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "json",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "validate config Logger",
		},
		{
			name: "invalid logger format uppercase",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "localhost",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "DEV",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "validate config Logger",
		},

		// multiple invalid fields

		{
			name: "empty host and invalid logger",
			env: map[string]string{
				configPrefixEnv:          "",
				serverHostEnv:            "",
				serverPortEnv:            "8080",
				serverShutdownTimeoutEnv: "10s",
				loggerFormatEnv:          "invalid",
				swaggerEnabledEnv:        "true",
			},
			wantErr: "validate config",
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
