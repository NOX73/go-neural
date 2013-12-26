package neural

import (
  . "launchpad.net/gocheck"
)

func ( s *SuiteT ) TestAttachNeurons (c *C) {
  n := NewNeuron()
  n2 := NewNeuron()
  w := 0.5

  n.SynapseTo(n2, w)

  c.Assert( n.Synapses[0].Weight, Equals, w )
}

func ( s *SuiteT ) TestInputs (c *C) {
  n := NewNeuron()

  n.AppendInput(0.1)
  n.AppendInput(0.1)
  n.AppendInput(0.1)

  c.Assert( n.Inputs, HasLen, 3 )

  n.ResetInputs()

  c.Assert( n.Inputs, HasLen, 0 )
}
