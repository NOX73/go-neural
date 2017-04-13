package evaluation

import (
	"math"

	neural "github.com/NOX73/go-neural"
)

const (
	TruePositive  = 0
	TrueNegative  = 1
	FalsePositive = 2
	FalseNegative = 3
)

type Evaluation struct {
	TruePositive  int
	TrueNegative  int
	FalsePositive int
	FalseNegative int
	Confusion     map[string]map[string]int
}

// TODO (abresk) add

// TODO (abresk) sum from all metrics

// TODO (abresk) getAccuracy

// TODO (abresk) getF1

// TODO (abresk) some other metrics

// TODO (abresk) confusion matrix handling

// Math Evaluation with Least squares method.
func shortEvaluation(n *neural.Network, in, ideal []float64) float64 {
	// This function was part of the former go-neural and moved to this package.
	out := n.Calculate(in)
	var e float64
	for i, _ := range out {
		e += math.Pow(out[i]-ideal[i], 2)
	}

	return e / 2
}
