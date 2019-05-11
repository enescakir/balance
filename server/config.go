package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port             int
	DatabaseHost     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabasePort     int
}

func ReadConfig(filename string) Config {
	file, err := os.Open(filename)
	var cfg Config
	if err != nil {
		log.Fatal("Can't open config file")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Fatal("Can't decode config JSON")
	}

	return cfg
}
