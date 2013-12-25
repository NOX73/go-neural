package go_neural

type Neuron struct {
  Synapses    []*Synapse
  Inputs      []float64
}

func NewNeuron () *Neuron {
  return &Neuron{}
}

func ( n *Neuron ) SynapseTo ( nTo *Neuron, weight float64 ) {
  syn := NewSynapse( nTo, weight )
  n.Synapses = append( n.Synapses, syn )
}

func ( n *Neuron ) ResetInputs () {
  n.Inputs = n.Inputs[:0]
}

func ( n *Neuron ) AppendInput ( val float64 ) {
  n.Inputs = append( n.Inputs, val)
}

