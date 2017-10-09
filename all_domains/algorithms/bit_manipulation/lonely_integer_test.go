package bit_manipulation

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type LonelyIntegerSuite struct{}

var _ = Suite(&LonelyIntegerSuite{})

func (s *LonelyIntegerSuite) TestLonelyIntegerCase1(c *C) {
	out := utils.Capture(LonelyInteger, "3\n1 1 2")

	c.Assert(out, Equals, "2\n")
}

func (s *LonelyIntegerSuite) TestLonelyIntegerCase2(c *C) {
	out := utils.Capture(LonelyInteger, "5\n0 0 1 2 1")

	c.Assert(out, Equals, "2\n")
}

func (s *LonelyIntegerSuite) TestLonelyIntegerCase3(c *C) {
	out := utils.Capture(LonelyInteger, "1\n1")

	c.Assert(out, Equals, "1\n")
}

func (s *LonelyIntegerSuite) TestLonelyIntegerCase4(c *C) {
	out := utils.Capture(LonelyInteger, "9\n4 9 95 93 57 4 57 93 9")

	c.Assert(out, Equals, "95\n")
}
