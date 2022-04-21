package server

import (
	"net/http"

	"github.com/truthray/url-shortening-service/internal/app/router"
)

type server struct {
	Addr   string
	Router http.HandlerFunc
}

func New() *server {
	r := router.New()

	return &server{
		Addr:   "localhost:8080",
		Router: r,
	}
}

func (s *server) Start() error {
	http.Handle("/", s.Router)
	return http.ListenAndServe(s.Addr, nil)
}
