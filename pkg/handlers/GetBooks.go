package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm/clause"
)

func (h handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if result := h.DB.Preload(clause.Associations).Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
