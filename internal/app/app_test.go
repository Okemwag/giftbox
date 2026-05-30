package app

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Okemwag/giftbox/internal/platform/config"
)

func TestNewApplicationRegistersHealthRoutesThroughDependencies(t *testing.T) {
	application := NewApplication(Dependencies{
		Config: config.Config{Environment: "test"},
		Logger: slog.Default(),
	})

	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	response := httptest.NewRecorder()

	application.mux.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", response.Code)
	}
}
