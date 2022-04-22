package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/truthray/url-shortening-service/internal/app/router"
	"github.com/truthray/url-shortening-service/internal/app/storage"
)

type server struct {
	Addr   string
	Router *mux.Router
}

func New() *server {
	s := storage.New()
	r := router.New(s)

	return &server{
		Addr:   "localhost:8080",
		Router: r,
	}
}

func (s *server) Start() error {
	server := &http.Server{
		Handler: s.Router,
		Addr:    s.Addr,
	}
	return server.ListenAndServe()
}
