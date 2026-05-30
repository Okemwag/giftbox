package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Okemwag/giftbox/internal/mpesa"
	"github.com/Okemwag/giftbox/internal/platform"
	"github.com/Okemwag/giftbox/internal/whatsapp"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := platform.LoadConfig()
	if cfg.WebhookAddr == "" {
		cfg.WebhookAddr = ":8081"
	}

	app := platform.NewApplication(cfg, logger)
	app.Register(mpesa.RegisterWebhookRoutes, whatsapp.RegisterWebhookRoutes)

	if err := app.Serve(context.Background(), cfg.WebhookAddr); err != nil {
		logger.Error("webhook gateway stopped", "error", err)
		os.Exit(1)
	}
}
