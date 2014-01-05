package lern

import (
  "github.com/NOX73/go-neural"
  "math"
)

// Math Evaluation with Least squares method.
func Evaluation (n *neural.Network, in, ideal []float64 ) float64 {
  out := n.Calculate(in)

  var e float64
  for i, _ := range out {
    e += math.Pow(out[i] - ideal[i], 2)
  }

  return e / 2
}


