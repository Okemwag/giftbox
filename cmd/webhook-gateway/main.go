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
	cfg, err := config.Load()
	if err != nil {
		logger.Error("invalid configuration", "error", err)
		os.Exit(1)
	}

	application := app.NewApplication(app.Dependencies{
		Config: cfg,
		Logger: logger,
	})
	application.Register(mpesa.RegisterWebhookRoutes, whatsapp.RegisterWebhookRoutes)

	if err := application.Serve(context.Background(), cfg.HTTPAddr()); err != nil {
		logger.Error("webhook gateway stopped", "error", err)
		os.Exit(1)
	}
}
