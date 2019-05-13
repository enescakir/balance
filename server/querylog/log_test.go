package querylog

import "testing"

func TestQueryLog(t *testing.T) {
	cases := []struct {
		query  string
		status Status
		time   int64
	}{
		{"[][][]()", Balanced, 12634642},
		{"[][|][]()", Invalid, 5343634},
		{"[][][[]()", Unbalanced, 6234643},
	}

	for _, c := range cases {
		l := New(c.query, c.status, c.time)

		if c.query != l.Query {
			t.Errorf("Query log query Actual: %v  Expected: %v", l.Query, c.query)
		}
		if c.status != l.Status {
			t.Errorf("Query log status Actual: %v  Expected: %v", l.Status, c.status)
		}
		if c.time != l.ResponseTime {
			t.Errorf("Query log response time Actual: %v  Expected: %v", l.ResponseTime, c.time)
		}
	}
}
