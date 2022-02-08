package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var author models.Author
	json.Unmarshal(body, &author)

	if result := h.DB.Create(&author); result.Error != nil {
		fmt.Println(result.Error)
	}

	// send 201
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
