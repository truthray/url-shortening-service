package main

import (
	"log"
	"github.com/truthray/url-shortening-service/internal/app/server"
)

func main() {
	s := server.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
