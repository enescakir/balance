package server

import (
	"github.com/enescakir/balance/server/querylog"
	"log"
	"net/http"
	"time"
)

func (s *Server) log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lw := querylog.NewLoggerWriter(w)
		h(lw, r)
		var cReq checkRequest

		cReq.fromJson(r.Body)

		defer querylog.NewQueryLog(cReq.Query, lw.Status, time.Since(start).Nanoseconds()).Save(s.db)

		defer log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}
