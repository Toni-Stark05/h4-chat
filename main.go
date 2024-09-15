package main

import (
	"h4-chat/internal/config"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log.Info("Env uploaded", slog.String("env", cfg.Env))
}
