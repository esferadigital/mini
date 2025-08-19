package server

import (
	"fmt"
	"net/http"

	"github.com/esferachill/mini/internal/components"
	"github.com/esferachill/mini/internal/platform"
	"github.com/esferachill/mini/internal/services"
)

func (s *Server) ShortenSubmit(w http.ResponseWriter, r *http.Request) {
	env := platform.GetPlatform().Env

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode := services.Shorten(url)
	shortURL := fmt.Sprintf("%s://%s:%d/%s", env.HostScheme, env.Host, env.Port, shortCode)

	components.Success(shortURL).Render(r.Context(), w)
}
