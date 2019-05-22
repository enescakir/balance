package querylog

import "fmt"

const (
	BinWidth    = 10000 // Histogram bin width in nanoseconds
	NanoToMicro = 1000  // Constant for converting nanoseconds to microseconds
)

// HistogramBin represents responseTime:count bins.
type HistogramBin struct {
	Label string `json:"label"`
	Count int    `json:"count"`
}

// createHistogramBins creates histogram bins from given logs
func createHistogramBins(logs []*QueryLog) []*HistogramBin {
	buckets := make(map[string]int, 0)
	max := int64(0)
	for _, l := range logs {
		if l.ResponseTime > max {
			max = l.ResponseTime
		}
		left := (l.ResponseTime / BinWidth) * BinWidth
		lbl := fmt.Sprintf("%d-%d", left/NanoToMicro, (left+BinWidth)/NanoToMicro)
		buckets[lbl]++
	}

	bins := make([]*HistogramBin, 0)
	for i := int64(0); i < max; i += BinWidth {
		lbl := fmt.Sprintf("%d-%d", i/NanoToMicro, (i+BinWidth)/NanoToMicro)
		bins = append(bins, &HistogramBin{lbl, buckets[lbl]})
	}

	return bins
}
