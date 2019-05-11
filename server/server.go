// Command server run parenthesis balance server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type server struct {
	//db     *someDatabase
	router *http.ServeMux
	port   int
}

func (s *server) routes() {
	s.router.HandleFunc("/isbalanced", s.log(s.handleCheck()))
	//s.router.HandleFunc("/about", s.handleAbout())
	//s.router.HandleFunc("/", s.handleIndex())
	//s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))

}

func (s *server) log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)

	}
}

func (s *server) start() {
	log.Printf("Listening on port %d", s.port)
	s.routes()
	address := fmt.Sprintf(":%d", s.port)
	log.Fatal(http.ListenAndServe(address, s.router))
}

func main() {
	mux := http.NewServeMux()

	s := server{router: mux, port: 8080}

	s.start()
}
