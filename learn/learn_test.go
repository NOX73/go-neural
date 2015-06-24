package learn

import (
	"github.com/NOX73/go-neural"
	. "launchpad.net/gocheck"
	"math/rand"
	"testing"
)

type lessThenChecker struct {
	*CheckerInfo
}

var LessThen Checker = &lessThenChecker{
	&CheckerInfo{Name: "LessThen", Params: []string{"a", "b"}},
}

func (checker *lessThenChecker) Check(params []interface{}, names []string) (result bool, error string) {
	a, _ := params[0].(float64)
	b, _ := params[1].(float64)

	return a < b, ""
}

type moreThenChecker struct {
	*CheckerInfo
}

var MoreThen Checker = &moreThenChecker{
	&CheckerInfo{Name: "MoreThen", Params: []string{"a", "b"}},
}

func (checker *moreThenChecker) Check(params []interface{}, names []string) (result bool, error string) {
	a, _ := params[0].(float64)
	b, _ := params[1].(float64)

	return a > b, ""
}

func Test(t *testing.T) { TestingT(t) }

type SuiteT struct{}

func (s *SuiteT) SetUpTest(c *C) {}

var _ = Suite(&SuiteT{})

func (s *SuiteT) TestLearn(c *C) {

	n := neural.NewNetwork(2, []int{5, 5, 2})
	n.RandomizeSynapses()

	for i := 0; i < 1000000; i++ {
		in := []float64{rand.Float64() * 10, rand.Float64() * 10}
		var ideal []float64
		if in[0] > in[1] {
			ideal = []float64{1, 0}
		} else {
			ideal = []float64{0, 1}
		}
		Learn(n, in, ideal, 0.1)
	}

	count := 1000
	success := 0
	for i := 0; i < count; i++ {
		in := []float64{rand.Float64() * 10, rand.Float64() * 10}
		res := n.Calculate(in)

		if in[0] > in[1] && res[0] > res[1] {
			success++
		}
		if in[0] < in[1] && res[0] < res[1] {
			success++
		}
	}

	// Success persente should be more then 90%
	successPersents := float64(success) / float64(count)
	c.Assert(successPersents, MoreThen, 0.99)

}
