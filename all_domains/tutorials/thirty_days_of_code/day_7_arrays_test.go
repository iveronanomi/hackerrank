package thirty_days_of_code

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day7ArraysSuite struct{}

var _ = Suite(&Day7ArraysSuite{})

func (s *Day7ArraysSuite) TestCase1(c *C) {
	out := utils.Capture(Day7Arrays, "4\n1 4 3 2")

	c.Assert(out, Equals, "2 3 4 1 ")
}
