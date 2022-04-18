package server

import (
	"net/http"

	"github.com/truthray/url-shortening-service/internal/app/router"
)

type Server struct {
	Addr   string
	Router http.HandlerFunc
}

func New() *Server {
	r := router.New()

	return &Server{
		Addr:   "localhost:8080",
		Router: r,
	}
}

func (s *Server) Start() error {
	http.Handle("/", s.Router)
	return http.ListenAndServe(s.Addr, nil)
}
