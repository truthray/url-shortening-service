package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/truthray/url-shortening-service/internal/app/storage"
)

func handleGet(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		code, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Query parameter is not integer", http.StatusBadRequest)
			return
		}
		if url, ok := storage.GetUrl(code); ok {
			w.Header().Set("Location", url)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTemporaryRedirect)
			fmt.Fprint(w, url)
			return
		}
		http.Error(w, "Id not found", http.StatusNotFound)
	}
}
