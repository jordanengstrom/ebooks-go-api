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

	subrouter := router.PathPrefix("/api").Subrouter()

	subrouter.HandleFunc("/authors", h.GetAuthors).Methods(http.MethodGet)
	subrouter.HandleFunc("/authors", h.CreateAuthor).Methods(http.MethodPost)
	subrouter.HandleFunc("/authors/{id}", h.GetAuthorById).Methods(http.MethodGet)
	subrouter.HandleFunc("/authors/{id}", h.UpdateAuthor).Methods(http.MethodPut)
	subrouter.HandleFunc("/authors/{id}", h.DeleteAuthor).Methods(http.MethodDelete)

	subrouter.HandleFunc("/books", h.GetBooks).Methods(http.MethodGet)
	subrouter.HandleFunc("/books", h.CreateBook).Methods(http.MethodPost)
	subrouter.HandleFunc("/books/{id}", h.GetBookById).Methods(http.MethodGet)
	subrouter.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	subrouter.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	subrouter.HandleFunc("/ping", h.HealthCheck).Methods(http.MethodGet)

	port := ":8080"
	log.Printf("API is running on http://localhost%v\n", port)
	http.ListenAndServe(port, router)
}
