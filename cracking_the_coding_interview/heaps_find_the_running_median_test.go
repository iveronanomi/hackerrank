package cracking_the_coding_interview

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type HeapsFindTheRunningMedianSuite struct{}

var _ = Suite(&HeapsFindTheRunningMedianSuite{})

func (s *HeapsFindTheRunningMedianSuite) TestHeapsFindTheRunningMedian_Case1(c *C) {
	out := utils.Capture(HeapsFindTheRunningMedian, "6\n12\n4\n5\n3\n8\n7")

	c.Assert(out, Equals, "12.0\n8.0\n5.0\n4.5\n5.0\n6.0\n")
}

func (s *HeapsFindTheRunningMedianSuite) TestHeapsFindTheRunningMedian_Case2(c *C) {
	out := utils.Capture(HeapsFindTheRunningMedian, "10\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10")

	c.Assert(out, Equals, "1.0\n1.5\n2.0\n2.5\n3.0\n3.5\n4.0\n4.5\n5.0\n5.5\n")
}
