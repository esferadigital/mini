package server

import (
	"fmt"
	"net/http"

	"github.com/esfxra/mini/repo"
)

type Server struct {
	port     int
	linkRepo *repo.LinkRepository
}

func NewServer(port int, linkRepo *repo.LinkRepository) *Server {
	return &Server{
		port:     port,
		linkRepo: linkRepo,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/shorten", s.ShortenLink)
	http.HandleFunc("/", s.ServeLink)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: http.DefaultServeMux,
	}
	return httpServer.ListenAndServe()
}
