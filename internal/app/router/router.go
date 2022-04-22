package router

import (
	"net/http"

	"github.com/truthray/url-shortening-service/internal/app/storage"
)

func New(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGet(w, r, storage)
		case http.MethodPost:
			handlePost(w, r, storage)
		default:
			http.Error(w, "Only GET or POST requests are allowed", http.StatusMethodNotAllowed)
		}
	}
}
