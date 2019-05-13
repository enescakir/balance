package querylog

import (
	"testing"
	"time"
)

var mCases = []struct {
	query  string
	status Status
	time   int64
}{
	{"[][][]()", Balanced, 4500},
	{"[][][][]()", Balanced, 10000},
	{"[][|][]()", Invalid, 10050},
	{"[][][[]()", Unbalanced, 40005},
}

func TestMemoryRepository(t *testing.T) {
	repo := NewMemoryRepository()
	defer repo.Flush()

	before := time.Now().Add(-time.Hour).Format("2006-01-02 15:04:05")
	after := time.Now().Add(time.Hour).Format("2006-01-02 15:04:05")

	// Test Store()
	for _, c := range mCases {
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

	if len(logs) != len(mCases) {
		t.Errorf("Find all count not matched  Actual: %d  Expected: %d", len(logs), len(mCases))
	}

	logs, err = repo.FindAll(before, after)
	if err != nil {
		t.Errorf("Repo can't get all logs in date range: %s", err.Error())
	}

	if len(logs) != len(mCases) {
		t.Errorf("Find all count not matched in date range  Actual: %d  Expected: %d", len(logs), len(mCases))
	}

	logs, err = repo.FindAll(before, before)
	if err != nil {
		t.Errorf("Repo can't get all logs in date range: %s", err.Error())
	}

	if len(logs) != 0 {
		t.Errorf("Find all count not matched in date range  Actual: %d  Expected: %d", len(logs), 0)
	}

	// Test GetCountByStatus(start, end)
	counts, err := repo.GetCountByStatus("", "")
	for _, pair := range counts {
		if pair.Status == Balanced && pair.Count != 2 {
			t.Errorf("Balanced logs counts Actual: %d  Expected: %d", pair.Count, 2)
		} else if pair.Status == Unbalanced && pair.Count != 1 {
			t.Errorf("Unbalanced logs counts Actual: %d  Expected: %d", pair.Count, 1)
		} else if pair.Status == Invalid && pair.Count != 1 {
			t.Errorf("Invalid logs counts Actual: %d  Expected: %d", pair.Count, 1)
		}
	}

	// Test GetCountByStatus(start, end)
	bins, err := repo.GetHistogramBins("", "")
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

func TestMemoryRepositoryConnectError(t *testing.T) {
	repo := NewMemoryRepository()
	// Drop tables. All queries will fail
	repo.Flush()

	// Test FindAll(start, end) Fail
	logs, err := repo.FindAll("asf", "asdf")
	if err == nil || len(logs) != 0 {
		t.Errorf("Repo shouldn't get all logs")
	}

	// Test GetCountByStatus(start, end) Fail
	counts, err := repo.GetCountByStatus("", "asdf")
	if err == nil || len(counts) != 0 {
		t.Errorf("Repo shouldn't get count by status")
	}

	// Test GetHistogramBins(start, end) Fail
	bins, err := repo.GetHistogramBins("asdf", "asdf")
	if err == nil || len(bins) != 0 {
		t.Errorf("Repo shouldn't get histogram bins")
	}
}
