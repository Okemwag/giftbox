package server

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Endpoint struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

func RegisterModuleRoutes(mux *http.ServeMux, module string, endpoints []Endpoint) {
	base := "/v1/" + strings.Trim(module, "/")

	mux.HandleFunc(base, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeMethodNotAllowed(w)
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{
			"module":    module,
			"endpoints": endpoints,
		})
	})

	mux.HandleFunc(base+"/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeMethodNotAllowed(w)
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{
			"module": module,
			"status": "ready",
		})
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeMethodNotAllowed(w http.ResponseWriter) {
	writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}
