package b

import (
	"math"

	"github.com/cybersamx/c"
)

func Variance(sample []float64) float64 {
	sum := 0.0
	mean := c.Mean(sample)

	for _, v := range sample {
		sum += math.Pow(v-mean, 2)
	}

	variance := sum / c.Count(sample)

	return variance
}
