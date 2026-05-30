package mpesa

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "mpesa", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/mpesa/reconcile", Description: "Reconcile M-Pesa transaction references"},
		{Method: http.MethodGet, Path: "/v1/mpesa/events", Description: "List M-Pesa ingestion events"},
	})
}

func RegisterWebhookRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "webhooks/mpesa", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/webhooks/mpesa/c2b", Description: "Receive M-Pesa C2B callbacks"},
		{Method: http.MethodPost, Path: "/v1/webhooks/mpesa/stk", Description: "Receive M-Pesa STK callbacks"},
	})
}
