package strings

import (
	"os"

	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type CamelCaseSuite struct{}

var _ = Suite(&CamelCaseSuite{})

func (s *CamelCaseSuite) TestCase1(c *C) {
	out := utils.CaptureStdout(func() {
		in := utils.WriteInput("saveChangesInTheEditor")
		defer in.Close()
		CamelCase(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "5\n")
}
