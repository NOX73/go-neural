package neural

import (
	. "launchpad.net/gocheck"
	//"log"
)

func (s *SuiteT) TestCreateNewNetwork(c *C) {
	net := NewNetwork(10, []int{5, 6, 2})

	c.Assert(net.Enters, HasLen, 10)
	c.Assert(net.Enters[0].OutSynapses, HasLen, 5)

	c.Assert(net.Layers, HasLen, 3)
	c.Assert(net.Layers[0].Neurons, HasLen, 5)

	c.Assert(net.Layers[0].Neurons[0].OutSynapses, HasLen, 6)
	c.Assert(net.Layers[1].Neurons[0].OutSynapses, HasLen, 2)
	c.Assert(net.Layers[2].Neurons[0].OutSynapses, HasLen, 0)
}

func (s *SuiteT) TestCalculcateNetwork(c *C) {

	net := NewNetwork(5, []int{5, 6, 1})

	out := net.Calculate([]float64{0, 0, 0, 0, 0})

	c.Assert(out[0], Equals, 0.5)

}
