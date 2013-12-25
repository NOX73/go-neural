package go_neural

import (
  . "launchpad.net/gocheck"
  //"log"
)

func ( s *SuiteT ) TestCreateNewNetwork (c *C) {
  net := NewNetwork(10, []int{5,6,2})

  c.Assert(len(net.Enters), Equals, 10)
  c.Assert(len(net.Enters[0].Synapses), Equals, 5)

  c.Assert(len(net.Layers), Equals, 3)
  c.Assert(len(net.Layers[0].Neurons), Equals, 5)

  c.Assert(len(net.Layers[0].Neurons[0].Synapses), Equals, 6)
  c.Assert(len(net.Layers[1].Neurons[0].Synapses), Equals, 2)
  c.Assert(len(net.Layers[2].Neurons[0].Synapses), Equals, 0)
}

