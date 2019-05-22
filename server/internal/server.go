// Package server start simple HTTP server for validating parentheses balance of strings
package internal

import (
	"fmt"
	"github.com/enescakir/balance/server/config"
	"github.com/enescakir/balance/server/database"
	"github.com/enescakir/balance/server/querylog"
	"log"
	"net/http"
)

// Server keeps shared dependencies.
type Server struct {
	srv    http.Server
	repo   querylog.Repository
	router *http.ServeMux
	port   int
}

// NewServer returns newly created Server reference with initialized mux and database connection.
func NewServer(cfg config.Config) *Server {
	mux := http.NewServeMux()

	var repo querylog.Repository

	if cfg.Database.Driver == config.MySQL {
		log.Printf("MySQL database is selected")
		db := database.New(cfg)
		database.Migrate(db)
		repo = querylog.NewMysqlRepository(db)
	} else {
		log.Printf("In memory database is selected")
		repo = querylog.NewMemoryRepository()
	}
	s := &Server{repo: repo, router: mux, port: cfg.Port}

	return s
}

// Start initialize HTTP server on given port address
func (s *Server) Start() {
	log.Printf("Listening on port %d\n", s.port)

	s.routes()

	address := fmt.Sprintf(":%d", s.port)
	s.srv = http.Server{Addr: address, Handler: s.router}

	log.Fatal(s.srv.ListenAndServe())
}
