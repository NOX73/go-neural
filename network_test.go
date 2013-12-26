package go_neural

import (
  . "launchpad.net/gocheck"
  //"log"
)

func ( s *SuiteT ) TestCreateNewNetwork (c *C) {
  net := NewNetwork(10, []int{5,6,2})

  c.Assert(net.Enters, HasLen, 10)
  c.Assert(net.Enters[0].Synapses, HasLen, 5)

  c.Assert(net.Layers, HasLen, 3)
  c.Assert(net.Layers[0].Neurons, HasLen, 5)

  c.Assert(net.Layers[0].Neurons[0].Synapses, HasLen, 6)
  c.Assert(net.Layers[1].Neurons[0].Synapses, HasLen, 2)
  c.Assert(net.Layers[2].Neurons[0].Synapses, HasLen, 0)
}

