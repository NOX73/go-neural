package neural

import (
  . "launchpad.net/gocheck"
)

func ( s *SuiteT ) TestAttachNeurons (c *C) {
  n := NewNeuron()
  n2 := NewNeuron()
  w := 0.5

  n.SynapseTo(n2, w)

  c.Assert( n.OutSynapses[0].Weight, Equals, w )
}

func ( s *SuiteT ) TestInputsSynapses (c *C) {
  n := NewNeuron()

  NewSynapseFromTo(NewNeuron(), n, 0.1)
  NewSynapseFromTo(NewNeuron(), n, 0.1)
  NewSynapseFromTo(NewNeuron(), n, 0.1)

  c.Assert( n.InSynapses, HasLen, 3 )
}
