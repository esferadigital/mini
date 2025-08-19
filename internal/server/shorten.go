package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esferachill/mini/internal/platform"
	"github.com/esferachill/mini/internal/services"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortCode string `json:"short_code"`
	ShortURL  string `json:"short_url"`
}

func (s *Server) Shorten(w http.ResponseWriter, r *http.Request) {
	env := platform.GetPlatform().Env

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

	shortCode := services.Shorten(req.URL)
	shortURL := fmt.Sprintf("%s://%s:%d/%s", env.HostScheme, env.Host, env.Port, shortCode)

	response := ShortenResponse{
		ShortCode: shortCode,
		ShortURL:  shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
