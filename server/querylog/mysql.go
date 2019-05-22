package querylog

import (
	"database/sql"
	"fmt"
	"github.com/enescakir/balance/server/config"
	"log"
	"strings"
)

// MysqlRepository implements Repository interface for MySQL.
type MysqlRepository struct {
	db *sql.DB
}

// NewMysqlRepository returns newly created MysqlRepository reference with given database.
func NewMysqlRepository(cfg config.Config) *MysqlRepository {
	db := NewMysqlDatabase(cfg)
	Migrate(db)

	return &MysqlRepository{db: db}
}

// Flush drops all tables on MySQL database.
func (r *MysqlRepository) Flush() {
	Rollback(r.db)
}

// Store saves given QueryLog to MySQL database.
func (r *MysqlRepository) Store(l *QueryLog) error {
	_, err := r.db.Exec("INSERT INTO logs (query, Status, response_time) VALUES (?, ?, ?)", l.Query, l.Status, l.ResponseTime)

	if err != nil {
		log.Printf("Can't insert log to DB: %s", err.Error())
		return err
	}

	return nil
}

// FindAll returns all logs for given data range from MySQL database.
func (r *MysqlRepository) FindAll(start string, end string) ([]*QueryLog, error) {
	where, args := buildDateRangeQuery(start, end)

	query := fmt.Sprintf("SELECT * FROM logs %s ORDER BY created_at DESC", where)

	results, err := r.db.Query(query, args...)

	if err != nil {
		log.Printf("Can't get logs from DB: %s", err.Error())
		return nil, err
	}

	defer results.Close()

	return parseLogs(results)
}

// CountByStatus returns status:count pairs for given data range from MySQL database.
func (r *MysqlRepository) CountByStatus(start string, end string) ([]*StatusCount, error) {
	where, args := buildDateRangeQuery(start, end)

	query := fmt.Sprintf("SELECT status, COUNT(*) as count FROM logs %s GROUP BY status", where)

	results, err := r.db.Query(query, args...)

	if err != nil {
		log.Printf("Can't get log counts by status from DB: %s", err.Error())
		return nil, err
	}

	defer results.Close()

	return parseStatusCounts(results)
}

// HistogramBins returns responseTime:count bins for given data range from MySQL database.
func (r *MysqlRepository) HistogramBins(start string, end string) ([]*HistogramBin, error) {
	where, args := buildDateRangeQuery(start, end)

	query := fmt.Sprintf("SELECT * FROM logs %s ORDER BY created_at DESC", where)

	results, err := r.db.Query(query, args...)

	if err != nil {
		log.Printf("Can't get logs from DB: %s", err.Error())
		return nil, err
	}

	defer results.Close()

	logs, err := parseLogs(results)
	if err != nil {
		log.Printf("Can't parse logs %s", err.Error())
		return nil, err
	}

	return createHistogramBins(logs), nil
}

// buildDateRangeQuery generates prepared SQL query for given start and end date.
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

// parseLogs converts MySQL results to QueryLog collection
func parseLogs(results *sql.Rows) ([]*QueryLog, error) {
	logs := make([]*QueryLog, 0)

	for results.Next() {
		var l QueryLog
		err := results.Scan(&l.Id, &l.Query, &l.Status, &l.ResponseTime, &l.CreatedAt)
		if err != nil {
			log.Printf("Can't parse log row: %s", err.Error())
			return nil, err
		}
		logs = append(logs, &l)
	}

	return logs, nil
}

// parseStatusCounts converts MySQL results to StatusCount collection
func parseStatusCounts(results *sql.Rows) ([]*StatusCount, error) {
	counts := make([]*StatusCount, 0)

	for results.Next() {
		var c StatusCount
		err := results.Scan(&c.Status, &c.Count)
		if err != nil {
			log.Printf("Can't parse log status count row: %s", err.Error())
			return nil, err
		}
		counts = append(counts, &c)
	}

	return counts, nil
}
