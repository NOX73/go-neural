package neural

func NewLayer(neurons int) *Layer {
	l := &Layer{}
	l.init(neurons)
	return l
}

type Layer struct {
	Neurons []*Neuron
}

func (l *Layer) ConnectTo(layer *Layer) {
	for _, n := range l.Neurons {
		for _, toN := range layer.Neurons {
			n.SynapseTo(toN, 0)
		}
	}
}

func (l *Layer) init(neurons int) {
	for ; neurons > 0; neurons-- {
		l.addNeuron()
	}
}

func (l *Layer) addNeuron() {
	n := NewNeuron()
	l.Neurons = append(l.Neurons, n)
}

func (l *Layer) Calculate() {
	for _, n := range l.Neurons {
		n.Calculate()
	}
}
