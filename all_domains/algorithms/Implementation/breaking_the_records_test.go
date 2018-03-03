package implementation

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type BreakingTheRecordsSuite struct{}

var _ = Suite(&BreakingTheRecordsSuite{})

func (s *BreakingTheRecordsSuite) TestBreakingTheRecords_Case1(c *C) {
	out := utils.Capture(BreakingTheRecords, "9\n10 5 20 20 4 5 2 25 1")

	c.Assert(out, Equals, "2 4\n")
}

func (s *BreakingTheRecordsSuite) TestBreakingTheRecords_Case2(c *C) {
	out := utils.Capture(BreakingTheRecords, "10\n3 4 21 36 10 28 35 5 24 42")

	c.Assert(out, Equals, "4 0\n")
}
