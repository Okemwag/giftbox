package whatsapp

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "whatsapp", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/whatsapp/messages", Description: "Queue WhatsApp messages"},
		{Method: http.MethodGet, Path: "/v1/whatsapp/templates", Description: "List approved templates"},
	})
}

func RegisterWebhookRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "webhooks/whatsapp", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/webhooks/whatsapp/messages", Description: "Receive WhatsApp message status callbacks"},
		{Method: http.MethodGet, Path: "/v1/webhooks/whatsapp/verify", Description: "Verify WhatsApp webhook ownership"},
	})
}
