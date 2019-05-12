package querylog

import (
	"encoding/json"
	"net/http"
)

// LoggerWriter intercepts server response and detects the status of response in the server middleware.
type LoggerWriter struct {
	writer http.ResponseWriter
	Status Status
}

// NewLoggerWriter returns newly created LoggerWriter reference.
func NewLoggerWriter(w http.ResponseWriter) *LoggerWriter {
	return &LoggerWriter{w, Unknown}
}

// Header returns the header map of http.ResponseWriter
func (w *LoggerWriter) Header() http.Header {
	return w.writer.Header()
}

// WriteHeader writes status code to header via http.ResponseWriter.
func (w *LoggerWriter) WriteHeader(statusCode int) {
	w.writer.WriteHeader(statusCode)
}

// Write decodes returned response and detects the status of response, then continues normal lifecycle/
func (w *LoggerWriter) Write(b []byte) (int, error) {
	type response struct {
		Valid bool   `json:"valid"`
		Error string `json:"error,omitempty"`
	}

	var r response
	_ = json.Unmarshal(b, &r)
	if r.Valid {
		w.Status = Balanced
	} else {
		w.Status = Unbalanced
	}

	return w.writer.Write(b)
}
