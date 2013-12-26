package persist

import ( 
  "go-neural"
  "io/ioutil"
  "encoding/json"
  //"log"
)

func FromFile ( path string ) *neural.Network {
  b, err := ioutil.ReadFile(path)
  if err != nil { panic(err) }

  n := &neural.Network{}
  err = json.Unmarshal(b, n)
  if err != nil { panic(err) }

  l := n.Layers[0]
  for _, e := range n.Enters {
    for i, s := range e.Synapses {
      s.Neuron = l.Neurons[i]
    }
  }

  for i, l := range n.Layers {
    if i+1 == len(n.Layers) {break}
    l2 := n.Layers[i+1]
    for _, n := range l.Neurons {
      for j, s := range n.Synapses {
        s.Neuron = l2.Neurons[j]
      }
    }
  }

  n.SetActivationFunction(neural.NewLogisticFunc(1))

  return n
}

func ToFile ( path string, n *neural.Network ) {
  j, _ := json.Marshal(n)

  err := ioutil.WriteFile(path, j, 0644)
  if err != nil { panic(err) }
}
