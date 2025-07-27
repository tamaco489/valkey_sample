package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/tamaco489/valkey_sample/internal/config"
	"github.com/tamaco489/valkey_sample/internal/router"
)

func main() {
	cfg := config.Load()
	mux := router.Setup()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	slog.Info("Starting server", "port", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, mux); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
