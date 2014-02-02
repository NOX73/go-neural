package neural

func NewSynapse(weight float64) *Synapse {
	return &Synapse{Weight: weight}
}

func NewSynapseFromTo(from, to *Neuron, weight float64) *Synapse {
	syn := NewSynapse(weight)

	from.OutSynapses = append(from.OutSynapses, syn)
	to.InSynapses = append(to.InSynapses, syn)

	return syn
}

type Synapse struct {
	Weight float64
	In     float64 `json:"-"`
	Out    float64 `json:"-"`
}

func (s *Synapse) Signal(value float64) {
	s.In = value
	s.Out = s.In * s.Weight
}
