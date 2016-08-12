package strings

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type MarsExplorationSuite struct{}

var _ = Suite(&MarsExplorationSuite{})

func (s *MarsExplorationSuite) TestCase1(c *C) {
	out := utils.CaptureOut(MarsExploration, "SOSSPSSQSSOR")

	c.Assert(out, Equals, "3\n")
}

func (s *MarsExplorationSuite) TestCase2(c *C) {
	out := utils.CaptureOut(MarsExploration, "SOSSOT")

	c.Assert(out, Equals, "1\n")
}
