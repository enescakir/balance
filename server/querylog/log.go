package querylog

import "time"

// QueryLog keeps record of the request to balance endpoint.
type QueryLog struct {
	Id           int       `json:"id"`
	Query        string    `json:"query"`
	Status       Status    `json:"status"`
	ResponseTime int64     `json:"response_time"`
	CreatedAt    time.Time `json:"created_at"`
}

// New returns newly created QueryLog reference.
func New(query string, status Status, rTime int64) *QueryLog {
	return &QueryLog{Query: query, Status: status, ResponseTime: rTime}
}

// StatusCount represents status:count pair.
type StatusCount struct {
	Status Status `json:"status"`
	Count  int    `json:"count"`
}

// HistogramBin represents responseTime:count bins.
type HistogramBin struct {
	Label string `json:"label"`
	Count int    `json:"count"`
}

// Status represents result of the request in QueryLog.
type Status int

const (
	Unknown    Status = 0
	Balanced   Status = 1
	Unbalanced Status = 2
	Invalid    Status = 3
)
