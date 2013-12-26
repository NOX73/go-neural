package go_neural

func NewSynapse (neuron *Neuron, weight float64) *Synapse {
  return &Synapse{ neuron, weight }
}

type Synapse struct {
  Neuron    *Neuron
  Weight    float64
}

func ( s *Synapse ) sendSignal (value float64) {
  s.Neuron.AppendInput(value * s.Weight)
}

