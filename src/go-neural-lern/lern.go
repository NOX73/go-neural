package lern

import (
  "go-neural"
)

func Backpropagation ( n neural.Network, in, ideal []float64 ) {

  out := n.Calculate(in)

}
