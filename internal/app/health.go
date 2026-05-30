package app

import (
	"context"
	"net/http"
	"time"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func registerHealthRoutes(mux *http.ServeMux, deps Dependencies) {
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		server.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if deps.DB == nil {
			server.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := deps.DB.Ping(ctx); err != nil {
			deps.Logger.Error("readiness check failed", "error", err)
			server.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "unavailable"})
			return
		}

		server.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
}
