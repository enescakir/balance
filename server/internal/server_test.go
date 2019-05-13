package internal

import (
	"context"
	"github.com/enescakir/balance/server/config"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	cfg := config.Read("../config/config.test.json")

	s := NewServer(cfg)
	go s.Start()

	req, _ := http.NewRequest("GET", "/logs", nil)
	rec := httptest.NewRecorder()

	http.HandlerFunc(s.handleLogIndex()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Error("Server couldn't start")
	}

	err := s.srv.Shutdown(context.Background())

	if err != nil {
		t.Error("Server couldn't stop")
	}

}

func TestLogHandler(t *testing.T) {
	cfg := config.Read("../config/config.test.json")

	s := NewServer(cfg)
	s.routes()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/logs", nil)
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Error("/logs endpoint has problem")
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/logs", nil)
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("/logs endpoint has problem: %v", rec.Code)
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/logs/status", nil)
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Error("/logs/status endpoint has problem")
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/logs/status", nil)
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("/logs/status endpoint has problem: %v", rec.Code)
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/logs/histogram", nil)
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Error("/logs/histogram endpoint has problem")
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/logs/histogram", nil)
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("/logs/histogram endpoint has problem: %v", rec.Code)
	}

}

func TestCheckHandler(t *testing.T) {
	cfg := config.Read("../config/config.test.json")

	s := NewServer(cfg)
	s.routes()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/isbalanced", strings.NewReader("{\"expr\": \"()(){}\"}"))
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("/isbalanced endpoint has problem: %v", rec.Code)
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/isbalanced", strings.NewReader("{\"expr\": \"()(){|}\"}"))
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("/isbalanced endpoint has problem: %v", rec.Code)
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/isbalanced", strings.NewReader("{\"expr\": \"()(){}\"}"))
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("/isbalanced endpoint has problem: %v", rec.Code)
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/isbalanced", strings.NewReader("{\"expr2\": \"()(){}\"}"))
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("/isbalanced endpoint has problem: %v", rec.Code)
	}

	s.repo.Flush()
}

func TestDashboardHandler(t *testing.T) {
	cfg := config.Read("../config/config.test.json")

	s := NewServer(cfg)
	s.routes()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	s.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("/ endpoint StatusOK has problem: %v", rec.Code)
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/", nil)
	s.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("/ endpoint StatusMethodNotAllowed has problem: %v", rec.Code)
	}
}

func TestMysqlServer(t *testing.T) {
	cfg := config.Read("../config/config.mysql.json")

	s := NewServer(cfg)
	go s.Start()

	req, _ := http.NewRequest("GET", "/logs", nil)
	rec := httptest.NewRecorder()

	http.HandlerFunc(s.handleLogIndex()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Error("MySQL Server couldn't start")
	}

	err := s.srv.Shutdown(context.Background())

	if err != nil {
		t.Error("MySQL Server couldn't stop")
	}

}
