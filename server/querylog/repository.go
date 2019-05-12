package querylog

// Repository is interface for log data storage. It's used for dependency injection.
type Repository interface {
	Store(l *QueryLog) error
	FindAll(start string, end string) ([]*QueryLog, error)
	GetCountByStatus(start string, end string) ([]*StatusCount, error)
	GetHistogramBins(start string, end string) ([]*HistogramBin, error)
}
