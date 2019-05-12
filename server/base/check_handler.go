package server

import (
	"bytes"
	"encoding/json"
	"github.com/enescakir/balance"
	"io"
	"io/ioutil"
	"net/http"
)

type checkRequest struct {
	Query string `json:"expr"`
}

func (req *checkRequest) fromJson(body io.ReadCloser) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(req)
	if err != nil {
		panic(err)
	}
}

type checkResponse struct {
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

func (res *checkResponse) fromJson(body io.ReadCloser) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(res)
	if err != nil {
		panic(err)
	}
}

func (s *Server) handleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse JSON request to struct
		var cReq checkRequest
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		cReq.fromJson(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// Validate given string
		valid, err := balance.Check(cReq.Query)

		// Convert result to JSON and return it
		var cRes checkResponse
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
