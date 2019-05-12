// Package server start simple HTTP server for validating parentheses balance of strings
package internal

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
	s := &Server{repo: repo, router: mux, port: cfg.Port}
	s.migrate(db)

	return s
}

// Start initialize HTTP server on given port address
func (s *Server) Start() {
	log.Printf("Listening on port %d\n", s.port)

	s.routes()

	address := fmt.Sprintf(":%d", s.port)

	log.Fatal(http.ListenAndServe(address, s.router))
}

// migrate creates data tables if not exists
func (s *Server) migrate(db *sql.DB) {
	migration := `
	CREATE TABLE IF NOT EXISTS logs (
 		id int(11) unsigned NOT NULL AUTO_INCREMENT,
 		query text COLLATE utf8mb4_general_ci,
 		status int(11) NOT NULL DEFAULT '0',
 		response_time bigint(11) NOT NULL,
 		created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  		PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=141 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;`

	_, err := db.Exec(migration)

	if err != nil {
		log.Fatalf("Can't not create logs table: %s", err.Error())
	}

}
