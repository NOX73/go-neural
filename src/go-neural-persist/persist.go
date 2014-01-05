package persist

import (
  "go-neural"
  "io/ioutil"
  "encoding/json"
)

type Weights [][][]float64
type Network struct {
  Enters        int
  Weights   Weights
}

func FromFile ( path string ) *neural.Network {
  b, err := ioutil.ReadFile(path)
  if err != nil { panic(err) }

  nDump := &Network{}
  err = json.Unmarshal(b, nDump)
  if err != nil { panic(err) }

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

func ToFile ( path string, n *neural.Network ) {

  nDump := Network{Enters: len(n.Enters), Weights: make([][][]float64, len(n.Layers))}

  for i, l := range n.Layers {
    nDump.Weights[i] = make([][]float64, len(l.Neurons))
    for j, n := range l.Neurons {
      nDump.Weights[i][j] = make([]float64, len(n.InSynapses))
      for k, s := range n.InSynapses {
        nDump.Weights[i][j][k] = s.Weight
      }
    }
  }

  j, _ := json.Marshal(nDump)

  err := ioutil.WriteFile(path, j, 0644)
  if err != nil { panic(err) }
}
