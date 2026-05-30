package consent

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform"
)

func RegisterRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "consent", []platform.Endpoint{
		{Method: http.MethodPost, Path: "/v1/consent", Description: "Capture communication consent"},
		{Method: http.MethodDelete, Path: "/v1/consent/{customer_id}", Description: "Withdraw customer consent"},
		{Method: http.MethodGet, Path: "/v1/consent/{customer_id}", Description: "Fetch consent status"},
	})
}
