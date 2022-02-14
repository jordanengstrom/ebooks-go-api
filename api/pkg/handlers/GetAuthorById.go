package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (h handler) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	log.Info("GET /api/authors/" + vars["id"])
	var author models.Author
	w.Header().Add("Content-Type", "application/json")

	if result := h.DB.First(&author, id); result.Error != nil {
		log.Error("author record with id=" + vars["id"] + " not found")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(author)
		log.Info("successfully retrieved author id=" + vars["id"])
	}
}
