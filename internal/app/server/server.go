package server

import (
	"net/http"

	"github.com/truthray/url-shortening-service/internal/app/router"
	"github.com/truthray/url-shortening-service/internal/app/storage"
)

type server struct {
	Addr   string
	Router http.HandlerFunc
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
	http.Handle("/", s.Router)
	return http.ListenAndServe(s.Addr, nil)
}
