package neural

import "fmt"
import "math/rand"

func NewNetwork(in int, layers []int) *Network {
	n := &Network{
		Enters: make([]*Enter, 0, in),
		Layers: make([]*Layer, 0, len(layers)),
	}
	n.init(in, layers, NewLogisticFunc(1))
	return n
}

type Network struct {
	Enters []*Enter
	Layers []*Layer
	Out    []float64 `json:"-"`
}

func (n *Network) init(in int, layers []int, aFunc ActivationFunction) {
	n.initLayers(layers)
	n.initEnters(in)
	n.ConnectLayers()
	n.ConnectEnters()
	n.SetActivationFunction(aFunc)
}

func (n *Network) initLayers(layers []int) {
	for _, count := range layers {
		layer := NewLayer(count)
		n.Layers = append(n.Layers, layer)
	}
}

func (n *Network) initEnters(in int) {
	for ; in > 0; in-- {
		e := NewEnter()
		n.Enters = append(n.Enters, e)
	}
}

func (n *Network) ConnectLayers() {
	for i := len(n.Layers) - 1; i > 0; i-- {
		n.Layers[i-1].ConnectTo(n.Layers[i])
	}
}

func (n *Network) ConnectEnters() {
	for _, e := range n.Enters {
		e.ConnectTo(n.Layers[0])
	}
}

func (n *Network) SetActivationFunction(aFunc ActivationFunction) {
	for _, l := range n.Layers {
		for _, n := range l.Neurons {
			n.SetActivationFunction(aFunc)
		}
	}
}

func (n *Network) setEnters(v *[]float64) {
	values := *v
	if len(values) != len(n.Enters) {
		panic(fmt.Sprint("Enters count ( ", len(n.Enters), " ) != count of elements in SetEnters function argument ( ", len(values), " ) ."))
	}

	for i, e := range n.Enters {
		e.Input = values[i]
	}

}

func (n *Network) sendEnters() {
	for _, e := range n.Enters {
		e.Signal()
	}
}

func (n *Network) calculateLayers() {
	for _, l := range n.Layers {
		l.Calculate()
	}
}

func (n *Network) generateOut() {
	outL := n.Layers[len(n.Layers)-1]
	n.Out = make([]float64, len(outL.Neurons))

	for i, neuron := range outL.Neurons {
		n.Out[i] = neuron.Out
	}
}

func (n *Network) Calculate(enters []float64) []float64 {
	n.setEnters(&enters)
	n.sendEnters()
	n.calculateLayers()
	n.generateOut()

	return n.Out
}

func (n *Network) RandomizeSynapses() {
	for _, l := range n.Layers {
		for _, n := range l.Neurons {
			for _, s := range n.InSynapses {
				s.Weight = 2 * (rand.Float64() - 0.5)
			}
		}
	}
}
