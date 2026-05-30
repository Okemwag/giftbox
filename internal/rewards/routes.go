package rewards

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform"
)

func RegisterRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "rewards", []platform.Endpoint{
		{Method: http.MethodPost, Path: "/v1/rewards", Description: "Create a reward"},
		{Method: http.MethodPost, Path: "/v1/rewards/redemptions", Description: "Redeem a reward"},
		{Method: http.MethodGet, Path: "/v1/rewards/redemptions", Description: "List reward redemptions"},
	})
}
