package logging

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestNewDevLogger(t *testing.T) {
	logger, output := newLoggerWithCapturedOutput(t, "gateway", "dev", func(logger *slog.Logger) {
		logger.Debug("request handled", slog.String("request_id", "req-1"))
	})

	if !logger.Enabled(context.Background(), slog.LevelDebug) {
		t.Fatal("expected debug level to be enabled")
	}

	for _, want := range []string{
		"level=DEBUG",
		`msg="request handled"`,
		"source=",
		"service=gateway",
		"request_id=req-1",
	} {
		if !strings.Contains(output, want) {
			t.Errorf("expected log output to contain %q, got %q", want, output)
		}
	}
}

func TestNewProdLogger(t *testing.T) {
	logger, output := newLoggerWithCapturedOutput(t, "gateway", "prod", func(logger *slog.Logger) {
		logger.Debug("debug message")
		logger.Info("request handled", slog.String("request_id", "req-1"))
	})

	if logger.Enabled(context.Background(), slog.LevelDebug) {
		t.Fatal("expected debug level to be disabled")
	}

	var entry map[string]any
	if err := json.Unmarshal([]byte(output), &entry); err != nil {
		t.Fatalf("expected valid JSON log output, got %q: %v", output, err)
	}

	for key, want := range map[string]string{
		"level":      "INFO",
		"msg":        "request handled",
		"service":    "gateway",
		"request_id": "req-1",
	} {
		if got := entry[key]; got != want {
			t.Errorf("expected %q field to be %q, got %v", key, want, got)
		}
	}

	if _, ok := entry["source"]; ok {
		t.Error("expected production log output not to contain source")
	}
}

func TestNewReturnsErrorForUnsupportedFormat(t *testing.T) {
	logger, err := New("gateway", configStub{format: "invalid"})

	if err == nil {
		t.Fatal("expected an error")
	}

	const want = `unsupported logger format: "invalid"`
	if err.Error() != want {
		t.Fatalf("expected error %q, got %q", want, err)
	}

	if logger != nil {
		t.Fatal("expected nil logger")
	}
}

func newLoggerWithCapturedOutput(
	t *testing.T,
	service string,
	format string,
	writeLog func(*slog.Logger),
) (*slog.Logger, string) {
	t.Helper()

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("create pipe: %v", err)
	}
	defer reader.Close()

	stdout := os.Stdout
	os.Stdout = writer
	logger, err := New(service, configStub{format: format})
	os.Stdout = stdout
	if err != nil {
		writer.Close()
		t.Fatalf("create logger: %v", err)
	}

	writeLog(logger)
	if err := writer.Close(); err != nil {
		t.Fatalf("close log writer: %v", err)
	}

	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("read log output: %v", err)
	}

	return logger, string(output)
}

type configStub struct {
	format string
}

func (c configStub) GetFormat() string {
	return c.format
}
