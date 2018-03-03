package implementation

import (
	. "gopkg.in/check.v1"

	"github.com/iveronanomi/hackerrank/utils"
)

type KangarooSuite struct{}

var _ = Suite(&KangarooSuite{})

func (s *KangarooSuite) TestKangaroo_Case1(c *C) {
	out := utils.Capture(Kangaroo, "0 3 4 2")

	c.Assert(out, Equals, "YES\n")
}

func (s *KangarooSuite) TestKangaroo_Case2(c *C) {
	out := utils.Capture(Kangaroo, "0 2 5 3")

	c.Assert(out, Equals, "NO\n")
}
