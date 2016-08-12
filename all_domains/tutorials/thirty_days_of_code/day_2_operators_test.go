package thirty_days_of_code

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day2OperatorsSuite struct{}

var _ = Suite(&Day2OperatorsSuite{})

func (s *Day2OperatorsSuite) TestCase1(c *C) {
	out := utils.CaptureOut(Day2Operators, "12.00\n20\n8")

	c.Assert(out, Equals, "The total meal cost is 15 dollars.")
}

func (s *Day2OperatorsSuite) TestCase2(c *C) {
	out := utils.CaptureOut(Day2Operators, "15.50\n15\n10")

	c.Assert(out, Equals, "The total meal cost is 19 dollars.")
}
