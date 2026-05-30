package analytics

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform"
)

func RegisterRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "analytics", []platform.Endpoint{
		{Method: http.MethodGet, Path: "/v1/analytics/overview", Description: "Fetch merchant performance summary"},
		{Method: http.MethodGet, Path: "/v1/analytics/loyalty", Description: "Fetch loyalty engagement metrics"},
		{Method: http.MethodGet, Path: "/v1/analytics/campaigns", Description: "Fetch campaign analytics"},
	})
}
