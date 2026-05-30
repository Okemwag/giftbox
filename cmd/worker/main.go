package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Okemwag/giftbox/internal/outbox"
	"github.com/Okemwag/giftbox/internal/platform"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := platform.LoadConfig()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	worker := outbox.NewWorker(logger, cfg.WorkerPollInterval)
	if err := worker.Run(ctx); err != nil {
		logger.Error("worker stopped", "error", err)
		os.Exit(1)
	}
}
