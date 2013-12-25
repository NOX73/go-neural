package go_neural

import "math/rand"

type Enter struct {
  Synapses    []*Synapse
  Input       float64
}

func NewEnter () *Enter {
  return &Enter{}
}

func ( e *Enter ) SynapseTo ( nTo *Neuron, weight float64 ) {
  syn := NewSynapse( nTo, weight )
  e.Synapses = append( e.Synapses, syn )
}

func ( e *Enter ) SetInput ( val float64 ) {
  e.Input = val
}

func ( e *Enter ) ConnectTo ( layer *Layer ) {
  for _, n := range layer.Neurons {
    e.SynapseTo(n, rand.Float64()*0.1 )
  }
}
