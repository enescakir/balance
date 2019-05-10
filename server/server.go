// Command server lists all the processes running on your system.
package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/isbalanced", checkHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
