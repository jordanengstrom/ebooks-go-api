package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var author models.Author

	if result := h.DB.First(&author, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&author)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
