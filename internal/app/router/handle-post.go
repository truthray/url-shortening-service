package router

import (
	"fmt"
	"io"
	"net/http"

	"github.com/truthray/url-shortening-service/internal/app/storage"
)

func handlePost(w http.ResponseWriter, r *http.Request, data storage.Storage) {
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
	data.AddUrl(stringBody)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "localhost:8080/", data.CurrentIndex())
}
