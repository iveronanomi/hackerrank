package thirty_days_of_code

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day3IntroTtoConditionalStatementsSuite struct{}

var _ = Suite(&Day3IntroTtoConditionalStatementsSuite{})

func (s *Day3IntroTtoConditionalStatementsSuite) TestCase1(c *C) {
	out := utils.Capture(Day3IntroTtoConditionalStatements, "3")

	c.Assert(out, Equals, "Weird")
}

func (s *Day3IntroTtoConditionalStatementsSuite) TestCase2(c *C) {
	out := utils.Capture(Day3IntroTtoConditionalStatements, "24")

	c.Assert(out, Equals, "Not Weird")
}
