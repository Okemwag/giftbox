package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Okemwag/giftbox/pkg/observability"
)

func TestHealthEndpoint(t *testing.T) {
	mux := http.NewServeMux()
	observability.RegisterRoutes(mux)

	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	response := httptest.NewRecorder()

	mux.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
