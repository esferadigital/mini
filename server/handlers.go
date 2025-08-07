package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/esferachill/mini/shortener"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortCode string `json:"short_code"`
	ShortURL  string `json:"short_url"`
}

func (s *Server) ShortenLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode := shortener.Shorten(s.linkRepo, req.URL)

	response := ShortenResponse{
		ShortCode: shortCode,
		ShortURL:  fmt.Sprintf("http://localhost:8080/%s", shortCode),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) ServeLink(w http.ResponseWriter, r *http.Request) {
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

	originalURL, exists := shortener.GetOriginalURL(s.linkRepo, path)
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
