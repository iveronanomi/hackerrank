package implementation

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type DivisibleSumPairsSuite struct{}

var _ = Suite(&DivisibleSumPairsSuite{})

func (s *DivisibleSumPairsSuite) TestCase1(c *C) {
	out := utils.Capture(DivisibleSumPairs, "6 3\n1 3 2 6 1 2")

	c.Assert(out, Equals, "5\n")
}
