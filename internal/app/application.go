package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/Okemwag/giftbox/internal/platform/config"
	"github.com/Okemwag/giftbox/internal/platform/server"
)

type RouteRegistrar func(*http.ServeMux)

type Application struct {
	config config.Config
	deps   Dependencies
	logger *slog.Logger
	mux    *http.ServeMux
}

func NewApplication(deps Dependencies) *Application {
	mux := http.NewServeMux()
	registerHealthRoutes(mux, deps)

	return &Application{
		config: deps.Config,
		deps:   deps,
		logger: deps.Logger,
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
		server.CORS,
		server.JSONContentType,
		server.Timeout(30*time.Second),
		server.Recoverer(a.logger),
		server.AccessLog(a.logger),
	)

	srv := server.NewHTTPServer(addr, handler)

	errs := make(chan error, 1)
	go func() {
		a.logger.Info("server listening", "addr", addr, "env", a.config.Environment)
		errs <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		return server.Shutdown(context.Background(), srv)
	case err := <-errs:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}
