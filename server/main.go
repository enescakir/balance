// Command run parenthesis balance server.
package main

import (
	"github.com/enescakir/balance/server/config"
	"github.com/enescakir/balance/server/internal"
)

func main() {
	cfg := config.Read("config/config.json")

	s := internal.NewServer(cfg)

	s.Start()
}
