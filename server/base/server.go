// Package server start simple HTTP server for validating parentheses balance of strings
package server

import (
	"database/sql"
	"fmt"
	"github.com/enescakir/balance/server/querylog"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

// Server keeps shared dependencies.
type Server struct {
	repo   querylog.Repository
	router *http.ServeMux
	port   int
}

// NewServer returns newly created Server reference with initialized mux and database connection.
func NewServer(cfg Config) *Server {
	mux := http.NewServeMux()

	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	db, err := sql.Open("mysql", address)

	if err != nil {
		log.Fatal("Can't not connect database")
	}

	repo := querylog.NewMysqlRepository(db)

	return &Server{repo: repo, router: mux, port: cfg.Port}
}

// Start initialize HTTP server on given port address
func (s *Server) Start() {
	log.Printf("Listening on port %d\n", s.port)

	s.routes()

	address := fmt.Sprintf(":%d", s.port)

	log.Fatal(http.ListenAndServe(address, s.router))
}
