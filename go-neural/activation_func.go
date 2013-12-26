package go_neural

import "math"

type ActivationFunction func(float64)float64;

func NewLogisticFunc(a float64) ActivationFunction {
  return func (x float64) float64 {
    return LogisticFunc(x, a)
  }
}
func LogisticFunc(x, a float64) float64 {
  return 1 / ( 1 + math.Exp(- a * x) )
}
