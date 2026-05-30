package notifications

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform"
)

func RegisterRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "notifications", []platform.Endpoint{
		{Method: http.MethodPost, Path: "/v1/notifications", Description: "Queue a notification"},
		{Method: http.MethodGet, Path: "/v1/notifications", Description: "List notification jobs"},
		{Method: http.MethodPost, Path: "/v1/notifications/{notification_id}/retry", Description: "Retry a failed notification"},
	})
}
