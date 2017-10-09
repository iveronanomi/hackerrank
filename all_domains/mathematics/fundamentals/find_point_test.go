package fundamentals

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type FindPointSuite struct{}

var _ = Suite(&FindPointSuite{})

func (s *FindPointSuite) TestCase1(c *C) {
	out := utils.Capture(FindPoint, "2\n0 0 1 1\n1 1 2 2")

	c.Assert(out, Equals, "2 2\n3 3\n")
}
