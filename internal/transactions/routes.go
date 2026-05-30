package transactions

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "transactions", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/transactions", Description: "Record a transaction"},
		{Method: http.MethodGet, Path: "/v1/transactions", Description: "List merchant transactions"},
		{Method: http.MethodGet, Path: "/v1/transactions/{transaction_id}", Description: "Fetch transaction details"},
	})
}
