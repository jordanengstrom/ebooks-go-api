package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	var book models.Book

	// if result := h.DB.First(&book, id); result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	if result := h.DB.Preload(clause.Associations).First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Authors = updatedBook.Authors
	book.Title = updatedBook.Title
	book.CopyrightYear = updatedBook.CopyrightYear
	book.About = updatedBook.About

	// h.DB.Save(&book)
	// h.DB.Session(&gorm.Session{FullSjk;.aveAssociations: false}).Updates(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
