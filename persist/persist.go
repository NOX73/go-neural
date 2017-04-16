package persist

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/flezzfx/gopher-neural"
)

type Weights [][][]float64
type NetworkDump struct {
	Enters    int
	Weights   Weights
	OutLabels map[string]string
}

func DumpFromFile(path string) (*NetworkDump, error) {
	b, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, err
	}
	dump := &NetworkDump{}
	err = json.Unmarshal(b, dump)
	if nil != err {
		return nil, err
	}

	return dump, nil
}

func FromFile(path string) (*neural.Network, error) {
	dump, err := DumpFromFile(path)
	if nil != err {
		return nil, err
	}
	n := FromDump(dump)
	return n, nil
}

func ToFile(path string, n *neural.Network) error {
	dump := ToDump(n)
	return DumpToFile(path, dump)
}

func DumpToFile(path string, dump *NetworkDump) error {
	j, err := json.Marshal(dump)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, j, 0644)
	return err
}

func ToDump(n *neural.Network) *NetworkDump {
	labels := intToStringMap(n.OutLabels)
	dump := &NetworkDump{Enters: len(n.Enters), Weights: make([][][]float64, len(n.Layers)), OutLabels: labels}

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
	labels := stringToIntMap(dump.OutLabels)
	n := neural.NewNetwork(dump.Enters, layers, labels)

	for i, l := range n.Layers {
		for j, n := range l.Neurons {
			for k, s := range n.InSynapses {
				s.Weight = dump.Weights[i][j][k]
			}
		}
	}

	return n
}

func intToStringMap(m map[int]string) map[string]string {
	ms := make(map[string]string)
	for k, v := range m {
		ms[strconv.Itoa(k)] = v
	}
	return ms
}

func stringToIntMap(m map[string]string) map[int]string {
	mi := make(map[int]string)
	for k, v := range m {
		index, err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		mi[index] = v
	}
	return mi
}
