package main

func main() {
	config := ReadConfig("config.json")

	s := NewServer(config)

	s.Start()
}
