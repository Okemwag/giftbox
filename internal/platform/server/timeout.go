package server

import (
	"net/http"
	"time"
)

func Timeout(duration time.Duration) Middleware {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, duration, `{"error":{"code":"INTERNAL_ERROR","message":"request timed out"}}`)
	}
}
