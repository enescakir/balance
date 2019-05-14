package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	expected := Config{
		Port: 8080,
		Database: DatabaseConfig{
			Host:     "127.0.0.1",
			Name:     "balance",
			User:     "root",
			Password: "secret",
			Port:     3306,
		},
	}
	data := `
		{
		  "Port": 8080,
		  "Database": {
			"Host": "127.0.0.1",
			"Name": "balance",
			"User": "root",
			"Password": "secret",
			"Port": 3306
		  }
		}
	`

	_ = ioutil.WriteFile("./test.json", []byte(data), 0644)

	actual := Read("./test.json")
	_ = os.Remove("./test.json")

	if expected.Port != actual.Port ||
		expected.Database.Port != actual.Database.Port ||
		expected.Database.Password != actual.Database.Password ||
		expected.Database.User != actual.Database.User ||
		expected.Database.Name != actual.Database.Name ||
		expected.Database.Host != actual.Database.Host {
		t.Errorf("Read config error Actual: %v  Expected: %v", actual, expected)
	}

	actual = Read("./dummy.json")
	if actual.Port != 8080 || actual.Database.Port != 3306 {
		t.Errorf("Read config error Actual: %v  Expected: %v", actual, expected)
	}
	_ = os.Setenv("DATABASE_HOST", "localhost:3306")
	actual = Read("./dummy.json")
	if actual.Database.Driver != MySQL {
		t.Errorf("Read config error Actual: %v  Expected: %v", actual.Database.Driver, MySQL)
	}

}
