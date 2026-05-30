package rewards

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "rewards", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/rewards", Description: "Create a reward"},
		{Method: http.MethodPost, Path: "/v1/rewards/redemptions", Description: "Redeem a reward"},
		{Method: http.MethodGet, Path: "/v1/rewards/redemptions", Description: "List reward redemptions"},
	})
}
