package campaigns

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "campaigns", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/campaigns", Description: "Create a campaign"},
		{Method: http.MethodPost, Path: "/v1/campaigns/{campaign_id}/send", Description: "Dispatch campaign messages"},
		{Method: http.MethodGet, Path: "/v1/campaigns/{campaign_id}/performance", Description: "Fetch campaign performance"},
	})
}
