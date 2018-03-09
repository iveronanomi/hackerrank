package cracking_the_coding_interview

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type QueuesATaleOfTwoStacksSuite struct{}

var _ = Suite(&QueuesATaleOfTwoStacksSuite{})

func (s *QueuesATaleOfTwoStacksSuite) TestQueuesATaleOfTwoStacks_Case1(c *C) {
	out := utils.Capture(QueuesATaleOfTwoStacks, "10\n1 42\n2\n1 14\n3\n1 28\n3\n1 60\n1 78\n2\n2")

	c.Assert(out, Equals, "14\n14\n")
}
