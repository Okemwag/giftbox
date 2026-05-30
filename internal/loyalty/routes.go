package loyalty

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform"
)

func RegisterRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "loyalty", []platform.Endpoint{
		{Method: http.MethodPost, Path: "/v1/loyalty/programs", Description: "Create a loyalty program"},
		{Method: http.MethodPost, Path: "/v1/loyalty/points/accruals", Description: "Accrue points from eligible transactions"},
		{Method: http.MethodGet, Path: "/v1/loyalty/customers/{customer_id}/balance", Description: "Fetch customer points balance"},
	})
}
