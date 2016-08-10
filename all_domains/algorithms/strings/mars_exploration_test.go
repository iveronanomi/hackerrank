package strings

import (
	"os"

	. "gopkg.in/check.v1"

	"github.com/ereminIvan/hackerrank/utils"
)

type MarsExplorationSuite struct{}

var _ = Suite(&MarsExplorationSuite{})

func (s *MarsExplorationSuite) TestCase1(c *C) {
	out := utils.CaptureStdout(func() {
		in := utils.WriteInput("SOSSPSSQSSOR")
		defer in.Close()
		MarsExploration(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "3\n")
}

func (s *MarsExplorationSuite) TestCase2(c *C) {
	out := utils.CaptureStdout(func() {
		in := utils.WriteInput("SOSSOT")
		defer in.Close()
		MarsExploration(func() *os.File {
			return in
		})
	})

	c.Assert(out, Equals, "1\n")
}
