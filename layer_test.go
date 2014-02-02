package neural

import (
	. "launchpad.net/gocheck"
)

func (s *SuiteT) TestLayer(c *C) {
	l := NewLayer(5)
	c.Assert(l.Neurons, HasLen, 5)
}

func (s *SuiteT) TestConnectToLayer(c *C) {
	count := 5
	l := NewLayer(count)
	l2 := NewLayer(count)

	l.ConnectTo(l2)

	for _, n := range l.Neurons {
		c.Assert(n.OutSynapses, HasLen, count)
	}

}
