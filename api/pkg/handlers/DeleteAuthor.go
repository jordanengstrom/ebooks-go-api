package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (h handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	log.Info("DELETE /api/authors/" + vars["id"])

	var author models.Author
	w.Header().Add("Content-Type", "application/json")

	if result := h.DB.First(&author, id); result.Error != nil {
		log.Error("unable to delete author with id=" + strconv.Itoa(id))
		w.WriteHeader(http.StatusBadRequest)
	} else {
		h.DB.Delete(&author)
		log.Info("successfully deleted author id=" + strconv.Itoa(id))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Deleted")
	}
}
