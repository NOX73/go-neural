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

func DumpFromFile(path string) *NetworkDump {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	dump := &NetworkDump{}
	err = json.Unmarshal(b, dump)
	if err != nil {
		panic(err)
	}

	return dump
}

func FromFile(path string) *neural.Network {
	dump := DumpFromFile(path)
	n := FromDump(dump)
	return n
}

func ToFile(path string, n *neural.Network) {
	dump := ToDump(n)

	DumpToFile(path, dump)
}

func DumpToFile(path string, dump *NetworkDump) {
	j, _ := json.Marshal(dump)

	err := ioutil.WriteFile(path, j, 0644)
	if err != nil {
		panic(err)
	}
}

func ToDump(n *neural.Network) *NetworkDump {

	dump := &NetworkDump{Enters: len(n.Enters), Weights: make([][][]float64, len(n.Layers))}

	for i, l := range n.Layers {
		dump.Weights[i] = make([][]float64, len(l.Neurons))
		for j, n := range l.Neurons {
			dump.Weights[i][j] = make([]float64, len(n.InSynapses))
			for k, s := range n.InSynapses {
				dump.Weights[i][j][k] = s.Weight
			}
		}
	}

	return dump
}

func FromDump(dump *NetworkDump) *neural.Network {
	layers := make([]int, len(dump.Weights))
	for i, layer := range dump.Weights {
		layers[i] = len(layer)
	}

	n := neural.NewNetwork(dump.Enters, layers)

	for i, l := range n.Layers {
		for j, n := range l.Neurons {
			for k, s := range n.InSynapses {
				s.Weight = dump.Weights[i][j][k]
			}
		}
	}

	return n
}
