package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

// Config represents customizable variables of server
type Config struct {
	Port     int
	Database DatabaseConfig
}

// DatabaseConfig represents database credentials
type DatabaseConfig struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     int
}

// Read imports config values from file and environment
func Read(filename string) Config {
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
		if port == "" {
			port = "3306"
		}
		cfg.Database.Port, _ = strconv.Atoi(port)
		cfg.Database.Name = os.Getenv("DATABASE_NAME")
		cfg.Database.User = os.Getenv("DATABASE_USER")
		cfg.Database.Password = os.Getenv("DATABASE_PASSWORD")
	}

	return cfg
}
