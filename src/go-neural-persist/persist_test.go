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
  path := "/tmp/network.json"
  ToFile(path, n)

  n2 := FromFile(path)

  c.Assert(len(n2.Enters), Equals, len(n.Enters))
  c.Assert(len(n2.Layers), Equals, len(n.Layers))
}

