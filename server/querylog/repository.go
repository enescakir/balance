package querylog

import (
	"database/sql"
	"fmt"
	"log"
)

func (l *QueryLog) Save(db *sql.DB) {
	q := fmt.Sprintf("INSERT INTO logs (query, Status, response_time) VALUES (%q, %d, %d)", l.Query, l.Status, l.ResponseTime)
	insert, err := db.Query(q)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer insert.Close()
}
