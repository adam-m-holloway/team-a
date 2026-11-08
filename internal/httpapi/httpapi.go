// Package httpapi wires the service's HTTP routes.
package httpapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/adam-m-holloway/team-a/internal/greet"
)

// NewHandler returns the service's HTTP handler.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handleHealthz)
	mux.HandleFunc("GET /greet", handleGreet)
	return mux
}

func handleHealthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	msg, err := greet.Greet(r.URL.Query().Get("name"))
	if errors.Is(err, greet.ErrEmptyName) {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name is required"})
		return
	}
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": msg})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("httpapi: write response: %v", err)
	}
}
