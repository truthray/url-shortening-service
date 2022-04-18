package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	Addr string
}

func New() *Server {
	return &Server{
		Addr: "localhost:8000",
	}
}

func (s *Server) Start() error {
	http.Handle("/", s.handleRequest())
	return http.ListenAndServe(s.Addr, nil)
}

func (s *Server) handleRequest() http.HandlerFunc {
	data := make(map[int]string)
	counter := 0

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
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
			if url, ok := data[code]; ok {
				w.Header().Set("Location", url)
				w.WriteHeader(http.StatusTemporaryRedirect)
				return
			}
			http.Error(w, "Id not found", http.StatusNotFound)

		case http.MethodPost:
			b, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Body is missing", http.StatusBadRequest)
				return
			}
			data[counter] = string(b)
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, s.Addr, "/", counter)
			counter += 1
		default:
			http.Error(w, "Only GET or POST requests are allowed", http.StatusMethodNotAllowed)
		}
	}
}
