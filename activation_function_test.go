package go_neural

import (
  . "launchpad.net/gocheck"
)

func ( s *SuiteT ) TestLogisticFunc (c *C) {
  f := NewLogisticFunc(1)

  c.Assert(f(0), Equals, 0.5)
  //c.Assert(1 - f(6))
}


