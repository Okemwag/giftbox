package auth

import (
	"net/http"

	"github.com/Okemwag/giftbox/internal/platform"
)

func RegisterRoutes(mux *http.ServeMux) {
	platform.RegisterModuleRoutes(mux, "auth", []platform.Endpoint{
		{Method: http.MethodPost, Path: "/v1/auth/login", Description: "Authenticate merchant dashboard users"},
		{Method: http.MethodPost, Path: "/v1/auth/refresh", Description: "Refresh access tokens"},
		{Method: http.MethodPost, Path: "/v1/auth/logout", Description: "Revoke active sessions"},
	})
}
