package server

import (
	"context"
	"net/http"
	"os"

	"github.com/esferachill/mini/internal/components"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	component := components.Home()
	component.Render(context.Background(), os.Stdout)
}
