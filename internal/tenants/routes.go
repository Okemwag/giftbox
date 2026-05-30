package tenants

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "tenants", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/tenants", Description: "Create a merchant tenant"},
		{Method: http.MethodGet, Path: "/v1/tenants/{tenant_id}", Description: "Fetch tenant profile and settings"},
		{Method: http.MethodPatch, Path: "/v1/tenants/{tenant_id}", Description: "Update tenant settings"},
	})
}
