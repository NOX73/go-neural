package go_neural

func NewNetwork(in, out, layers int) *Network {
  return &Network{}
}

type Network struct {
  Layers      []*Layer
}
