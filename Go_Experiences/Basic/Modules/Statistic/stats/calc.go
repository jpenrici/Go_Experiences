// stats/calc.go

package stats // same as the directory name

import (
	"errors"
	"math"
)

// Exported
type Summary struct {
	Mean   float64
	StdDev float64
}

type DataProcessor interface {
	Process() (Summary, error)
}

type SimpleStats struct {
	Values []float64
}

// Process implements DataProcessor
func (s SimpleStats) Process() (Summary, error) {
	if len(s.Values) == 0 {
		return Summary{}, errors.New("empty data")
	}

	var sum float64
	for _, v := range s.Values {
		sum += v
	}
	mean := sum / float64(len(s.Values))

	var vSum float64
	for _, v := range s.Values {
		vSum += math.Pow(v-mean, 2)
	}
	sd := math.Sqrt(vSum / float64(len(s.Values)))

	return Summary{Mean: mean, StdDev: sd}, nil
}
