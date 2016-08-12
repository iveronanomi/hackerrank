package strings

import (
	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type CamelCaseSuite struct{}

var _ = Suite(&CamelCaseSuite{})

func (s *CamelCaseSuite) TestCase1(c *C) {
	out := utils.CaptureOut(CamelCase, "saveChangesInTheEditor")

	c.Assert(out, Equals, "5\n")
}
