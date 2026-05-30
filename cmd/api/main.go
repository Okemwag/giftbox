package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/Okemwag/giftbox/internal/analytics"
	"github.com/Okemwag/giftbox/internal/app"
	"github.com/Okemwag/giftbox/internal/audit"
	"github.com/Okemwag/giftbox/internal/auth"
	"github.com/Okemwag/giftbox/internal/campaigns"
	"github.com/Okemwag/giftbox/internal/consent"
	"github.com/Okemwag/giftbox/internal/customers"
	"github.com/Okemwag/giftbox/internal/loyalty"
	"github.com/Okemwag/giftbox/internal/mpesa"
	"github.com/Okemwag/giftbox/internal/notifications"
	"github.com/Okemwag/giftbox/internal/outbox"
	"github.com/Okemwag/giftbox/internal/platform/config"
	"github.com/Okemwag/giftbox/internal/platform/database"
	"github.com/Okemwag/giftbox/internal/rewards"
	"github.com/Okemwag/giftbox/internal/segments"
	"github.com/Okemwag/giftbox/internal/tenants"
	"github.com/Okemwag/giftbox/internal/transactions"
	"github.com/Okemwag/giftbox/internal/whatsapp"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg, err := config.Load()
	if err != nil {
		logger.Error("invalid configuration", "error", err)
		os.Exit(1)
	}

	startupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.Open(startupCtx, database.Config{
		DatabaseURL:     cfg.DatabaseURL,
		MaxConns:        10,
		MinConns:        1,
		MaxConnLifetime: time.Hour,
	})
	if err != nil {
		logger.Error("database startup failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	application := app.NewApplication(app.Dependencies{
		Config: cfg,
		DB:     db,
		Logger: logger,
	})
	application.Register(
		auth.RegisterRoutes,
		tenants.RegisterRoutes,
		customers.RegisterRoutes,
		consent.RegisterRoutes,
		transactions.RegisterRoutes,
		mpesa.RegisterRoutes,
		whatsapp.RegisterRoutes,
		loyalty.RegisterRoutes,
		rewards.RegisterRoutes,
		campaigns.RegisterRoutes,
		segments.RegisterRoutes,
		analytics.RegisterRoutes,
		notifications.RegisterRoutes,
		audit.RegisterRoutes,
		outbox.RegisterRoutes,
	)

	if err := application.Serve(context.Background(), cfg.HTTPAddr()); err != nil {
		logger.Error("api stopped", "error", err)
		os.Exit(1)
	}
}
