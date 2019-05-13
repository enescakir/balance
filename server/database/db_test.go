package database

import (
	"github.com/enescakir/balance/server/config"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := config.Read("../config/config.mysql.json")

	db := New(cfg)

	err := db.Ping()
	if err != nil {
		t.Errorf("Can't connect database: %s", err.Error())
	}
}

func TestErrors(t *testing.T) {
	cfg := config.Read("")

	db := New(cfg)

	Migrate(db)

	_, err := db.Query("SELECT 1 FROM logs LIMIT 1;")

	if err == nil {
		t.Errorf("Table shouldn't exist")
	}

	Rollback(db)

	_, err = db.Query("SELECT 1 FROM logs LIMIT 1;")
	if err == nil {
		t.Errorf("Table not dropped")
	}
}
