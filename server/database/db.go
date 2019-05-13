package database

import (
	"database/sql"
	"fmt"
	"github.com/enescakir/balance/server/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// New creates new SQL DB instance end returns its reference.
func New(cfg config.Config) *sql.DB {
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	db, err := sql.Open("mysql", address)

	if err != nil {
		log.Fatalf("Can't not connect database: %s", err.Error())
	}

	return db
}

// Migrate creates data tables if not exists.
func Migrate(db *sql.DB) {
	migration := `
	CREATE TABLE IF NOT EXISTS logs (
 		id int(11) unsigned NOT NULL AUTO_INCREMENT,
 		query text COLLATE utf8mb4_general_ci,
 		status int(11) NOT NULL DEFAULT '0',
 		response_time bigint(11) NOT NULL,
 		created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  		PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=141 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;`

	_, err := db.Exec(migration)

	if err != nil {
		log.Printf("Can't not create logs table: %s", err.Error())
	}
}

// Rollback drops data tables.
func Rollback(db *sql.DB) {
	migration := "DROP TABLE IF EXISTS logs"

	_, err := db.Exec(migration)

	if err != nil {
		log.Printf("Can't not drop logs table: %s", err.Error())
	}
}
