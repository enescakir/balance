package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) handleAllLogs() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// if only one expected
		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")
		log.Printf("%q - %q \n", start, end)

		bq := "SELECT * FROM logs"

		var results *sql.Rows
		var err error

		if start != "" && end != "" {
			results, err = s.db.Query(bq+" WHERE created_at > ? AND created_at < ?", start, end)
		} else if start != "" {
			results, err = s.db.Query(bq+" WHERE created_at > ?", start)
		} else if end != "" {
			results, err = s.db.Query(bq+" WHERE created_at < ?", end)
		} else {
			results, err = s.db.Query(bq)
		}

		// Execute the query
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		logs := QueryLogs{}
		for results.Next() {
			var l QueryLog
			// for each row, scan the result into our tag composite object
			err = results.Scan(&l.Id, &l.Query, &l.Status, &l.ResponseTime, &l.CreatedAt)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			logs = append(logs, l)

		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(logs)

		if err != nil {
			panic(err)
		}
	}
}
