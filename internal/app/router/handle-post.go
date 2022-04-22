package router

import (
	"fmt"
	"io"
	"net/http"

	"github.com/truthray/url-shortening-service/internal/app/storage"
)

func handlePost(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Body is missing", http.StatusBadRequest)
			return
		}
		stringBody := string(b)

		if stringBody == "" {
			http.Error(w, "Body is missing", http.StatusBadRequest)
			return
		}
		storage.AddURL(stringBody)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "http://localhost:8080/", storage.CurrentIndex())
	}
}
