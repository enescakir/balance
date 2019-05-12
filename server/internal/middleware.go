package internal

import (
	"encoding/json"
	"github.com/enescakir/balance/server/querylog"
	"log"
	"net/http"
	"time"
)

// log intercepts and saves check balance query to data storage
func (s *Server) log(h http.HandlerFunc) http.HandlerFunc {

	// request represents logged request.
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

		q := querylog.NewQueryLog(*req.Query, lw.Status, time.Since(start).Nanoseconds())
		s.repo.Store(q)

		defer log.Printf("%s\t%10s %10dms %10q", r.Method, r.RequestURI, q.ResponseTime/1000, q.Query)
	}
}
