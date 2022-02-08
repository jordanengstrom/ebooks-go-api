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
)

func (h handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedAuthor models.Author
	json.Unmarshal(body, &updatedAuthor)

	var author models.Author

	if result := h.DB.First(&author, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	author.FirstName = updatedAuthor.FirstName
	author.MiddleName = updatedAuthor.MiddleName
	author.LastName = updatedAuthor.LastName

	h.DB.Save(&author)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
