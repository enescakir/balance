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

// QueryLogs is a collection of QueryLog.
type QueryLogs []QueryLog

// Status represents result of the request in QueryLog.
type Status int

const (
	Unknown    Status = 0
	Balanced   Status = 1
	Unbalanced Status = 2
	Invalid    Status = 3
)

// NewQueryLog returns newly created QueryLog reference.
func NewQueryLog(query string, status Status, rTime int64) *QueryLog {
	return &QueryLog{Query: query, Status: status, ResponseTime: rTime}
}