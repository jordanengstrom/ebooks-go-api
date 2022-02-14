package handlers

import (
	"ebooks/pkg/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func (h handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var validate *validator.Validate = validator.New()

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	w.Header().Add("Content-Type", "application/json")

	var author models.Author
	json.Unmarshal(body, &author)
	err := validate.Struct(author)

	if err != nil {
		log.Error("unable to create author with the given data")
		for _, e := range err.(validator.ValidationErrors) {
			log.Error(e)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result := h.DB.Create(&author); result.Error != nil {
		log.Error("unable to create author with the given data")
		log.Error(result.Error)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		log.WithFields(log.Fields{"firstName": author.FirstName, "middleName": author.MiddleName, "lastName": author.LastName}).Info("successfully created author:")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Created")
	}
}
