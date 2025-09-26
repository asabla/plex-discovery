package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/asabla/plex-discovery/internal/plex"
)

func NewRouter(home templ.Component, plexClient *plex.Client) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/healthz", http.HandlerFunc(healthHandler))
	mux.Handle("/api/identity", identityHandler(plexClient))
	mux.Handle("/", templ.Handler(home))
	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func identityHandler(client *plex.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if client == nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "plex client not configured"})
			return
		}

		resp, err := client.GetIdentity(r.Context())
		if err != nil {
			log.Printf("plex identity lookup failed: %v", err)
			w.WriteHeader(http.StatusBadGateway)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "plex identity unavailable"})
			return
		}

		if resp.JSON200 == nil {
			w.WriteHeader(http.StatusBadGateway)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "unexpected plex response"})
			return
		}

		if err := json.NewEncoder(w).Encode(resp.JSON200); err != nil {
			log.Printf("failed to encode plex identity: %v", err)
		}
	})
}
