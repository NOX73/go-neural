package persist

import (
	"encoding/json"
	"github.com/NOX73/go-neural"
	"io/ioutil"
)

type Weights [][][]float64
type NetworkDump struct {
	Enters  int
	Weights Weights
}

func FromFile(path string) *neural.Network {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	nDump := &NetworkDump{}
	err = json.Unmarshal(b, nDump)
	if err != nil {
		panic(err)
	}

	n := FromDump(nDump)
	return n
}

func ToFile(path string, n *neural.Network) {
	nDump := ToDump(n)
	j, _ := json.Marshal(nDump)

	err := ioutil.WriteFile(path, j, 0644)
	if err != nil {
		panic(err)
	}
}

func ToDump(n *neural.Network) *NetworkDump {

	nDump := &NetworkDump{Enters: len(n.Enters), Weights: make([][][]float64, len(n.Layers))}

	for i, l := range n.Layers {
		nDump.Weights[i] = make([][]float64, len(l.Neurons))
		for j, n := range l.Neurons {
			nDump.Weights[i][j] = make([]float64, len(n.InSynapses))
			for k, s := range n.InSynapses {
				nDump.Weights[i][j][k] = s.Weight
			}
		}
	}

	return nDump
}

func FromDump(nDump *NetworkDump) *neural.Network {
	layers := make([]int, len(nDump.Weights))
	for i, layer := range nDump.Weights {
		layers[i] = len(layer)
	}

	n := neural.NewNetwork(nDump.Enters, layers)

	for i, l := range n.Layers {
		for j, n := range l.Neurons {
			for k, s := range n.InSynapses {
				s.Weight = nDump.Weights[i][j][k]
			}
		}
	}

	return n
}
