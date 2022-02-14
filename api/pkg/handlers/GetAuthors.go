package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func (h handler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	log.Info("GET /api/authors")
	var authors []models.Author
	w.Header().Add("Content-Type", "application/json")

	if result := h.DB.Find(&authors); result.Error != nil {
		log.Error("unable to retrieve authors")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(authors)
		log.Info("successfully retrieved " + strconv.Itoa(len(authors)) + " authors")
	}
}
