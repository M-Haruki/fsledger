package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/M-Haruki/fsledger/api/internal/server"
)

func main() {
	ctx := context.Background()

	cfg := server.Config{
		IsDev: os.Getenv("APP_ENV") == "dev",
	}
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		cfg.LogLevel = slog.LevelDebug
	case "info":
		cfg.LogLevel = slog.LevelInfo
	case "warn":
		cfg.LogLevel = slog.LevelWarn
	case "error":
		cfg.LogLevel = slog.LevelError
	default:
		cfg.LogLevel = slog.LevelInfo
	}

	server, err := server.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(":1323"); err != nil {
		log.Fatal(err)
	}
}
