package persist

import ( 
  "go-neural"
  "io/ioutil"
  "encoding/json"
  //"log"
)

type MarchalNetwork struct {
  Enters int
  Layers []int
}

func FromFile ( path string ) *neural.Network {
  b, err := ioutil.ReadFile(path)
  if err != nil { panic(err) }

  h := MarchalNetwork{}
  err = json.Unmarshal(b, &h)
  if err != nil { panic(err) }

  n := neural.NewNetwork(h.Enters, h.Layers)

  return n
}

func ToFile ( path string, n *neural.Network ) {
  h := ToHash(n)

  j, _ := json.Marshal(h)

  err := ioutil.WriteFile(path, j, 0644)
  if err != nil { panic(err) }
}

type Hash map[string]interface{}

func ToHash ( n *neural.Network ) MarchalNetwork {
  h := MarchalNetwork{}

  h.Enters = len(n.Enters)
  for _, l := range n.Layers {
    h.Layers = append(h.Layers, len(l.Neurons))
  }

  return h
}
