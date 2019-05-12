package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/enescakir/balance"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) handleCheck() http.HandlerFunc {
	type request struct {
		Query *string `json:"expr"`
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
		var cReq request
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(bodyBytes, &cReq)

		if err != nil || cReq.Query == nil {
			log.Printf("Check handle couldn't parse request")
			fmt.Fprint(w, "`expr` is required")
			return
		}

		// Put body content to request body again, because log middleware will read it
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// Validate given string
		valid, err := balance.Check(*cReq.Query)

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
