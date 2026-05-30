package segments

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform"
)

func RegisterRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "segments", []platform.Endpoint{
		{Method: http.MethodPost, Path: "/v1/segments", Description: "Create a customer segment"},
		{Method: http.MethodGet, Path: "/v1/segments", Description: "List segments"},
		{Method: http.MethodPost, Path: "/v1/segments/{segment_id}/preview", Description: "Preview matching customers"},
	})
}
