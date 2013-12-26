package persist

import (
  "testing"
  . "launchpad.net/gocheck"
  "go-neural"
)

func Test(t *testing.T) { TestingT(t) }
type SuiteT struct { }
func (s *SuiteT) SetUpTest (c *C) { }
var _ = Suite( &SuiteT{} )

func ( s *SuiteT ) TestPersist (c *C) {
  n := neural.NewNetwork(5, []int{5,5,5})
  n.RandomizeSynapses()

  path := "/tmp/network.json"
  ToFile(path, n)

  n2 := FromFile(path)

  c.Assert(n2.Enters, HasLen, len(n.Enters))
  c.Assert(n2.Layers, HasLen, len(n.Layers))

  for i, l := range n2.Layers {
    for j, nr := range l.Neurons {
      for h, s := range nr.Synapses {
        c.Assert(s.Weight, Equals, n.Layers[i].Neurons[j].Synapses[h].Weight)
      }
    }
  }

  in := []float64{0.5, 0.5, 0.5, 0.5, 0.5}
  out := n.Calculate(in)
  out2 := n2.Calculate(in)

  for i, o := range out2 {
    c.Assert(o, Equals, out[i])
  }

}

