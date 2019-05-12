package querylog

type Repository interface {
	Store(l *QueryLog) error
	FindAll(start string, end string) ([]*QueryLog, error)
	GetCountByStatus(start string, end string) ([]*StatusCount, error)
	GetHistogramBins(start string, end string) ([]*HistogramBin, error)
}
