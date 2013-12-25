package go_neural

import (
  "testing"
  . "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }
type SuiteT struct { }
func (s *SuiteT) SetUpTest (c *C) { }
var _ = Suite( &SuiteT{} )


