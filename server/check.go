package main

import (
	"encoding/json"
	"github.com/enescakir/balance"
	"log"
	"net/http"
)

type checkRequest struct {
	Query string `json:"expr"`
}

type checkResponse struct {
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

func checkHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(req.Body)
	var cReq checkRequest
	err := decoder.Decode(&cReq)
	if err != nil {
		panic(err)
	}

	log.Printf("/isbalanced - %q", cReq.Query)

	valid, err := balance.Check(cReq.Query)

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
