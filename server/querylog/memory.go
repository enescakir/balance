package querylog

import (
	"log"
	"sync"
	"time"
)

// MemoryRepository implements Repository QueryLogy interface for MySQL.
type MemoryRepository struct {
	data  []QueryLog
	mutex *sync.Mutex
}

// NewMemoryRepository returns newly created MemoryRepository reference with given database.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{data: make([]QueryLog, 0), mutex:
	&sync.Mutex{}}
}

// Flush removes all data from in memory database.
func (r *MemoryRepository) Flush() {
	r.data = make([]QueryLog, 0)
}

// Store saves given QueryLog to in memory database.
func (r *MemoryRepository) Store(l *QueryLog) error {
	r.mutex.Lock()
	l.Id = len(r.data) + 1
	l.CreatedAt = time.Now()
	r.data = append(r.data, *l)
	defer r.mutex.Unlock()

	return nil
}

// FindAll returns all logs for given data range from in memory database.
func (r *MemoryRepository) FindAll(start string, end string) ([]*QueryLog, error) {
	logs, err := r.logsByDateRange(start, end)

	if err != nil {
		log.Printf("FindAll date parsing error: %s", err.Error())
		return nil, err
	}

	return logs, nil
}

// CountByStatus returns status:count pairs for given data range from in memory database.
func (r *MemoryRepository) CountByStatus(start string, end string) ([]*StatusCount, error) {
	logs, err := r.logsByDateRange(start, end)

	if err != nil {
		log.Printf("CountByStatus date parsing error: %s", err.Error())
		return nil, err
	}

	pairs := make(map[Status]int, 0)
	for _, l := range logs {
		if val, ok := pairs[l.Status]; ok {
			pairs[l.Status] = val + 1
		} else {
			pairs[l.Status] = 1
		}
	}

	counts := make([]*StatusCount, 0)
	for k, v := range pairs {
		counts = append(counts, &StatusCount{k, v})
	}

	return counts, nil
}

// HistogramBins returns responseTime:count bins for given data range from in memory database.
func (r *MemoryRepository) HistogramBins(start string, end string) ([]*HistogramBin, error) {
	logs, err := r.logsByDateRange(start, end)

	if err != nil {
		log.Printf("HistogramBins date parsing error: %s", err.Error())
		return nil, err
	}

	return createHistogramBins(logs), nil
}

// logsByDateRange returns logs in given start and end date.
func (r *MemoryRepository) logsByDateRange(start string, end string) ([]*QueryLog, error) {
	sDate, err := parseTimeWithLocation(start)
	if err != nil {
		return nil, err
	}
	eDate, err := parseTimeWithLocation(end)
	if err != nil {
		return nil, err
	}
	return r.createLogsByDateRange(sDate, eDate), nil
}

// createLogsByDateRange returns logs in rage of given Time references
func (r *MemoryRepository) createLogsByDateRange(sDate *time.Time, eDate *time.Time) []*QueryLog {
	logs := make([]*QueryLog, 0)
	for i := 0; i < len(r.data); i++ {
		l := &r.data[i]
		s := !(sDate != nil && l.CreatedAt.Before(*sDate))
		e := !(eDate != nil && l.CreatedAt.After(*eDate))

		if s && e {
			logs = append(logs, l)
		}

	}
	return logs
}

// parseTimeWithLocation converts given string to Time
func parseTimeWithLocation(timeStr string) (*time.Time, error) {
	if timeStr == "" {
		return nil, nil
	}
	date, err := time.ParseInLocation(TimeFormat, timeStr, defaultLoc)
	if err != nil {
		log.Printf("Date convert error: %v", err)
		return nil, err
	}
	return &date, nil
}
