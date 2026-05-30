package customers

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "customers", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/customers", Description: "Register a customer"},
		{Method: http.MethodGet, Path: "/v1/customers", Description: "Search merchant customers"},
		{Method: http.MethodPatch, Path: "/v1/customers/{customer_id}", Description: "Update a customer profile"},
	})
}
