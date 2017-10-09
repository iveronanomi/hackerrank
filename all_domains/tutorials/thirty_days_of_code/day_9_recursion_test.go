package thirty_days_of_code

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day9RecursionSuite struct{}

var _ = Suite(&Day9RecursionSuite{})

func (s *Day9RecursionSuite) TestCase1(c *C) {
	out := utils.Capture(Recursion, "3")

	c.Assert(out, Equals, "6")
}
