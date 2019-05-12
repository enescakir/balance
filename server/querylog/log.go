package querylog

import (
	"time"
)

type QueryLog struct {
	Id           int       `json:"id"`
	Query        string    `json:"query"`
	Status       Status    `json:"Status"`
	ResponseTime int64     `json:"response_time"`
	CreatedAt    time.Time `json:"created_at"`
}

type QueryLogs []QueryLog

type Status int

const (
	Unknown    Status = 0
	Balanced   Status = 1
	Unbalanced Status = 2
	Invalid    Status = 3
)

func NewQueryLog(query string, status Status, rTime int64) *QueryLog {
	return &QueryLog{Query: query, Status: status, ResponseTime: rTime}
}
