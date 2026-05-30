package audit

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform"
)

func RegisterRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "audit", []platform.Endpoint{
		{Method: http.MethodGet, Path: "/v1/audit/events", Description: "Search audit events"},
		{Method: http.MethodPost, Path: "/v1/audit/events", Description: "Append an audit event"},
	})
}
