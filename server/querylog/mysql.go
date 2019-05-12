package querylog

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) *MysqlRepository {
	return &MysqlRepository{db: db}
}

func (r *MysqlRepository) Store(l *QueryLog) error {
	insert, err := r.db.Query("INSERT INTO logs (query, Status, response_time) VALUES (?, ?, ?)", l.Query, l.Status, l.ResponseTime)

	if err != nil {
		log.Printf("Can't insert log to DB: %s", err.Error())
		return err
	}

	defer insert.Close()
	return nil
}

func (r *MysqlRepository) FindAll(start string, end string) ([]*QueryLog, error) {
	where, args := buildDateRangeQuery(start, end)

	query := fmt.Sprintf("SELECT * FROM logs %s ORDER BY created_at DESC", where)

	results, err := r.db.Query(query, args...)

	if err != nil {
		log.Printf("Can't get logs from DB: %s", err.Error())
		return nil, err
	}

	logs := []*QueryLog{}

	for results.Next() {
		var l QueryLog
		err = results.Scan(&l.Id, &l.Query, &l.Status, &l.ResponseTime, &l.CreatedAt)
		if err != nil {
			log.Printf("Can't parse log row: %s", err.Error())
			return nil, err
		}
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *MysqlRepository) GetCountByStatus(start string, end string) ([]*StatusCount, error) {
	where, args := buildDateRangeQuery(start, end)

	query := fmt.Sprintf("SELECT status, COUNT(*) as count FROM logs %s GROUP BY status", where)

	results, err := r.db.Query(query, args...)

	if err != nil {
		log.Printf("Can't get log counts by status from DB: %s", err.Error())
		return nil, err
	}

	counts := []*StatusCount{}

	for results.Next() {
		var c StatusCount
		err = results.Scan(&c.Status, &c.Count)
		if err != nil {
			log.Printf("Can't parse log status count row: %s", err.Error())
			return nil, err
		}
		counts = append(counts, &c)
	}

	return counts, nil
}

func (r *MysqlRepository) GetHistogramBins(start string, end string) ([]*HistogramBin, error) {
	type Bucket struct {
		Value int `json:"value"`
		Count int `json:"count"`
	}

	where, args := buildDateRangeQuery(start, end)

	query := fmt.Sprintf("SELECT ROUND(response_time, -4) AS value, COUNT(*) AS count FROM logs %s GROUP BY value ORDER BY value", where)

	results, err := r.db.Query(query, args...)

	if err != nil {
		log.Printf("Can't get log response time histogram from DB: %s", err.Error())
		return nil, err
	}

	prev := 0
	bins := []*HistogramBin{}

	for results.Next() {
		var b Bucket
		err = results.Scan(&b.Value, &b.Count)
		if err != nil {
			log.Printf("Can't parse log response time histogram row: %s", err.Error())
			return nil, err
		}
		for b.Value > prev {
			label := fmt.Sprintf("%d-%d", prev/1000, (prev+10000)/1000)
			bins = append(bins, &HistogramBin{label, 0})
			prev = prev + 10000
		}
		prev = prev + 10000
		label := fmt.Sprintf("%d-%d", b.Value/1000, (b.Value+10000)/1000)
		bins = append(bins, &HistogramBin{label, b.Count})
	}

	return bins, nil
}

func buildDateRangeQuery(start string, end string) (string, []interface{}) {
	var placeholders []string
	var values []interface{}

	if start != "" {
		placeholders = append(placeholders, "created_at > ?")
		values = append(values, start)
	}

	if end != "" {
		placeholders = append(placeholders, "created_at < ?")
		values = append(values, end)
	}

	where := ""
	if len(placeholders) > 0 {
		where = fmt.Sprintf("WHERE %s", strings.Join(placeholders, " AND "))
	}

	return where, values
}
