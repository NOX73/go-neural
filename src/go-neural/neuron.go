package neural

type Neuron struct {
  OutSynapses           []*Synapse
  InSynapses            []*Synapse `json:"-"`
  ActivationFunction    ActivationFunction `json:"-"`
  Out                   float64 `json:"-"`
}

func NewNeuron () *Neuron {
  return &Neuron{}
}

func ( n *Neuron ) SynapseTo ( nTo *Neuron, weight float64 ) {
  NewSynapseFromTo(n, nTo, weight)
}

func ( n *Neuron ) SetActivationFunction ( aFunc ActivationFunction ) {
  n.ActivationFunction = aFunc
}

func ( n *Neuron ) Calculate () {
  var sum float64
  for _, s := range n.InSynapses { sum += s.Out }

  n.Out = n.ActivationFunction(sum)

  for _, s := range n.OutSynapses {
    s.Signal(n.Out)
  }
}
