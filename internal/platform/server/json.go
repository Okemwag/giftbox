package server

import (
	"encoding/json"
	"net/http"

	apperrors "github.com/Okemwag/giftbox/internal/shared/errors"
)

func JSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, err error) {
	status, payload := apperrors.ToHTTP(err)
	WriteJSON(w, status, map[string]apperrors.HTTPError{"error": payload})
}

func DecodeJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		return apperrors.Wrap(apperrors.CodeInvalidInput, "invalid JSON request body", err)
	}
	return nil
}
