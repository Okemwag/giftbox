package notifications

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "notifications", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/notifications", Description: "Queue a notification"},
		{Method: http.MethodGet, Path: "/v1/notifications", Description: "List notification jobs"},
		{Method: http.MethodPost, Path: "/v1/notifications/{notification_id}/retry", Description: "Retry a failed notification"},
	})
}
