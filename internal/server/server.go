package server

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/esferachill/mini/internal/components"
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
	http.HandleFunc("/shorten-submit", s.ShortenSubmit)
	http.HandleFunc("/", s.Redirect)

	// Site
	home := components.Home()
	http.Handle("/home", templ.Handler(home))

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: http.DefaultServeMux,
	}
	return httpServer.ListenAndServe()
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
