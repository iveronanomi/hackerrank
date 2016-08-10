package strings

import (
	. "gopkg.in/check.v1"
	"os"
)

type CamelCaseSuite struct{}

var _ = Suite(&CamelCaseSuite{})

func (s *CamelCaseSuite) TestCase1(c *C) {
	out := captureStdout(func() {
		in := writeInput("saveChangesInTheEditor")
		defer in.Close()
		CamelCase(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "5\n")
}
