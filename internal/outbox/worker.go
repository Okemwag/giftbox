package outbox

import (
	"context"
	"log/slog"
	"time"
)

type Worker struct {
	logger       *slog.Logger
	pollInterval time.Duration
}

func NewWorker(logger *slog.Logger, pollInterval time.Duration) *Worker {
	if pollInterval <= 0 {
		pollInterval = 5 * time.Second
	}
	return &Worker{logger: logger, pollInterval: pollInterval}
}

func (w *Worker) Run(ctx context.Context) error {
	ticker := time.NewTicker(w.pollInterval)
	defer ticker.Stop()

	w.logger.Info("outbox worker started", "poll_interval", w.pollInterval.String())
	for {
		select {
		case <-ctx.Done():
			w.logger.Info("outbox worker stopped")
			return nil
		case <-ticker.C:
			w.logger.Info("outbox poll completed", "processed", 0)
		}
	}
}
