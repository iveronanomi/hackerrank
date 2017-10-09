package search

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type SherlockAndArraySuite struct{}

var _ = Suite(&SherlockAndArraySuite{})

func (s *SherlockAndArraySuite) TestCase1(c *C) {
	out := utils.Capture(SherlockAndArray, "2\n3\n1 2 3\n4\n1 2 3 3")

	c.Assert(out, Equals, "NO\nYES\n")
}
