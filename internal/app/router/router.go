package router

import (
	"github.com/gorilla/mux"
	"github.com/truthray/url-shortening-service/internal/app/storage"
)

func New(storage storage.Storage) *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/{id}", handleGet(storage)).Methods("GET")
	r.HandleFunc("/", handlePost(storage)).Methods("POST")

	return r
}
