package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Okemwag/giftbox/internal/app"
	"github.com/Okemwag/giftbox/internal/mpesa"
	"github.com/Okemwag/giftbox/internal/platform/config"
	"github.com/Okemwag/giftbox/internal/whatsapp"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := config.LoadConfig()
	if cfg.WebhookAddr == "" {
		cfg.WebhookAddr = ":8081"
	}

	application := app.NewApplication(cfg, logger)
	application.Register(mpesa.RegisterWebhookRoutes, whatsapp.RegisterWebhookRoutes)

	if err := application.Serve(context.Background(), cfg.WebhookAddr); err != nil {
		logger.Error("webhook gateway stopped", "error", err)
		os.Exit(1)
	}
}
