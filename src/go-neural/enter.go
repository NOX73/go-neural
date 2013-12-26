package neural

type Enter struct {
  Synapses    []*Synapse
  Input       float64 `json:"-"`
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
    e.SynapseTo(n, 0)
  }
}

func ( e *Enter ) sendSignal () {
  for _, s := range e.Synapses {
    s.sendSignal(e.Input)
  }
}
