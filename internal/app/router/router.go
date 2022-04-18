package router

import (
	"net/http"

	"github.com/truthray/url-shortening-service/internal/app/storage"
)

func New() http.HandlerFunc {
	data := storage.New()

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGet(w, r, data)
		case http.MethodPost:
			handlePost(w, r, data)
		default:
			http.Error(w, "Only GET or POST requests are allowed", http.StatusMethodNotAllowed)
		}
	}
}
