package go_neural

func NewNetwork(in int, layers []int) *Network {
  n := &Network{
    make([]*Enter, 0, in),
    make([]*Layer, 0, len(layers)),
  }
  n.init(in, layers, NewLogisticFunc(1))
  return n
}

type Network struct {
  Enters      []*Enter
  Layers      []*Layer
}

func ( n *Network ) init (in int, layers []int, aFunc ActivationFunction) {
  n.initLayers(layers)
  n.initEnters(in)
  n.connectLayers()
  n.connectEnters()
  n.SetActivationFunction(aFunc)
}

func ( n *Network ) initLayers (layers []int) {
  for _, count := range layers {
    layer := NewLayer( count )
    n.Layers = append(n.Layers, layer)
  }
}

func ( n *Network ) initEnters (in int) {
  for ;in > 0; in-- {
    e := NewEnter()
    n.Enters = append(n.Enters, e)
  }
}

func ( n *Network ) connectLayers () {
  for i := len(n.Layers) - 1;i > 0;i-- {
    n.Layers[i-1].ConnectTo( n.Layers[i] )
  }
}

func ( n *Network ) connectEnters () {
  for _, e := range n.Enters {
    e.ConnectTo( n.Layers[0] )
  }
}

func ( n *Network ) EachLayer (f func(*Layer)) {
  for _, l := range n.Layers { f(l) }
}

func ( n *Network ) SetActivationFunction ( aFunc ActivationFunction ) {
  n.EachLayer(func (l *Layer){
    l.EachNeuron(func(n *Neuron){
      n.SetActivationFunction(aFunc)
    })
  })
}
