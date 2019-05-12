package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port     int
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     int
}

func ReadConfig(filename string) Config {
	file, err := os.Open(filename)
	var cfg Config
	cfg.Port = 8080

	if err != nil {
		log.Println("Can't open config file")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Println("Can't decode config JSON")
		cfg.Database.Host = os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		cfg.Database.Port, _ = strconv.Atoi(port)
		cfg.Database.Name = os.Getenv("DATABASE_NAME")
		cfg.Database.User = os.Getenv("DATABASE_USER")
		cfg.Database.Password = os.Getenv("DATABASE_PASSWORD")
	}

	fmt.Println(cfg)
	return cfg
}
