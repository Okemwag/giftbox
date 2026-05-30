package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/Okemwag/giftbox/internal/platform/config"
	"github.com/Okemwag/giftbox/internal/platform/observability"
	"github.com/Okemwag/giftbox/internal/platform/server"
)

type RouteRegistrar func(*http.ServeMux)

type Application struct {
	config config.Config
	logger *slog.Logger
	mux    *http.ServeMux
}

func NewApplication(cfg config.Config, logger *slog.Logger) *Application {
	mux := http.NewServeMux()
	observability.RegisterRoutes(mux)

	return &Application{
		config: cfg,
		logger: logger,
		mux:    mux,
	}
}

func (a *Application) Register(registrars ...RouteRegistrar) {
	for _, register := range registrars {
		register(a.mux)
	}
}

func (a *Application) Serve(ctx context.Context, addr string) error {
	handler := server.Chain(
		a.mux,
		server.RequestID,
		server.Recoverer(a.logger),
		server.AccessLog(a.logger),
	)

	srv := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	errs := make(chan error, 1)
	go func() {
		a.logger.Info("server listening", "addr", addr, "env", a.config.Environment)
		errs <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return srv.Shutdown(shutdownCtx)
	case err := <-errs:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}
