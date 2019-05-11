package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type QueryLog struct {
	Id           int       `json:"id"`
	Query        string    `json:"query"`
	Status       LogStatus `json:"status"`
	ResponseTime int64     `json:"response_time"`
	CreatedAt    time.Time `json:"created_at"`
}

type QueryLogs []QueryLog

type LogStatus int

const (
	Unknown    LogStatus = 0
	Balanced   LogStatus = 1
	Unbalanced LogStatus = 2
	Invalid    LogStatus = 3
)

func newQueryLog(query string, status LogStatus, rTime int64) *QueryLog {
	return &QueryLog{Query: query, Status: status, ResponseTime: rTime}
}

func (l *QueryLog) save(db *sql.DB) {
	q := fmt.Sprintf("INSERT INTO logs (query, status, response_time) VALUES (%q, %d, %d)", l.Query, l.Status, l.ResponseTime)
	insert, err := db.Query(q)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer insert.Close()
}

func (s *Server) log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lw := NewLoggerWriter(w)
		h(lw, r)
		var cReq checkRequest

		cReq.fromJson(r.Body)

		defer newQueryLog(cReq.Query, lw.status, time.Since(start).Nanoseconds()).save(s.db)

		defer log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}

type LoggerWriter struct {
	writer http.ResponseWriter
	status LogStatus
}

func NewLoggerWriter(w http.ResponseWriter) *LoggerWriter {
	var lw LoggerWriter
	lw.writer = w
	return &lw
}

func (w *LoggerWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *LoggerWriter) WriteHeader(statusCode int) {
	w.writer.WriteHeader(statusCode)
}

func (w *LoggerWriter) Write(b []byte) (int, error) {
	var cRes checkResponse
	err := json.Unmarshal(b, &cRes)
	if cRes.Valid {
		w.status = 1
	} else {
		w.status = 2
	}
	if err != nil {
		panic(err)
	}

	return w.writer.Write(b)
}
