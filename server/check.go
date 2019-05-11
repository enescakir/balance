package main

import (
	"encoding/json"
	"github.com/enescakir/balance"
	"net/http"
)

func (s *Server) handleCheck() http.HandlerFunc {
	type request struct {
		Query string `json:"expr"`
	}

	type response struct {
		Valid bool   `json:"valid"`
		Error string `json:"error,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse JSON request to struct
		decoder := json.NewDecoder(r.Body)
		var cReq request
		err := decoder.Decode(&cReq)
		if err != nil {
			panic(err)
		}

		// Validate given string
		valid, err := balance.Check(cReq.Query)

		// Convert result to JSON and return it
		var cRes response
		cRes.Valid = valid
		if err != nil {
			cRes.Error = err.Error()
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(cRes)

		if err != nil {
			panic(err)
		}
	}
}
