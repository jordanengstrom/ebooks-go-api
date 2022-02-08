package main

import (
	"ebooks/pkg/db"
	"ebooks/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()

	router.HandleFunc("/authors", h.GetAuthors).Methods(http.MethodGet)
	router.HandleFunc("/authors", h.CreateAuthor).Methods(http.MethodPost)
	router.HandleFunc("/authors/{id}", h.GetAuthorById).Methods(http.MethodGet)
	router.HandleFunc("/authors/{id}", h.UpdateAuthor).Methods(http.MethodPut)
	router.HandleFunc("/authors/{id}", h.DeleteAuthor).Methods(http.MethodDelete)

	router.HandleFunc("/books", h.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", h.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":8080", router)
}
