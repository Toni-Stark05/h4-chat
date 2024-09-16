package logger

import (
	"github.com/samber/do"
	"log/slog"
	"os"
)

func NewLogger(_ *do.Injector) (*slog.Logger, error) {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	logger.Info("Logger initialized")
	return logger, nil
}
