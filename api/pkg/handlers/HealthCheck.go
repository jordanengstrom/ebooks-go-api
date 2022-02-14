package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Info("GET /api/ping")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("pong")
}
