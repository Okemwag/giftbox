package whatsapp

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform"
)

func RegisterRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "whatsapp", []platform.Endpoint{
		{Method: http.MethodPost, Path: "/v1/whatsapp/messages", Description: "Queue WhatsApp messages"},
		{Method: http.MethodGet, Path: "/v1/whatsapp/templates", Description: "List approved templates"},
	})
}

func RegisterWebhookRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "webhooks/whatsapp", []platform.Endpoint{
		{Method: http.MethodPost, Path: "/v1/webhooks/whatsapp/messages", Description: "Receive WhatsApp message status callbacks"},
		{Method: http.MethodGet, Path: "/v1/webhooks/whatsapp/verify", Description: "Verify WhatsApp webhook ownership"},
	})
}
