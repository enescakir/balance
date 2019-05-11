package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) handleLogIndex() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// if only one expected
		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")

		bq := "SELECT * FROM logs"

		var results *sql.Rows
		var err error

		if start != "" && end != "" {
			results, err = s.db.Query(bq+" WHERE created_at > ? AND created_at < ? ORDER BY created_at DESC", start, end)
		} else if start != "" {
			results, err = s.db.Query(bq+" WHERE created_at > ? ORDER BY created_at DESC", start)
		} else if end != "" {
			results, err = s.db.Query(bq+" WHERE created_at < ? ORDER BY created_at DESC", end)
		} else {
			results, err = s.db.Query(bq + " ORDER BY created_at DESC")
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

func (s *Server) handleLogStatusCounts() http.HandlerFunc {
	type StatusCount struct {
		Status LogStatus `json:"status"`
		Count  int       `json:"count"`
	}

	type StatusCounts []StatusCount

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// if only one expected
		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")

		bq := "SELECT status, COUNT(*) as count FROM logs"

		var results *sql.Rows
		var err error

		if start != "" && end != "" {
			results, err = s.db.Query(bq+" WHERE created_at > ? AND created_at < ? GROUP BY status", start, end)
		} else if start != "" {
			results, err = s.db.Query(bq+" WHERE created_at > ? GROUP BY status", start)
		} else if end != "" {
			results, err = s.db.Query(bq+" WHERE created_at < ? GROUP BY status", end)
		} else {
			results, err = s.db.Query(bq + " GROUP BY status")
		}

		// Execute the query
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		counts := StatusCounts{}
		for results.Next() {
			var c StatusCount
			// for each row, scan the result into our tag composite object
			err = results.Scan(&c.Status, &c.Count)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			counts = append(counts, c)
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(counts)

		if err != nil {
			panic(err)
		}
	}
}

func (s *Server) handleLogResponseHistogram() http.HandlerFunc {
	type Bucket struct {
		Value int `json:"value"`
		Count int `json:"count"`
	}

	type Bin struct {
		Label string `json:"label"`
		Count int    `json:"count"`
	}

	type Bins []Bin

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// if only one expected
		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")

		bq := "SELECT ROUND(response_time, -4) AS value, COUNT(*) AS count FROM logs"

		var results *sql.Rows
		var err error

		if start != "" && end != "" {
			results, err = s.db.Query(bq+" WHERE created_at > ? AND created_at < ? GROUP BY value ORDER BY value", start, end)
		} else if start != "" {
			results, err = s.db.Query(bq+" WHERE created_at > ? GROUP BY value ORDER BY value", start)
		} else if end != "" {
			results, err = s.db.Query(bq+" WHERE created_at < ? GROUP BY value ORDER BY value", end)
		} else {
			results, err = s.db.Query(bq + " GROUP BY value ORDER BY value")
		}

		// Execute the query
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		prev := 0
		bins := Bins{}

		for results.Next() {
			var b Bucket
			// for each row, scan the result into our tag composite object
			err = results.Scan(&b.Value, &b.Count)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			for b.Value > prev {
				label := fmt.Sprintf("%d-%d", prev/1000, (prev+10000)/1000)
				bins = append(bins, Bin{label, 0})
				prev = prev + 10000
			}
			prev = prev + 10000
			label := fmt.Sprintf("%d-%d", b.Value/1000, (b.Value+10000)/1000)
			bins = append(bins, Bin{label, b.Count})
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(bins)

		if err != nil {
			panic(err)
		}
	}
}
