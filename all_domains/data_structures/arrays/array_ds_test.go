package arrays

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type Day2OperatorsSuite struct{}

var _ = Suite(&Day2OperatorsSuite{})

func (s *Day2OperatorsSuite) TestCase1(c *C) {
	out := utils.CaptureOut(ArrayDS, "4\n1 4 3 2")

	c.Assert(out, Equals, "2 3 4 1")
}
