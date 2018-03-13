package cracking_the_coding_interview

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type StacksBalancedBracketsSuite struct{}

var _ = Suite(&StacksBalancedBracketsSuite{})

func (s *StacksBalancedBracketsSuite) TestStacksBalancedBrackets_Case1(c *C) {
	out := utils.Capture(StacksBalancedBrackets, "3\n{[()]}\n{[(])}\n{{[[(())]]}}")

	c.Assert(out, Equals, "YES\nNO\nYES\n")
}
