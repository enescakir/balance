package server

import (
	"encoding/json"
	"github.com/enescakir/balance/server/querylog"
	"log"
	"net/http"
	"time"
)

func (s *Server) log(h http.HandlerFunc) http.HandlerFunc {
	type request struct {
		Query *string `json:"expr"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lw := querylog.NewLoggerWriter(w)
		h(lw, r)

		var req request
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)

		if err != nil || req.Query == nil {
			return
		}

		s.repo.Store(querylog.NewQueryLog(*req.Query, lw.Status, time.Since(start).Nanoseconds()))

		defer log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, time.Since(start),
		)
	}
}
