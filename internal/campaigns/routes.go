package campaigns

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform"
)

func RegisterRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "campaigns", []platform.Endpoint{
		{Method: http.MethodPost, Path: "/v1/campaigns", Description: "Create a campaign"},
		{Method: http.MethodPost, Path: "/v1/campaigns/{campaign_id}/send", Description: "Dispatch campaign messages"},
		{Method: http.MethodGet, Path: "/v1/campaigns/{campaign_id}/performance", Description: "Fetch campaign performance"},
	})
}
