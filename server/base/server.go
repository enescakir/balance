// Command Server run parenthesis balance Server.
package server

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Server struct {
	db     *sql.DB
	router *http.ServeMux
	port   int
}

func NewServer(cfg Config) *Server {
	mux := http.NewServeMux()

	dbAddress := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)
	db, err := sql.Open("mysql", dbAddress)

	if err != nil {
		panic(err.Error())
	}

	s := Server{db: db, router: mux, port: cfg.Port}

	return &s
}

func (s *Server) Start() {
	log.Printf("Listening on port %d", s.port)
	s.routes()
	address := fmt.Sprintf(":%d", s.port)
	log.Fatal(http.ListenAndServe(address, s.router))
}
