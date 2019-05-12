package querylog

import (
	"encoding/json"
	"net/http"
)

type LoggerWriter struct {
	writer http.ResponseWriter
	Status Status
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
	type response struct {
		Valid bool   `json:"valid"`
		Error string `json:"error,omitempty"`
	}

	var r response
	err := json.Unmarshal(b, &r)
	if r.Valid {
		w.Status = 1
	} else {
		w.Status = 2
	}
	if err != nil {
		panic(err)
	}

	return w.writer.Write(b)
}
