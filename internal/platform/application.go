package platform

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/Okemwag/giftbox/pkg/middleware"
	"github.com/Okemwag/giftbox/pkg/observability"
)

type RouteRegistrar func(*http.ServeMux)

type Application struct {
	config Config
	logger *slog.Logger
	mux    *http.ServeMux
}

func NewApplication(config Config, logger *slog.Logger) *Application {
	mux := http.NewServeMux()
	observability.RegisterRoutes(mux)

	return &Application{
		config: config,
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
	handler := middleware.Chain(
		a.mux,
		middleware.RequestID,
		middleware.Recoverer(a.logger),
		middleware.AccessLog(a.logger),
	)

	server := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	errs := make(chan error, 1)
	go func() {
		a.logger.Info("server listening", "addr", addr, "env", a.config.Environment)
		errs <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	case err := <-errs:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}
