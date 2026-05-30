package app

import (
	"log/slog"

	"github.com/Okemwag/giftbox/internal/platform/config"
	"github.com/Okemwag/giftbox/internal/platform/database"
)

type Dependencies struct {
	Config config.Config
	DB     *database.Postgres
	Logger *slog.Logger
}
