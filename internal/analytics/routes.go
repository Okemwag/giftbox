package analytics

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "analytics", []server.Endpoint{
		{Method: http.MethodGet, Path: "/v1/analytics/overview", Description: "Fetch merchant performance summary"},
		{Method: http.MethodGet, Path: "/v1/analytics/loyalty", Description: "Fetch loyalty engagement metrics"},
		{Method: http.MethodGet, Path: "/v1/analytics/campaigns", Description: "Fetch campaign analytics"},
	})
}
