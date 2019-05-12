package main

import server "github.com/enescakir/balance/server/base"

func main() {
	config := server.ReadConfig("config.json")

	s := server.NewServer(config)

	s.Start()
}
