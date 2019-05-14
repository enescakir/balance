package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

// handleLogIndex returns all logs as JSON for given date range
func (s *Server) handleLogIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")

		logs, err := s.repo.FindAll(start, end)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(logs)

		if err != nil {
			log.Printf("Log handler can't convert logs to JSON")
			http.Error(w, "JSON convert error", http.StatusInternalServerError)
			return
		}
	}
}

// handleLogStatusCounts returns status:count pairs at given date range
func (s *Server) handleLogStatusCounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")

		counts, err := s.repo.CountByStatus(start, end)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(counts)

		if err != nil {
			log.Printf("Log handler can't convert status counts to JSON")
			http.Error(w, "JSON convert error", http.StatusInternalServerError)
			return
		}
	}
}

// handleLogResponseHistogram returns label:responseTime bins at given date range.
func (s *Server) handleLogResponseHistogram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")

		bins, err := s.repo.HistogramBins(start, end)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(bins)

		if err != nil {
			log.Printf("Log handler can't convert histogram bins to JSON")
			http.Error(w, "JSON convert error", http.StatusInternalServerError)
			return
		}
	}
}
