package neural

type Neuron struct {
  Synapses    []*Synapse
  Inputs      []float64 `json:"-"`
  ActivationFunction  ActivationFunction `json:"-"`
  Out         float64 `json:"-"`
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

func ( n *Neuron ) SetActivationFunction ( aFunc ActivationFunction ) {
  n.ActivationFunction = aFunc
}

func ( n *Neuron ) Calculate () {
  var sum float64
  for _, i := range n.Inputs { sum += i }

  n.Out = n.ActivationFunction(sum)

  for _, s := range n.Synapses {
    s.sendSignal(n.Out)
  }
}
