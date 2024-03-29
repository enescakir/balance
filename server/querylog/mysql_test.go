package querylog

import (
	"github.com/enescakir/balance/server/config"
	"testing"
	"time"
)

var cases = []struct {
	query  string
	status Status
	time   int64
}{
	{"[][][]()", Balanced, 4500},
	{"[][][][]()", Balanced, 10000},
	{"[][|][]()", Invalid, 10050},
	{"[][][[]()", Unbalanced, 40005},
}

func TestingConfig() config.Config {
	cfg := config.Read("../config/config.mysql.json")
	return cfg
}

func TestMysqlRepository(t *testing.T) {
	cfg := TestingConfig()
	repo := NewMysqlRepository(cfg)
	Rollback(repo.db)
	Migrate(repo.db)
	defer repo.Flush()

	before := time.Now().Add(-time.Hour).Format(TimeFormat)
	after := time.Now().Add(time.Hour).Format(TimeFormat)

	// Test Store()
	for _, c := range cases {
		err := repo.Store(New(c.query, c.status, c.time))

		if err != nil {
			t.Errorf("Query %q log can't store: %s", c.query, err.Error())
		}
	}

	// Test FindAll(start, end)
	logs, err := repo.FindAll("", "")
	if err != nil {
		t.Errorf("Repo can't get all logs: %s", err.Error())
	}

	if len(logs) != len(cases) {
		t.Errorf("Find all count not matched  Actual: %d  Expected: %d", len(logs), len(cases))
	}

	logs, err = repo.FindAll(before, after)
	if err != nil {
		t.Errorf("Repo can't get all logs in date range: %s", err.Error())
	}

	if len(logs) != len(cases) {
		t.Errorf("Find all count not matched in date range  Actual: %d  Expected: %d", len(logs), len(cases))
	}

	logs, err = repo.FindAll(before, before)
	if err != nil {
		t.Errorf("Repo can't get all logs in date range: %s", err.Error())
	}

	if len(logs) != 0 {
		t.Errorf("Find all count not matched in date range  Actual: %d  Expected: %d", len(logs), 0)
	}

	// Test CountByStatus(start, end)
	counts, err := repo.CountByStatus("", "")
	for _, pair := range counts {
		if pair.Status == Balanced && pair.Count != 2 {
			t.Errorf("Balanced logs counts Actual: %d  Expected: %d", pair.Count, 2)
		} else if pair.Status == Unbalanced && pair.Count != 1 {
			t.Errorf("Unbalanced logs counts Actual: %d  Expected: %d", pair.Count, 1)
		} else if pair.Status == Invalid && pair.Count != 1 {
			t.Errorf("Invalid logs counts Actual: %d  Expected: %d", pair.Count, 1)
		}
	}

	// Test CountByStatus(start, end)
	bins, err := repo.HistogramBins("", "")
	for _, bin := range bins {
		if bin.Label == "0-10" && bin.Count != 1 {
			t.Errorf("Histogram bin %s Actual: %d  Expected: %d", bin.Label, bin.Count, 1)
		} else if bin.Label == "10-20" && bin.Count != 2 {
			t.Errorf("Histogram bin %s Actual: %d  Expected: %d", bin.Label, bin.Count, 2)
		} else if bin.Label == "20-30" && bin.Count != 0 {
			t.Errorf("Histogram bin %s Actual: %d  Expected: %d", bin.Label, bin.Count, 0)
		} else if bin.Label == "30-40" && bin.Count != 0 {
			t.Errorf("Histogram bin %s Actual: %d  Expected: %d", bin.Label, bin.Count, 0)
		} else if bin.Label == "40-50" && bin.Count != 1 {
			t.Errorf("Histogram bin %s Actual: %d  Expected: %d", bin.Label, bin.Count, 1)
		}
	}
}

func TestMysqlRepositoryConnectError(t *testing.T) {
	cfg := TestingConfig()
	repo := NewMysqlRepository(cfg)
	// Drop tables. All queries will fail
	repo.Flush()
	defer repo.Flush()

	// Test Store() Fail
	for _, c := range cases {
		err := repo.Store(New(c.query, c.status, c.time))

		if err == nil {
			t.Errorf("Query %q log shouldn't be store", c.query)
		}
	}

	// Test FindAll(start, end) Fail
	logs, err := repo.FindAll("", "")
	if err == nil || len(logs) != 0 {
		t.Errorf("Repo shouldn't get all logs")
	}

	// Test CountByStatus(start, end) Fail
	counts, err := repo.CountByStatus("", "")
	if err == nil || len(counts) != 0 {
		t.Errorf("Repo shouldn't get count by status")
	}

	// Test HistogramBins(start, end) Fail
	bins, err := repo.HistogramBins("", "")
	if err == nil || len(bins) != 0 {
		t.Errorf("Repo shouldn't get histogram bins")
	}
}

func TestMysqlRepositoryParseError(t *testing.T) {
	cfg := TestingConfig()
	repo := NewMysqlRepository(cfg)
	Rollback(repo.db)
	Migrate(repo.db)
	defer repo.Flush()

	// Test Store()
	for _, c := range cases {
		err := repo.Store(New(c.query, c.status, c.time))

		if err != nil {
			t.Errorf("Query %q log can't store: %s", c.query, err.Error())
		}
	}

	// Drop column
	_, err := repo.db.Exec("ALTER TABLE logs MODIFY status VARCHAR(255);")
	_, err = repo.db.Exec("ALTER TABLE logs MODIFY response_time VARCHAR(255);")
	_, err = repo.db.Exec("ALTER TABLE logs MODIFY created_at VARCHAR(255);")

	// Test FindAll(start, end)
	logs, err := repo.FindAll("", "")
	if err == nil || len(logs) != 0 {
		t.Errorf("Repo shouldn't get all logs")
	}

}
