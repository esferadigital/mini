package server

import (
	"net/http"
	"strings"

	"github.com/esferachill/mini/internal/services"
)

func (s *Server) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/")

	// Handle root path
	if path == "" {
		w.Write([]byte("Mini URL Shortener Service"))
		return
	}

	// Handle favicon requests
	if path == "favicon.ico" {
		http.NotFound(w, r)
		return
	}

	targetURL, err := services.Target(path)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, targetURL, http.StatusMovedPermanently)
}
