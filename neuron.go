package go_neural

func NewNeuron () *Neuron {
  return &Neuron{}
}

type Neuron struct {
  Synapses    []*Synapse
  Inputs      []float64
}

func ( n *Neuron ) SynapseTo ( neuron *Neuron, weight float64 ) {
  syn := NewSynapse( neuron, weight )
  n.Synapses = append( n.Synapses, syn )
}

func ( n *Neuron ) ResetInputs () {
  n.Inputs = n.Inputs[:0]
}

func ( n *Neuron ) AppendInput ( val float64 ) {
  n.Inputs = append( n.Inputs, val)
}
