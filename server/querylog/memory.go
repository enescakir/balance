package querylog

import (
	"fmt"
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
	return &MemoryRepository{data: []QueryLog{}, mutex:
	&sync.Mutex{}}
}

// Flush removes all data from in memory database.
func (r *MemoryRepository) Flush() {
	r.data = []QueryLog{}
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
	logs, err := r.getDateRangeLogs(start, end)

	if err != nil {
		log.Printf("FindAll date parsing error: %s", err.Error())
		return nil, err
	}

	return logs, nil
}

// GetCountByStatus returns status:count pairs for given data range from in memory database.
func (r *MemoryRepository) GetCountByStatus(start string, end string) ([]*StatusCount, error) {
	logs, err := r.getDateRangeLogs(start, end)

	if err != nil {
		log.Printf("GetCountByStatus date parsing error: %s", err.Error())
		return nil, err
	}

	pairs := map[Status]int{}
	for _, l := range logs {
		if val, ok := pairs[l.Status]; ok {
			pairs[l.Status] = val + 1
		} else {
			pairs[l.Status] = 1
		}
	}
	counts := []*StatusCount{}
	for k, v := range pairs {
		counts = append(counts, &StatusCount{k, v})
	}

	return counts, nil
}

// GetHistogramBins returns responseTime:count bins for given data range from in memory database.
func (r *MemoryRepository) GetHistogramBins(start string, end string) ([]*HistogramBin, error) {
	logs, err := r.getDateRangeLogs(start, end)

	if err != nil {
		log.Printf("GetHistogramBins date parsing error: %s", err.Error())
		return nil, err
	}

	buckets := map[string]int{}
	max := int64(0)
	for _, l := range logs {
		if l.ResponseTime > max {
			max = l.ResponseTime
		}
		left := (l.ResponseTime / 10000) * 10000
		lbl := fmt.Sprintf("%d-%d", left/1000, (left+10000)/1000)
		if val, ok := buckets[lbl]; ok {
			buckets[lbl] = val + 1
		} else {
			buckets[lbl] = 1
		}
	}

	bins := []*HistogramBin{}
	for i := int64(0); i < max; i += 10000 {
		lbl := fmt.Sprintf("%d-%d", i/1000, (i+10000)/1000)
		if val, ok := buckets[lbl]; ok {
			bins = append(bins, &HistogramBin{lbl, val})
		} else {
			bins = append(bins, &HistogramBin{lbl, 0})
		}
	}

	return bins, nil
}

// getDateRangeLogs returns logs in given start and end date.
func (r *MemoryRepository) getDateRangeLogs(start string, end string) ([]*QueryLog, error) {
	logs := []*QueryLog{}
	var sDate *time.Time
	var eDate *time.Time
	loc, _ := time.LoadLocation("Europe/Istanbul")

	if start != "" {
		date, err := time.ParseInLocation("2006-01-02 15:04:05", start, loc)
		sDate = &date

		if err != nil {
			log.Printf("Start date convert error: %v", err.Error())
			return nil, err
		}
	}

	if end != "" {
		date, err := time.ParseInLocation("2006-01-02 15:04:05", end, loc)
		eDate = &date

		if err != nil {
			log.Printf("End date convert error: %v", err.Error())
			return nil, err
		}
	}

	for i := 0; i < len(r.data); i++ {
		l := &r.data[i]
		s := !(sDate != nil && l.CreatedAt.Before(*sDate))
		e := !(eDate != nil && l.CreatedAt.After(*eDate))

		if s && e {
			logs = append(logs, l)
		}

	}
	return logs, nil
}
