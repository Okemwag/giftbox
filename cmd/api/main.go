package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Okemwag/giftbox/internal/analytics"
	"github.com/Okemwag/giftbox/internal/audit"
	"github.com/Okemwag/giftbox/internal/auth"
	"github.com/Okemwag/giftbox/internal/campaigns"
	"github.com/Okemwag/giftbox/internal/consent"
	"github.com/Okemwag/giftbox/internal/customers"
	"github.com/Okemwag/giftbox/internal/loyalty"
	"github.com/Okemwag/giftbox/internal/mpesa"
	"github.com/Okemwag/giftbox/internal/notifications"
	"github.com/Okemwag/giftbox/internal/outbox"
	"github.com/Okemwag/giftbox/internal/platform"
	"github.com/Okemwag/giftbox/internal/rewards"
	"github.com/Okemwag/giftbox/internal/segments"
	"github.com/Okemwag/giftbox/internal/tenants"
	"github.com/Okemwag/giftbox/internal/transactions"
	"github.com/Okemwag/giftbox/internal/whatsapp"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := platform.LoadConfig()

	app := platform.NewApplication(cfg, logger)
	app.Register(
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

	if err := app.Serve(context.Background(), cfg.HTTPAddr); err != nil {
		logger.Error("api stopped", "error", err)
		os.Exit(1)
	}
}
