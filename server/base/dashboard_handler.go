package server

import (
	"net/http"
	"path/filepath"
)

func (s *Server) handleDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		fp := filepath.Join("templates", "dashboard.html")
		http.ServeFile(w, r, fp)
	}
}
