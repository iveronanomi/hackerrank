package strings

import (
	. "gopkg.in/check.v1"
)

type CamelCaseSuite struct{}

var _ = Suite(&CamelCaseSuite{})

func (s *CamelCaseSuite) TestCase1(c *C) {
	out := captureStdout(func() {
		CamelCase()
	}, "saveChangesInTheEditor")

	c.Assert(out, Equals, "5\n")
}
