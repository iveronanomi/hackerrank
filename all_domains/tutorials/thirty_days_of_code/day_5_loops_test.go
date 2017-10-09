package thirty_days_of_code

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day2LoopsSuite struct{}

var _ = Suite(&Day2LoopsSuite{})

func (s *Day2LoopsSuite) TestCase1(c *C) {
	out := utils.Capture(Day5Loops, "-5")

	c.Assert(out, Equals, "-5 x 1 = -5-5 x 2 = -10-5 x 3 = -15-5 x 4 = -20-5 x 5 = -25-5 x 6 = -30-5 x 7 = -35-5 x 8 = -40-5 x 9 = -45-5 x 10 = -50")
}

func (s *Day2LoopsSuite) TestCase2(c *C) {
	out := utils.Capture(Day5Loops, "5")

	c.Assert(out, Equals, "5 x 1 = 55 x 2 = 105 x 3 = 155 x 4 = 205 x 5 = 255 x 6 = 305 x 7 = 355 x 8 = 405 x 9 = 455 x 10 = 50")
}

func (s *Day2LoopsSuite) TestCase3(c *C) {
	out := utils.Capture(Day5Loops, "0")

	c.Assert(out, Equals, "0 x 1 = 00 x 2 = 00 x 3 = 00 x 4 = 00 x 5 = 00 x 6 = 00 x 7 = 00 x 8 = 00 x 9 = 00 x 10 = 0")
}
