package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (h handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var validate *validator.Validate = validator.New()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	log.Info("PUT /api/authors/" + vars["id"])

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	w.Header().Add("Content-Type", "application/json")

	var updatedAuthor models.Author
	json.Unmarshal(body, &updatedAuthor)
	err := validate.Struct(updatedAuthor)

	if err != nil {
		log.Error("unable to update author with the given data")
		for _, e := range err.(validator.ValidationErrors) {
			log.Error(e)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var author models.Author
	if result := h.DB.First(&author, id); result.Error != nil {
		log.Error("author record with id=" + vars["id"] + " not found")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		author.FirstName = updatedAuthor.FirstName
		author.MiddleName = updatedAuthor.MiddleName
		author.LastName = updatedAuthor.LastName

		h.DB.Save(&author)
		log.Info("successfully updated author id=" + vars["id"])
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Updated")
	}
}
