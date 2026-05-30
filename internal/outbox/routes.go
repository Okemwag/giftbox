package outbox

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "outbox", []server.Endpoint{
		{Method: http.MethodGet, Path: "/v1/outbox/messages", Description: "List outbox messages"},
		{Method: http.MethodPost, Path: "/v1/outbox/messages/{message_id}/retry", Description: "Retry outbox delivery"},
	})
}
