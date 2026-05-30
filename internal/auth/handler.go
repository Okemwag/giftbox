package auth

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform/server"
)

func RegisterRoutes(mux *http.ServeMux) {
	server.RegisterModuleRoutes(mux, "auth", []server.Endpoint{
		{Method: http.MethodPost, Path: "/v1/auth/login", Description: "Authenticate merchant dashboard users"},
		{Method: http.MethodPost, Path: "/v1/auth/refresh", Description: "Refresh access tokens"},
		{Method: http.MethodPost, Path: "/v1/auth/logout", Description: "Revoke active sessions"},
	})
}
