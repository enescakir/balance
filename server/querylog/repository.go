package querylog

import (
	"time"
)

// TimeFormat is format for MySQL dates
const TimeFormat = "2006-01-02 15:04:05"

// defaultLoc is default timezone for date operations
var defaultLoc, _ = time.LoadLocation("Europe/Istanbul")

// Repository is interface for log data storage. It's used for dependency injection.
type Repository interface {
	Flush()
	Store(l *QueryLog) error
	FindAll(start string, end string) ([]*QueryLog, error)
	CountByStatus(start string, end string) ([]*StatusCount, error)
	HistogramBins(start string, end string) ([]*HistogramBin, error)
}
