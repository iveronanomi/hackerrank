package warmup

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type CircularArrayRotationSuite struct{}

var _ = Suite(&CircularArrayRotationSuite{})

func (s *CircularArrayRotationSuite) TestCase1(c *C) {
	out := utils.Capture(CircularArrayRotation, "3 2 3\n1 2 3\n0\n1\n2")

	c.Assert(out, DeepEquals, "2\n3\n1\n")
}
