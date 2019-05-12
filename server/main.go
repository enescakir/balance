// Command run parenthesis balance server.
package main

import server "github.com/enescakir/balance/server/internal"

func main() {
	config := server.ReadConfig("config.json")

	s := server.NewServer(config)

	s.Start()
}
