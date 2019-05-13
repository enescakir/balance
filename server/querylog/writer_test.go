package querylog

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewLoggerWriter(t *testing.T) {

	w := httptest.NewRecorder()
	lw := NewLoggerWriter(w)

	if lw.Status != Unknown {
		t.Errorf("LoggerWriter status Actual: %v  Expected: %v", lw.Status, Unknown)
	}

	if lw.writer == nil {
		t.Errorf("LoggerWriter writer is nil")
	}
}

func TestLoggerWriter_Header(t *testing.T) {
	w := httptest.NewRecorder()
	lw := NewLoggerWriter(w)

	lw.WriteHeader(http.StatusAccepted)
	if len(lw.Header()) != 0 {
		t.Errorf("LoggerWriter header should be empty")
	}

	i, err := lw.Write([]byte{})

	if err != nil && i != 0 {
		t.Errorf("LoggerWriter could't write")
	}

	res := "{\"valid\": true}"

	i, err = lw.Write([]byte(res))

	if err != nil && i != 0 {
		t.Errorf("LoggerWriter could't write")
	}

	if lw.Status != Balanced {
		t.Errorf("LoggerWriter status Actual: %v  Expected: %v", lw.Status, Balanced)
	}

	res = "{\"valid\": false, \"error\": \"Mismatch error\"}"

	i, err = lw.Write([]byte(res))

	if err != nil && i != 0 {
		t.Errorf("LoggerWriter could't write")
	}

	if lw.Status != Unbalanced {
		t.Errorf("LoggerWriter status Actual: %v  Expected: %v", lw.Status, Unbalanced)
	}
}
