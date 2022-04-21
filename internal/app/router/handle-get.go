package router

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/truthray/url-shortening-service/internal/app/storage"
)

func handleGet(w http.ResponseWriter, r *http.Request, data storage.Storage) {
	params := strings.Split(r.URL.Path, "/")
	if len(params) < 2 {
		http.Error(w, "Query parameter is missing", http.StatusBadRequest)
		return
	}
	code, err := strconv.Atoi(params[len(params)-1])
	if err != nil {
		http.Error(w, "Query parameter is not integer", http.StatusBadRequest)
		return
	}
	if url, ok := data.GetUrl(code); ok {
		w.Header().Set("Location", url)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTemporaryRedirect)
		fmt.Fprint(w, url)
		return
	}
	http.Error(w, "Id not found", http.StatusNotFound)
}
