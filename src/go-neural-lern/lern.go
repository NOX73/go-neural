package lern

import (
  "go-neural"
  //"log"
)

type Deltas [][]float64

func Lern ( n *neural.Network, in, ideal []float64, speed float64 ) {
  Backpropagation(n, in, ideal, speed)
}

func Backpropagation ( n *neural.Network, in, ideal []float64, speed float64 ) {
  n.Calculate(in)

  deltas := make([][]float64, len(n.Layers))

  last := len(n.Layers) - 1
  l := n.Layers[last]
  deltas[last] = make([]float64, len(l.Neurons))
  for i, n := range l.Neurons {
    deltas[last][i] = n.Out * ( 1 - n.Out ) * ( ideal[i] - n.Out )
  }

  for i := last - 1; i >= 0; i-- {
    l := n.Layers[i]
    deltas[i] = make([]float64, len(l.Neurons))
    for j, n := range l.Neurons {

      var sum float64 = 0
      for k, s := range n.OutSynapses {
        sum += s.Weight * deltas[i + 1][k]
      }

      deltas[i][j] = n.Out * ( 1 - n.Out ) * sum
    }
  }

  for i, l := range n.Layers {
    for j, n := range l.Neurons {
      for _, s := range n.InSynapses {
        s.Weight += speed * deltas[i][j] * s.In
      }
    }
  }

}

