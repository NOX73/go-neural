package neural

type Enter struct {
	OutSynapses []*Synapse
	Input       float64 `json:"-"`
}

func NewEnter() *Enter {
	return &Enter{}
}

func (e *Enter) SynapseTo(nTo *Neuron, weight float64) {
	syn := NewSynapse(weight)

	e.OutSynapses = append(e.OutSynapses, syn)
	nTo.InSynapses = append(nTo.InSynapses, syn)
}

func (e *Enter) SetInput(val float64) {
	e.Input = val
}

func (e *Enter) ConnectTo(layer *Layer) {
	for _, n := range layer.Neurons {
		e.SynapseTo(n, 0)
	}
}

func (e *Enter) Signal() {
	for _, s := range e.OutSynapses {
		s.Signal(e.Input)
	}
}
