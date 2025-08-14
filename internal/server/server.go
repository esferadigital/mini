package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	port int
}

func NewServer(port int) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/health", health)
	http.HandleFunc("/shorten", s.Shorten)
	http.HandleFunc("/", s.Redirect)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: http.DefaultServeMux,
	}
	return httpServer.ListenAndServe()
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
